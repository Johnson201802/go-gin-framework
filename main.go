package main

import (
	"awesomeProject/databases"
	"awesomeProject/router"
)

func main() {
	defer databases.Connect().Close()
	r := router.InitRouter()
	//models.ConnectRedis()

	r.Run(":8080")

}

