package main

import (
	"awesomeProject/router"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	//设置GIN运行模式
	gin.SetMode(gin.ReleaseMode)

	//初始化路由
	r := router.InitRouter()

	fmt.Println("服务已启动  ^_^  服务端口8090 请配置Nginx反向代理 注意云主机防火墙配置和服务商安全规则配置  Author: Jason")
	//启动GIN监听服务
	_ = r.Run(":8090")
}
