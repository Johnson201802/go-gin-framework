package databases

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

func Connect() *gorm.DB{
	db , err := gorm.Open("mysql", "root:root@(127.0.0.1)/test?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil{
		fmt.Println("连接失败！")
		log.Fatal(err)
	}else{
		fmt.Println("连接成功！")
	}
	return  db
}