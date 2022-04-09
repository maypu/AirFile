package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Static(r *gin.Engine) *gin.Engine {
	// static
	r.Static("/static", "./web/static")

	//html template
	r.LoadHTMLGlob("./web/view/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	return r
}
