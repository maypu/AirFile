package middleware

import (
	"AirFile/utils"
	"github.com/gin-gonic/gin"
)

func InitGinMode() {
	if utils.GetConfig("common.environment") == "dev" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}
