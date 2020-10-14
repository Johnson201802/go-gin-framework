package router

import (
	"awesomeProject/apis"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine{
	router := gin.Default()
	router.GET("/", apis.IndexApi)
	return router
}