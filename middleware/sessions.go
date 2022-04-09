package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
)

//  调用文档 https://juejin.cn/post/6978879775808946207

// InitSessions 初始化session
func InitSessions(r *gin.Engine) gin.HandlerFunc {
	// 创建基于cookie的存储引擎，参数是用于加密的密钥
	//store := cookie.NewStore([]byte("AirFile"))
	store := memstore.NewStore([]byte("AirFile"))
	//设置session中间件，参数session的名字，也是cookie的名字；store存储引擎,可以替换成其他存储引擎
	r.Use(sessions.Sessions("sessionid", store))

	return func(c *gin.Context) {
		//// 初始化session对象
		//session := sessions.Default(c)
		////设置sessions的相关参数
		//session.Options(sessions.Options{MaxAge: 3600}) //单位：秒S
		c.Next()
	}
}

