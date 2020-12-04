package main

import (
	"awesomeProject/databases"
	"awesomeProject/router"
	"github.com/gin-gonic/gin"
)

func main() {

	//设置GIN运行模式
	gin.SetMode(gin.ReleaseMode)

	//程序执行完关闭数据库连接
	defer databases.Connect().Close()

	//初始化路由
	r := router.InitRouter()

	//启动GIN监听服务
	r.Run(":8090")

}
