package database

import (
	"AirFile/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

var DB *gorm.DB

func InitDatabase() {
	dbDriver := utils.GetConfig("database.driver")
	if dbDriver == "sqlite" {
		DB = InitSQLite()
	} else if dbDriver == "mysql" {
		DB = InitMysql()
	}
	sqlDB, err := DB.DB()
	if err != nil {
		panic("连接数据库失败！")
	}
	sqlDB.SetMaxIdleConns(50)             // 空闲连接数
	sqlDB.SetMaxOpenConns(50)             // 最大连接数
	sqlDB.SetConnMaxLifetime(time.Minute) // 高并发情况下连接频繁失效，可修改为time.Hour
	//defer sqlDB.Close()
	//return db
}

func LoggerLevel() logger.Interface {
	loggerLevel := logger.Default.LogMode(logger.Silent)
	if utils.GetConfig("common.environment") == "dev" {
		loggerLevel = logger.Default.LogMode(logger.Info)
	}
	return loggerLevel
}
