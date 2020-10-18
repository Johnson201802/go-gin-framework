package models

import (
	"awesomeProject/databases"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Admin struct {
	Id         string
	Admin_name string
	Admin_pwd  string
}

//管理员登录
func GetAdmin(Admin_name string, Admin_pwd string) (flag bool, id string) {
	var admin Admin
	db := databases.Connect()
	db.Table("admin").Where("admin_name = ?", Admin_name).First(&admin).Scan(&admin)
	fmt.Println(admin.Admin_name + "___" + admin.Admin_pwd)
	if admin.Admin_name == Admin_name {
		if admin.Admin_pwd == Admin_pwd {
			return true, admin.Id
		} else {
			return false, ""
		}
	} else {
		return false, ""
	}

}

//连接REDIS
func ConnectRedis() {
	databases.Connect_redis()
}

//获取系统配置
//基础配置
type BaseConfig struct {
	Config_id         int
	Config_type       int
	Config_name       string
	Config_name_other string
	Config_value      string
}

func GetBaseconfig(c *gin.Context) {
	Baseconfig := make(map[string]interface{}, 0)
	var baseconfig BaseConfig
	db := databases.Connect()
	db.Table("config").Where("config_id = ?", 5).First(&baseconfig).Scan(&baseconfig)
	Baseconfig["qrcode"] = baseconfig.Config_value

	db.Table("config").Where("config_id = ?", 7).First(&baseconfig).Scan(&baseconfig)
	Baseconfig["miniapp_id"] = baseconfig.Config_value

	db.Table("config").Where("config_id = ?", 8).First(&baseconfig).Scan(&baseconfig)
	Baseconfig["telphone"] = baseconfig.Config_value

	db.Table("config").Where("config_id = ?", 9).First(&baseconfig).Scan(&baseconfig)
	Baseconfig["miniapp_secrets"] = baseconfig.Config_value

	db.Table("config").Where("config_id = ?", 13).First(&baseconfig).Scan(&baseconfig)
	Baseconfig["tell_content"] = baseconfig.Config_value

	c.JSON(200, gin.H{
		"code": 200,
		"data": Baseconfig,
	})
}
