package models

import (
	"awesomeProject/databases"
)

type User struct{
	Name string
	Age int
	Birthday string
}

func GetAllUser() (user []User){
	db := databases.Connect()
	db.Find(&user).Scan(&user)

	databases.Connect_redis()

	return  user
}

func ConnectRedis(){
	databases.Connect_redis()
}
