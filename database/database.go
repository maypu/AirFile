package database

import (
	"AirFile/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func InitDatabase() *gorm.DB {
	dbDriver := utils.GetConfig("database.driver")
	var db *gorm.DB
	if dbDriver == "sqlite" {
		db = InitSQLite()
	} else if dbDriver == "mysql" {
		db = InitMysql()
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic("连接数据库失败！")
	}
	sqlDB.SetMaxIdleConns(50) 			// 空闲连接数
	sqlDB.SetMaxOpenConns(50)			// 最大连接数
	sqlDB.SetConnMaxLifetime(time.Minute)	// 高并发情况下连接频繁失效，可修改为time.Hour
	return db
}

func LoggerLevel() logger.Interface {
	loggerLevel := logger.Default.LogMode(logger.Silent)
	if utils.GetConfig("common.environment") == "dev" {
		loggerLevel = logger.Default.LogMode(logger.Info)
	}
	return loggerLevel
}
