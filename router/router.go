package router

import (
	"AirFile/database"
	"AirFile/middleware"
	"AirFile/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routers(r *gin.Engine) *gin.Engine {
	// database
	db := database.InitDatabase()
	//db.Debug()
	db = middleware.Migrate(db)
	// api router
	v1 := r.Group("/api/v1")
	v1.POST("/config", func(c *gin.Context) {
		response := service.Config(c)
		c.JSON(http.StatusOK, response)
	})
	v1.POST("/auth", func(c *gin.Context) {
		response := service.Auth(c, db)
		c.JSON(http.StatusOK, response)
	})
	v1.POST("/upload", func(c *gin.Context) {
		response := service.Upload(c, db)
		c.JSON(http.StatusOK, response)
	})
	v1.POST("/download", func(c *gin.Context) {
		response := service.Download(c, db)
		c.JSON(http.StatusOK, response)
	})
	v1.POST("/verifyPwd", func(c *gin.Context) {
		response := service.VerifyPwd(c, db)
		c.JSON(http.StatusOK, response)
	})
	v1.GET("/file/:fileId", func(c *gin.Context) {
		service.File(c, db)
	})
	v1.POST("/history", func(c *gin.Context) {
		response := service.History(c, db)
		c.JSON(http.StatusOK, response)
	})
	return r
}
