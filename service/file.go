package service

import (
	"AirFile/database"
	"AirFile/model"
	"AirFile/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/OrlovEvgeny/go-mcache"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func Upload(c *gin.Context) *model.Response {
	response := model.NewResponse()
	uuidString := c.DefaultPostForm("uuid", "")
	password := c.DefaultPostForm("password", "")
	frequency := c.DefaultPostForm("frequency", "")
	limithours := c.DefaultPostForm("limithours", "")
	clintIp := c.ClientIP()
	f, err := c.FormFile("file")

	if err != nil {
		panic(err)
	}
	//取userid
	var mUser model.User
	database.DB.Where(&model.User{UUID: uuidString}).Find(&mUser)
	if err != nil {
		response.Code = 500
		response.Message = "获取文件失败"
		fmt.Println(err)
		c.JSON(http.StatusOK, response)
		return nil
	}
	fileExt := strings.ToLower(path.Ext(f.Filename))
	//if fileExt != ".png" && fileExt != ".jpg" && fileExt != ".gif" && fileExt != ".jpeg" && fileExt != ".apk" {
	//	response.Code = 500
	//	response.Message = "上传失败!只允许png,jpg,gif,jpeg文件"
	//	c.JSON(http.StatusOK, response)
	//	return
	//}
	maxSize, _ := strconv.Atoi(utils.GetConfig("upload.size"))
	if f.Size > int64(maxSize*1024*1024) {
		response.Code = 500
		response.Message = "文件最大限制为" + utils.GetConfig("upload.size") + "M！"
		c.JSON(http.StatusOK, response)
		return nil
	}
	random := utils.Random(6)
	fileName := random
	corrPath, _ := os.Getwd() //获取项目的执行路径
	fileDir := path.Join(utils.GetConfig("upload.path"), time.Now().Format("200601")) + "/"
	if runtime.GOOS == "windows" {
		strings.Replace(fileDir, "/", "\\", 5)
	} else {
		strings.Replace(fileDir, "\\", "/", -1)
	}
	//创建文件夹
	if _, err := os.Stat(path.Join(corrPath, fileDir)); os.IsNotExist(err) {
		err := os.MkdirAll(path.Join(corrPath, fileDir), os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
		err2 := os.Chmod(path.Join(corrPath, fileDir), os.ModePerm)
		if err2 != nil {
			log.Fatal(err2)
		}
	}
	//组装完整路径与文件名和后缀
	filepath := fmt.Sprintf("%s%s%s", corrPath+fileDir, fileName, fileExt)

	file, _, err := c.Request.FormFile("file")
	fileBuf := bytes.NewBuffer(nil)
	if _, err := io.Copy(fileBuf, file); err != nil {
		fmt.Println(err)
	}
	errEnc := utils.FileEncrypt(*fileBuf, filepath)

	//if err := c.SaveUploadedFile(f, filepath); err != nil {
	if !errEnc {
		fmt.Println(err)
		response.Code = 500
		response.Message = "上传失败"
	} else {
		// 存表
		frequencyInt, _ := strconv.Atoi(frequency)
		//n小时之后的时间
		hour, _ := time.ParseDuration(limithours + "h")
		nowAfterHour := time.Now().Add(hour)

		common := model.Common{IpAddress: clintIp}
		mFile := model.File{
			User:         mUser.ID,
			FileName:     f.Filename,
			Path:         fmt.Sprintf("%s%s%s", fileDir, fileName, fileExt),
			RandomCode:   random,
			Password:     password,
			NumDownloads: 0,
			LimitTimes:   frequencyInt,
			ExpiryTime:   nowAfterHour,
			Common:       common,
		}
		if err := database.DB.Create(&mFile).Error; err != nil {
			response.Code = 500
			response.Message = "上传失败"
			return nil
		}
		// 返回请求
		response.Message = "上传成功"
		response.Result = random
	}
	return response
}

func Download(c *gin.Context) *model.Response {
	response := model.NewResponse()
	params := make(map[string]interface{})
	c.BindJSON(&params)
	RandomCode := fmt.Sprintf("%v", params["code"])
	var mFile *model.File
	database.DB.Where(&model.File{RandomCode: RandomCode}).Find(&mFile)
	if mFile.Common.ID < 1 || mFile.NumDownloads >= mFile.LimitTimes {
		response.Code = 404
		response.Message = "资源不存在"
	} else {
		response.Code = 200
		type Result struct {
			File     bool `json:"file"`
			Password int  `json:"password"`
		}
		isPassword := 0
		if mFile.Password != "" {
			isPassword = 1
		}
		result := Result{File: true, Password: isPassword}
		result2, _ := json.Marshal(result)
		response.Result = string(result2)
	}
	return response
}

func VerifyPwd(c *gin.Context) *model.Response {
	response := model.NewResponse()
	params := make(map[string]interface{})
	c.BindJSON(&params)
	fileCode := fmt.Sprintf("%v", params["fileCode"])
	password := fmt.Sprintf("%v", params["password"])
	var mFile *model.File
	database.DB.Where(&model.File{RandomCode: fileCode}).Find(&mFile)
	if mFile.Common.ID < 1 {
		response.Code = 404
		response.Message = "资源不存在"
	} else {
		if password == mFile.Password {
			response.Code = 200
			response.Result = "true"
		} else {
			response.Code = 500
			response.Message = "密码错误"
			response.Result = "false"
		}
	}
	return response
}

func File(c *gin.Context) {
	response := model.NewResponse()
	fileCode := c.Param("fileCode")
	random := c.Query("random")
	if fileCode == "" || random == "" {
		response.Code = 500
		response.Message = "错误请求"
		c.JSON(http.StatusOK, response)
		return
	}
	MCache := mcache.New()
	var mFile *model.File
	database.DB.Where(&model.File{RandomCode: fileCode}).Find(&mFile)
	if mFile.Common.ID < 1 {
		response.Code = 404
		response.Message = "资源不存在"
		c.JSON(http.StatusOK, response)
		return
	}
	corrPath, _ := os.Getwd() //获取项目的执行路径
	wholePath := fmt.Sprintf("%s%s", corrPath, mFile.Path)
	file, errByOpenFile := os.Open(wholePath) //打开文件
	if errByOpenFile != nil {
		fmt.Println(errByOpenFile)
		response.Code = 404
		response.Message = "资源不存在了"
		c.JSON(http.StatusOK, response)
		return
	}
	defer file.Close()

	fileBuffer := utils.FileDecrypt(wholePath)
	if reflect.ValueOf(fileBuffer).IsNil() == true {
		response.Code = 404
		response.Message = "资源不存在了"
		c.JSON(http.StatusOK, response)
		return
	}

	// 下载次数+1
	randomCache, _ := MCache.Get("random")
	if randomCache != random {
		database.DB.Model(&model.File{}).Where("ID = ?", mFile.ID).Update("num_downloads", gorm.Expr("num_downloads + ?", 1))
		err := MCache.Set("random", random, time.Hour*1) //过期时间，1小时
		if err != nil {
			fmt.Println("MCache:", err)
		}
	}

	//fileStat, _ := file.Stat()

	// 发送下载header
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+mFile.FileName)
	c.Header("Content-Transfer-Encoding", "binary")
	//c.Header("Content-Length", strconv.FormatInt(fileStat.Size(), 10))
	//c.File(wholePath)

	//fileBuffer := bytes.NewfileBufferfer(nil)
	//if _, err := io.Copy(fileBuffer, file); err != nil {
	//	fmt.Println(err)
	//}

	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write(fileBuffer.Bytes())
}

func History(c *gin.Context) *model.Response {
	response := model.NewResponse()
	params := make(map[string]interface{})
	c.BindJSON(&params)
	uuidString := fmt.Sprintf("%v", params["uuid"])

	var mUser *model.User
	database.DB.Where(&model.User{UUID: uuidString}).First(&mUser)

	var mFile []model.File
	selectField := []string{"FileName", "RandomCode", "NumDownloads", "LimitTimes", "ExpiryTime", "CreatedAt", "DeletedAt", "Status"}
	database.DB.Select(selectField).Where(&model.File{User: mUser.ID}).Unscoped().Order("id desc").Find(&mFile) //Unscoped method to prevent adding deleted_at IS NULL
	mFile2, _ := json.Marshal(mFile)

	response.Code = 200
	response.Result = string(mFile2)
	return response
}
