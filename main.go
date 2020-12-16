package main

import (
	"awesomeProject/router"
	"github.com/gin-gonic/gin"
)

func main() {

	//设置GIN运行模式
	gin.SetMode(gin.ReleaseMode)

	//初始化路由
	r := router.InitRouter()

	//启动GIN监听服务
	r.Run(":8090")

}
