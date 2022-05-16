package router

import (
	"AirFile/database"
	"AirFile/middleware"
	"AirFile/model"
	"AirFile/service"
	"AirFile/utils"
	"encoding/json"
	"fmt"
	"github.com/OrlovEvgeny/go-mcache"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func Routers(r *gin.Engine) *gin.Engine {
	// database
	db := database.InitDatabase()
	//db.Debug()
	db = middleware.Migrate(db)
	MCache := mcache.New()

	// api router
	v1 := r.Group("/api/v1")
	v1.POST("/config", func(c *gin.Context) {
		key := c.DefaultPostForm("key", "")
		response := service.Config(key)
		c.JSON(http.StatusOK, response)
	})

	v1.POST("/auth", func(c *gin.Context) {
		params := make(map[string]interface{})
		c.BindJSON(&params)
		uuidString := fmt.Sprintf("%v", params["uuid"])
		response := service.Auth(uuidString, db, c)
		c.JSON(http.StatusOK, response)
	})

	//// 新建/修改计算类型
	v1.POST("/upload", func(c *gin.Context) {
		response := model.NewResponse()
		uuidString := c.DefaultPostForm("uuid", "")
		//取userid
		var mUser model.User
		db.Where(&model.User{UUID: uuidString}).Find(&mUser)

		password := c.DefaultPostForm("password", "")
		frequency := c.DefaultPostForm("frequency", "")
		limithours := c.DefaultPostForm("limithours", "")
		clintIp := c.ClientIP()
		f, err := c.FormFile("file")
		if err != nil {
			response.Code = 500
			response.Message = "获取文件失败"
			fmt.Println(err)
			c.JSON(http.StatusOK, response)
			return
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
			return
		}
		random := utils.Random(6)
		fileName := random
		corrPath, _ := os.Getwd() //获取项目的执行路径
		fileDir := path.Join(utils.GetConfig("upload.path"), time.Now().Format("200601")) + "/"
		fmt.Println(fileDir)
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
		if err := c.SaveUploadedFile(f, filepath); err != nil {
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
			if err := db.Create(&mFile).Error; err != nil {
				response.Code = 500
				response.Message = "上传失败"
				return
			}
			// 返回请求
			response.Message = "上传成功"
			response.Result = random
		}
		c.JSON(http.StatusOK, response)
	})

	v1.POST("/download", func(c *gin.Context) {
		response := model.NewResponse()
		params := make(map[string]interface{})
		c.BindJSON(&params)
		RandomCode := fmt.Sprintf("%v", params["code"])

		var mFile *model.File
		db.Where(&model.File{RandomCode: RandomCode}).Find(&mFile)
		if mFile.Common.ID < 1 || mFile.NumDownloads >= mFile.LimitTimes {
			response.Code = 404
			response.Message = "资源不存在"
		} else {
			response.Code = 200
			type Result struct {
				File     uint `json:"file"`
				Password int  `json:"password"`
			}
			if mFile.Password != "" {
				result := Result{File: mFile.Common.ID, Password: 1}
				result2, _ := json.Marshal(result)
				response.Result = string(result2)
			} else {
				result := Result{File: mFile.Common.ID, Password: 0}
				result2, _ := json.Marshal(result)
				response.Result = string(result2)
			}
		}
		c.JSON(http.StatusOK, response)
	})
	v1.POST("/verifyPwd", func(c *gin.Context) {
		response := model.NewResponse()
		params := make(map[string]interface{})
		c.BindJSON(&params)
		fileId := fmt.Sprintf("%v", params["fileId"])
		password := fmt.Sprintf("%v", params["password"])

		var mFile *model.File
		db.Where("ID = " + fileId).Find(&mFile)
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
		c.JSON(http.StatusOK, response)
	})
	v1.GET("/file/:fileId", func(c *gin.Context) {
		response := model.NewResponse()
		fileId := c.Param("fileId")
		random := c.Query("random")
		if fileId == "" || random == "" {
			response.Code = 500
			response.Message = "错误请求"
			c.JSON(http.StatusOK, response)
			return
		}
		var mFile *model.File
		db.Where("ID = " + fileId).Find(&mFile)
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
		// 下载次数+1
		randomCache, _ := MCache.Get("random")
		if randomCache != random {
			db.Model(&model.File{}).Where("ID = "+fileId).Update("num_downloads", gorm.Expr("num_downloads + ?", 1))
			err := MCache.Set("random", random, time.Hour*1) //过期时间，1小时
			if err != nil {
				fmt.Println("MCache:", err)
			}
		}
		// 发送下载header
		c.Header("Content-Type", "application/octet-stream")
		c.Header("Content-Disposition", "attachment; filename="+mFile.FileName)
		c.Header("Content-Transfer-Encoding", "binary")
		c.File(wholePath)
	})
	v1.POST("/history", func(c *gin.Context) {
		response := model.NewResponse()
		params := make(map[string]interface{})
		c.BindJSON(&params)
		uuidString := fmt.Sprintf("%v", params["uuid"])

		var mUser *model.User
		db.Where(&model.User{UUID: uuidString}).First(&mUser)

		var mFile []model.File
		db.Where(&model.File{User: mUser.ID}).Unscoped().Order("id desc").Find(&mFile) //Unscoped method to prevent adding deleted_at IS NULL
		mFile2, _ := json.Marshal(mFile)

		response.Code = 200
		response.Result = string(mFile2)
		c.JSON(http.StatusOK, response)
	})
	return r
}
