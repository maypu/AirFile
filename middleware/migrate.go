package middleware

import (
	"AirFile/model"
	"fmt"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) *gorm.DB {

	if !db.Migrator().HasTable(&model.User{}) {
		db.AutoMigrate(&model.User{}, &model.File{}, &model.Text{}, &model.CornHistory{})
		fmt.Println("初始化允许，自动创建数据表完成！")
	}
	return db
}