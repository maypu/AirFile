package database

import (
	"AirFile/utils"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
)

var Sqlite *gorm.DB

func InitSQLite() *gorm.DB {
	corrPath, err := os.Getwd() //获取项目的执行路径
	if err != nil {
		fmt.Println(err)
	}
	loggerLevel := LoggerLevel()
	db, err := gorm.Open(sqlite.Open(corrPath+utils.GetConfig("sqlite.path")), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: utils.GetConfig("sqlite.tablePrefix"), // 表名前缀
		},
		Logger: loggerLevel,
	})
	if err != nil {
		panic(err)
	}
	Sqlite = db
	//Sqlite.DB().SetMaxIdleConns(1000)
	//Sqlite.DB().SetMaxOpenConns(100000)
	//Sqlite.DB().SetConnMaxLifetime(-1)

	//defer Sqlite.Close()
	return Sqlite
}
