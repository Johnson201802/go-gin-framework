package databases

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

func Connect() *gorm.DB {
	db, err := gorm.Open("mysql", "root:1e80a9ad15576fb9@tcp(127.0.0.1)/gin?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println("MYSQL连接失败了！")
		log.Fatal(err)
	} else {
		fmt.Println("MYSQL连接成功啦！")
	}


	return db

}
