package middleware

import (
	"AirFile/database"
	"AirFile/model"
	"fmt"
)

func Migrate() {
	database.DB.AutoMigrate(&model.User{}, &model.File{}, &model.Text{}, &model.CornHistory{})
	fmt.Println("数据表结构同步完成！")
}
