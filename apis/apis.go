package apis

import (
	"awesomeProject/databases"
	"awesomeProject/extend/oss"
	"awesomeProject/extend/sms"
	"awesomeProject/models"
	"awesomeProject/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
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

func GetMsmConfig(c *gin.Context) {
	models.GetMsmConfig(c)
}

func GetMchConfig(c *gin.Context) {
	models.GetMchConfig(c)
}

func ImgUploads(c *gin.Context) {
	f, err := c.FormFile("file")
	if err != nil {
		c.JSON(200, gin.H{
			"error": err,
		})
		return
	} else {
		err := c.SaveUploadedFile(f, f.Filename)
		if err != nil {
			c.JSON(200, gin.H{
				"error": err,
			})
			return
		} else {
			UploadFile(c, f.Filename)
			err := os.Remove(f.Filename)
			db := databases.Connect()
			db.Table("config").Where("config_id = ?", 5).Update("config_value","https://img-c-jason.oss-accelerate.aliyuncs.com/"+f.Filename)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func SendMSM(c *gin.Context) {
	sms.SingleSms()
}

func GetOss(c *gin.Context) {
	oss.ListFile(c)
}

func UploadFile(c *gin.Context, filename string) {
	oss.UploadFile(c, filename)
}

func GetStorageList(c *gin.Context) {
	oss.GetStorageList(c)
}

func SaveConfigBase(c *gin.Context) {
	models.SaveConfigBase(c)
}

func SaveConfigSms(c *gin.Context){
	models.SaveConfigSms(c)
}

func SaveConfigMch(c *gin.Context){
	models.SaveConfigSms(c)
}

func GetAuthList(c *gin.Context){
	models.GetAuthList(c)
}

func DelRule(c *gin.Context){
	models.DelRule(c)
}

func CreateRule(c *gin.Context){
	models.CreateRule(c)
}