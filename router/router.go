package router

import (
	"awesomeProject/apis"
	"github.com/gin-gonic/gin"
)

//跨域中间件配置
func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("origin")
		if len(origin) == 0 {
			origin = c.Request.Header.Get("Origin")
		}
		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With,x-token")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST")
		c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}
		c.Next()
	}
}

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors())
	router.POST("/adminLogin", apis.IndexApi)

	//路由组没写加上JWT鉴权
	router.GET("/getInfo", apis.GetInfo)
	router.GET("/serverInfo", apis.GetServerInfo)
	router.GET("/getBaseconfig", apis.GetBaseconfig)
	router.GET("/getMsmConfig", apis.GetMsmConfig)
	router.GET("/getMchConfig", apis.GetMchConfig)
	router.POST("/imgUploads", apis.ImgUploads)
	router.GET("/sendMSM", apis.SendMSM)

	return router
}
