package middleware

import (
	"AirFile/database"
	"AirFile/model"
	"AirFile/utils"
	"fmt"
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
	"os"
	"time"
)

func MainCron() {
	c := cron.New()
	i := 1
	c.AddFunc("*/1 * * * *", func() {
		fmt.Println("自动任务执行：", i, "次")
		AutoDeleteFile()
		i++
	})
	c.Start()
}

func AutoDeleteFile()  {
	// database
	dbDriver := utils.GetConfig("database.driver")
	var db *gorm.DB
	if dbDriver == "sqlite" {
		db = database.InitSQLite()
	} else if dbDriver == "mysql" {
		db = database.InitMysql()
	}
	corrPath, _ := os.Getwd() //获取项目的执行路径

	var mFile []model.File
	db.Find(&mFile)
	for i := range mFile {
		if mFile[i].NumDownloads >= mFile[i].LimitTimes ||  mFile[i].ExpiryTime.Unix() < time.Now().Unix() {
			wholePath := fmt.Sprintf("%s%s", corrPath, mFile[i].Path)
			err := os.Remove(wholePath)
			if err != nil {
				fmt.Println(err)
				fmt.Println("文件删除失败！")
			}
			db.Delete(&mFile)
		}
	}
}