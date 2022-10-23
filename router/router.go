package router

import (
	"AirFile/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routers(r *gin.Engine) *gin.Engine {
	// api router
	v1 := r.Group("/api/v1")
	v1.POST("/config", func(c *gin.Context) {
		response := service.Config(c)
		c.JSON(http.StatusOK, response)
	})
	v1.POST("/auth", func(c *gin.Context) {
		response := service.Auth(c)
		c.JSON(http.StatusOK, response)
	})
	v1.POST("/upload", func(c *gin.Context) {
		response := service.Upload(c)
		c.JSON(http.StatusOK, response)
	})
	v1.POST("/download", func(c *gin.Context) {
		response := service.Download(c)
		c.JSON(http.StatusOK, response)
	})
	v1.POST("/verifyPwd", func(c *gin.Context) {
		response := service.VerifyPwd(c)
		c.JSON(http.StatusOK, response)
	})
	v1.GET("/file/:fileCode", func(c *gin.Context) {
		service.File(c)
	})
	v1.POST("/history", func(c *gin.Context) {
		response := service.History(c)
		c.JSON(http.StatusOK, response)
	})
	return r
}
