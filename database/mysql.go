package database

import (
	"AirFile/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm/schema"
)
import "gorm.io/gorm"

func InitMysql() *gorm.DB {
	loggerLevel := LoggerLevel()
	// "用户名:密码@(127.0.0.1:端口号)/数据库名称?charset=utf8&parseTime=True&loc=Local"
	dsn := utils.GetConfig("mysql.user") + ":" + utils.GetConfig("mysql.password") + "@tcp(" + utils.GetConfig("mysql.domain") + ":" + utils.GetConfig("mysql.port") + ")/" + utils.GetConfig("mysql.dbname") + "?charset=" + utils.GetConfig("mysql.charset") + "&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: utils.GetConfig("mysql.prefix"), // 表名前缀
			//SingularTable: true,
		},
		Logger: loggerLevel,
	})
	if err != nil {
		panic(err)
	}
	return db
}
