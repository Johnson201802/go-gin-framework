package databases

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func Connect_redis(){
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
	}else{
		fmt.Println("Connect to redis success!")
	}
	c.Do("SET", "mykey", "superWang")
	defer c.Close()
}