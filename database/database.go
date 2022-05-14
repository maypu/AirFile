package database

import (
	"AirFile/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDatabase() *gorm.DB {
	dbDriver := utils.GetConfig("database.driver")
	var db *gorm.DB
	if dbDriver == "sqlite" {
		db = InitSQLite()
	} else if dbDriver == "mysql" {
		db = InitMysql()
	}
	return db
}

func LoggerLevel() logger.Interface {
	loggerLevel := logger.Default.LogMode(logger.Silent)
	if utils.GetConfig("common.environment") == "dev" {
		loggerLevel = logger.Default.LogMode(logger.Info)
	}
	return loggerLevel
}
