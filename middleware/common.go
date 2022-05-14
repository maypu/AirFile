package middleware

import (
	"AirFile/utils"
	"github.com/gin-gonic/gin"
)

func InitGin() *gin.Engine {
	if utils.GetConfig("common.environment") == "dev" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	return gin.Default()
}
