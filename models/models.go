package models

import (
	"awesomeProject/databases"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
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

func GetMsmConfig(c *gin.Context) {
	MsmConfig := make(map[string]interface{}, 0)
	var config BaseConfig
	db := databases.Connect()
	db.Table("config").Where("config_id = ?", 1).First(&config).Scan(&config)
	MsmConfig["sms_app_id"] = config.Config_value

	db.Table("config").Where("config_id = ?", 2).First(&config).Scan(&config)
	MsmConfig["sms_app_key"] = config.Config_value

	db.Table("config").Where("config_id = ?", 3).First(&config).Scan(&config)
	MsmConfig["sms_sign"] = config.Config_value

	c.JSON(200, gin.H{
		"code": 200,
		"data": MsmConfig,
	})
}

func GetMchConfig(c *gin.Context) {
	MchConfig := make(map[string]interface{}, 0)
	var config BaseConfig
	db := databases.Connect()
	db.Table("config").Where("config_id = ?", 10).First(&config).Scan(&config)
	MchConfig["mch_appid"] = config.Config_value

	db.Table("config").Where("config_id = ?", 11).First(&config).Scan(&config)
	MchConfig["mch_key"] = config.Config_value

	db.Table("config").Where("config_id = ?", 12).First(&config).Scan(&config)
	MchConfig["url"] = config.Config_value

	c.JSON(200, gin.H{
		"code": 200,
		"data": MchConfig,
	})
}

type ConfigBase struct {
	Miniapp_id string
	Qrcode string
	Telphone string
	Miniapp_secrets string
	Tell_content string
}

func SaveConfigBase(c *gin.Context) {
	configBase := &ConfigBase{}
	if c.BindJSON(&configBase) == nil{
		db := databases.Connect()
		db.Table("config").Where("config_id = ?", 7).Update("config_value",configBase.Miniapp_id)
		db.Table("config").Where("config_id = ?", 9).Update("config_value",configBase.Miniapp_secrets)
		db.Table("config").Where("config_id = ?", 5).Update("config_value",configBase.Qrcode)
		db.Table("config").Where("config_id = ?", 13).Update("config_value",configBase.Tell_content)
		db.Table("config").Where("config_id = ?", 8).Update("config_value",configBase.Telphone)
		c.JSON(200,gin.H{"code":200,"msg":"OK！"})
	}else{
		c.JSON(200,gin.H{"msg":"非法请求！"})
	}

}

type ConfigSms struct {
	Sms_app_id string
	Sms_app_key string
	Sms_sign string
}

func SaveConfigSms(c *gin.Context){
	configSms := &ConfigSms{}
	if c.BindJSON(&configSms) == nil{
		db := databases.Connect()
		db.Table("config").Where("config_id = ?", 1).Update("config_value",configSms.Sms_app_id)
		db.Table("config").Where("config_id = ?", 2).Update("config_value",configSms.Sms_app_key)
		db.Table("config").Where("config_id = ?", 3).Update("config_value",configSms.Sms_sign)
		c.JSON(200,gin.H{"code":200,"msg":"OK！"})
	}else{
		c.JSON(200,gin.H{"msg":"非法请求！"})
	}
}

type ConfigMch struct {
	Mch_appid string
	Mch_key string
	Url string
}

func SaveConfigMch(c *gin.Context){
	configMch := &ConfigMch{}
	if c.BindJSON(&configMch) == nil{
		db := databases.Connect()
		db.Table("config").Where("config_id = ?", 10).Update("config_value",configMch.Mch_appid)
		db.Table("config").Where("config_id = ?", 11).Update("config_value",configMch.Mch_key)
		db.Table("config").Where("config_id = ?", 12).Update("config_value",configMch.Url)
		c.JSON(200,gin.H{"code":200,"msg":"OK！"})
	}else{
		c.JSON(200,gin.H{"msg":"非法请求！"})
	}
}

type AuthRule struct {
	Id int
	Name string
	Title string
	Status_t string
	Pid int
	Type int
}

func GetAuthList(c *gin.Context){
	var authrules[] AuthRule
	var count int
	page, _ := strconv.Atoi(c.Request.FormValue("page"))
	limit, _ := strconv.Atoi(c.Request.FormValue("limit"))
	start := (page-1)*limit
	sort := c.Request.FormValue("sort")

	if(sort=="+id"){
		sort = "Id DESC"
	}else{
		sort = "Id ASC"
	}

	db := databases.Connect()
	db.Table("auth_rule").Limit(limit).Offset(start).Order("Id desc").Find(&authrules).Scan(&authrules)
	db.Table("auth_rule").Count(&count)
	c.JSON(200,gin.H{"data":authrules,"total":count})

}

func DelRule(c *gin.Context){
	authrule := AuthRule{}
	id,_ := strconv.Atoi(c.Request.FormValue("id"))
	authrule.Id = id
	db := databases.Connect()
	db.Table("auth_rule").Where("id = ?", id).Delete(&authrule)
	c.JSON(200,gin.H{"code":200,"msg":"ok"})
}

func CreateRule(c *gin.Context){
	rule := &AuthRule{}
	rule.Pid = 0
	rule.Type = 1
	rule.Name = c.Request.FormValue("name")
	rule.Title = c.Request.FormValue("title")
	rule.Status_t = c.Request.FormValue("status_t")

	db := databases.Connect()
	db.Table("auth_rule").Create(&rule)
	c.JSON(200,gin.H{"code":200,"msg":"OK！"})
}
