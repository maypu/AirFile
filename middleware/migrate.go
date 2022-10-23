package middleware

import (
	"AirFile/database"
	"AirFile/model"
	"fmt"
)

func Migrate() {
	if !database.DB.Migrator().HasTable(&model.User{}) {
		database.DB.AutoMigrate(&model.User{}, &model.File{}, &model.Text{}, &model.CornHistory{})
		fmt.Println("初始化允许，自动创建数据表完成！")
	}
	//return db
}
