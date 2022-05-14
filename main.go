package main

import (
	"AirFile/middleware"
	"AirFile/router"
	"AirFile/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)


// go env -w CGO_ENABLED=0
// go env -w GOOS=linux
// go env -w GOARCH=amd64
// go build
func main()  {
	r := gin.Default()
	//gin.SetMode(gin.ReleaseMode)
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
