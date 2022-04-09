package main

import (
	"AirFile/middleware"
	"AirFile/router"
	"AirFile/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)


// CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
func main()  {
	r := gin.Default()
	r.Use(middleware.Cors())
	//r.Use(middleware.InitSessions(r))
	r = router.Routers(r)	// 加载路由
	r = router.Static(r)
	middleware.MainCron()	//挂载自动任务
	runPort := utils.GetConfig("common.port")
	if runPort == "" {
		runPort = "8085"
	}
	err := r.Run(":" + runPort) //run in terminal: fresh
	fmt.Println(err)
}
