package apis

import (
	"awesomeProject/databases"
	"awesomeProject/models"
	"awesomeProject/tools"
	"fmt"
	"github.com/gin-gonic/gin"
)

var roles = [2]string{"admin", "edit"}

func IndexApi(c *gin.Context) {
	name := c.Request.FormValue("username")
	pwd := tools.MD5(c.Request.FormValue("password"))
	flag, uid := models.GetAdmin(name, pwd)

	//生成TOKEN
	token, _ := tools.CreateToken(uid, "johnson2018092789728376273672364wey")
	red := databases.Connect_redis()
	defer red.Close()
	var err error
	_, err = red.Do("SET", uid, token)
	_, err = red.Do("expire", uid, 7200) //两个小时过期
	if err != nil {
		fmt.Println("set expire error: ", err)
		return
	}
	if flag {
		c.JSON(200, gin.H{
			"code":         200,
			"roles":        roles,
			"name":         name,
			"avatar":       "234567",
			"introduction": "",
			"token":        token,
		})
	} else {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  "",
		})
	}
}

func GetInfo(c *gin.Context) {
	token := c.Request.Header.Get("X-Token")
	uid, _ := tools.ParseToken(token, "johnson2018092789728376273672364wey")
	red := databases.Connect_redis()
	v, _ := red.Do("GET", uid)
	if v == nil {
		c.JSON(200, gin.H{
			"code": 600,
		})
	} else {
		c.JSON(200, gin.H{
			"code":         200,
			"roles":        "admin",
			"name":         "jason",
			"avatar":       "12345",
			"introduction": "234567",
		})
	}
	defer red.Close()
}

func GetServerInfo(c *gin.Context) {
	tools.ServerInfo(c)
}

func GetBaseconfig(c *gin.Context) {
	models.GetBaseconfig(c)
}
