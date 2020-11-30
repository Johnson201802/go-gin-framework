package models

import (
	"awesomeProject/databases"
	"awesomeProject/extend/oss"
	"fmt"
	"github.com/EDDYCJY/go-gin-example/pkg/qrcode"
	"github.com/boombuler/barcode/qr"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
	"strings"
	"time"
)

//定义管理员结构体
type Admin struct {
	Id         int
	Admin_name string
	Avatar string
	Admin_pwd  string
	Status  int
	Admin_login_ip  string
	Group_id int
	Introduction string
	Login_time string
	Title string
	Is_first_login int
}

//管理员登录
func GetAdmin(Admin_name string, Admin_pwd string) (flag bool, id int) {
	var admin Admin
	db := databases.Connect()
	db.Table("admin").Where("admin_name = ?", Admin_name).First(&admin).Scan(&admin)
	fmt.Println(admin.Admin_name + "___" + admin.Admin_pwd)
	if admin.Admin_name == Admin_name {
		if admin.Admin_pwd == Admin_pwd {
			return true, admin.Id
		} else {
			return false, 0
		}
	} else {
		return false, 0
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
	Status_t int
	Pid int
	Type string
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

	if c.BindJSON(&rule) == nil{
		db := databases.Connect()
		db.Table("auth_rule").Create(&rule)
		c.JSON(200,gin.H{"code":200,"msg":"OK！"})
	}else{
		c.JSON(200,gin.H{"msg":"非法请求！"})
	}
}

func UpdateRule(c *gin.Context){
	rule := &AuthRule{}
	if c.BindJSON(&rule) == nil{
		db := databases.Connect()
		db.Table("auth_rule").Update(&rule)
		c.JSON(200,gin.H{"code":200,"msg":"OK！"})
	}else{
		c.JSON(200,gin.H{"msg":"非法请求！"})
	}

}

func GetAdminList(c *gin.Context){
	var admin[] Admin

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
	db.Table("admin").Limit(limit).Offset(start).Order("Id desc").Select("admin.*, auth_group.title").Joins("left join auth_group on auth_group.id = admin.group_id").Find(&admin).Scan(&admin)
	db.Table("admin").Count(&count)
	c.JSON(200,gin.H{"data":admin,"total":count})
}

type Grouplist struct{
	Id int
	Title string
	Status int
	Rules string
	Description string
}

func GetGroupList(c *gin.Context){
	var grouplist[] Grouplist

	db := databases.Connect()
	db.Table("auth_group").Find(&grouplist)
	c.JSON(200,gin.H{"code":200,"data":grouplist})
}

//定义管理员结构体2
type Admin2 struct {
	Id         int
	Admin_name string
	Avatar string
	Admin_pwd  string
	Status  int
	Admin_login_ip  string
	Group_id int
	Introduction string
	Login_time string
	Is_first_login int
}

func CreateAdmin(c *gin.Context){
	admin2 := &Admin2{}

	if c.BindJSON(&admin2) == nil{
		db := databases.Connect()
		db.Table("admin").Create(&admin2)
		c.JSON(200,gin.H{"code":200,"msg":"OK！"})
	}else{
		c.JSON(200,gin.H{"msg":"非法请求！"})
	}
}

func DelAdmin(c *gin.Context){
	admin2 := Admin2{}
	id,_ := strconv.Atoi(c.Request.FormValue("id"))
	admin2.Id = id
	db := databases.Connect()
	db.Table("admin").Where("id = ?", id).Delete(&admin2)
	c.JSON(200,gin.H{"code":200,"msg":"ok"})
}

func UpdateAdmin(c *gin.Context){
	admin2 := &Admin2{}
	if c.BindJSON(&admin2) == nil{
		db := databases.Connect()
		db.Table("admin").Where("id = ?",admin2.Id).Update(&admin2)
		c.JSON(200,gin.H{"code":200,"msg":"OK！"})
	}else{
		c.JSON(200,gin.H{"msg":"非法请求！"})
	}
}

type Auth_group struct{
	Id int
	Title string
	Status int
	Rules string
	Description string
	Name string
}

//用户组管理
func GetRoleList(c *gin.Context){
	var auth_group[] Auth_group
	var authrule[] AuthRule
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
	db.Table("auth_group").Limit(limit).Offset(start).Order("id desc").Find(&auth_group).Scan(&auth_group)

	for key , items := range auth_group {
		arr := strings.Split(items.Rules,",")

		var dd[] int
		for _ ,item := range arr{
			dd2,_ := strconv.Atoi(item)
			dd = append(dd, dd2)
		}

		db.Table("auth_rule").Where("id in (?)", dd).Find(&authrule)

		for _ ,item2 := range authrule{
			auth_group[key].Name += item2.Title + ","
		}

	}

	db.Table("auth_group").Count(&count)
	c.JSON(200,gin.H{"data":auth_group,"total":count})
}


func GetRuleList(c *gin.Context){
	var authrule[] AuthRule

	db := databases.Connect()
	db.Table("auth_rule").Find(&authrule)
	c.JSON(200,gin.H{"code":200,"data":authrule})
}

type Auth_group2 struct{
	Id int
	Title string
	Status int
	Rules string
	Description string
}

func AddRole(c *gin.Context){
	var authgroup Auth_group2

	if c.BindJSON(&authgroup)==nil {
		db := databases.Connect()
		db.Table("auth_group").Create(&authgroup)
		c.JSON(200,gin.H{"code":200})
	}
}

func DelRole(c *gin.Context){
	var authgroup Auth_group2

	id := c.Request.FormValue("id")
	id2 , _:= strconv.Atoi(id)
	authgroup.Id = id2
	db := databases.Connect()
	db.Table("auth_group").Delete(&authgroup)
	c.JSON(200,gin.H{"code":200})
}

func GetOneRole(c *gin.Context){
	var authgroup Auth_group2

	id := c.Request.FormValue("Id")
	id2 , _:= strconv.Atoi(id)
	db := databases.Connect()
	db.Table("auth_group").First(&authgroup,id2)

	ids := strings.Split(authgroup.Rules,",")
	c.JSON(200,gin.H{"code":200,"data":ids})
}

func UpdateRole(c *gin.Context){
	var authgroup Auth_group2

	if c.BindJSON(&authgroup)==nil{
		db := databases.Connect()
		db.Table("auth_group").Where("id=?",authgroup.Id).Update(&authgroup)
		c.JSON(200,gin.H{"code":200})
	}
}

//定义用户结构体
type User struct{
	User_id int
	Nick_name string
	Avatar string
	Phone int
	Open_id string
	Merchant_id int
	Is_vip string
	Expire_vip_time int
}

func FetchUserList(c *gin.Context){
	var user[] User
	var count int

	page, _ := strconv.Atoi(c.Request.FormValue("page"))
	limit, _ := strconv.Atoi(c.Request.FormValue("limit"))
	start := (page-1)*limit
	sort := c.Request.FormValue("sort")
	title := c.Request.FormValue("title")

	if(sort=="+id"){
		sort = "user_id DESC"
	}else{
		sort = "user_id ASC"
	}
	db := databases.Connect()
	if title == "" {
		db.Table("user").Limit(limit).Offset(start).Order(sort).Select("user_id,phone,nick_name,merchant_id,avatar,is_vip,expire_vip_time").Find(&user).Scan(&user)
	}else{
		db.Table("user").Limit(limit).Offset(start).Order(sort).Select("user_id,phone,nick_name,merchant_id,avatar,is_vip,expire_vip_time").Where("nick_name like ?",title).Find(&user).Scan(&user)
	}

	db.Table("user").Count(&count)

	c.JSON(200,gin.H{"code":200,"data":user,"total":count})
}

func DelUser(c *gin.Context){
	var user User

	id := c.Request.FormValue("id")
	id2 ,_ := strconv.Atoi(id)
	user.User_id = id2

	db := databases.Connect()
	db.Table("user").Delete(&user)
	c.JSON(200,gin.H{"code":200})
}

//文章结构体
type Articles struct{
	Article_id int
	Title string
	Content string
	Pv int
	Create_time int64
}

func FetchArticleList(c *gin.Context){
	var article[] Articles
	var count int

	page, _ := strconv.Atoi(c.Request.FormValue("page"))
	limit, _ := strconv.Atoi(c.Request.FormValue("limit"))
	start := (page-1)*limit
	sort := c.Request.FormValue("sort")
	title := c.Request.FormValue("title")

	if(sort=="+id"){
		sort = "article_id DESC"
	}else{
		sort = "article_id ASC"
	}

	db := databases.Connect()
	if title == "" {
		db.Table("articles").Limit(limit).Offset(start).Order(sort).Find(&article).Scan(&article)
	}else{
		db.Table("articles").Limit(limit).Offset(start).Order(sort).Where("title like ?",title).Find(&article).Scan(&article)
	}

	db.Table("articles").Count(&count)

	c.JSON(200,gin.H{"code":200,"data":article,"total":count})

}

func DelArticle(c *gin.Context){
	var article Articles

	id := c.Request.FormValue("id")
	id2 ,_ := strconv.Atoi(id)
	article.Article_id = id2

	db := databases.Connect()
	db.Table("articles").Delete(&article)
	c.JSON(200,gin.H{"code":200})
}

func CreateArticle(c *gin.Context){
	var article Articles

	article.Create_time = time.Now().Unix()
	if c.BindJSON(&article)==nil{
		db := databases.Connect()
		db.Table("articles").Create(&article)
		c.JSON(200,gin.H{"code":200})
	}else{
		c.JSON(200,gin.H{"code":400})
	}
}

func FetchOneArticle(c *gin.Context){
	var article Articles

	id := c.Request.FormValue("id")
	id2 ,_ := strconv.Atoi(id)

	db := databases.Connect()
	db.Table("articles").Where("article_id=?",id2).First(&article)
	c.JSON(200,gin.H{"code":200,"data":article})
}

func UpdateArticle(c *gin.Context){
	var article Articles

	if c.BindJSON(&article)==nil{
		db := databases.Connect()
		db.Table("articles").Update(&article)
		c.JSON(200,gin.H{"code":200})
	}else{
		c.JSON(200,gin.H{"code":400})
	}
}

//商户结构体
type Merchant struct {
	Merchant_id int
	Name        string
	Mobile      string
	Img1        string
	Img2        string
	Longitude   float32
	Latitude    float32
	Address     string
	Order       int
	Create_time int64
	ID          interface{}
	Qrcode string
}

func FetchMerchantList(c *gin.Context){
	var merchant[] Merchant
	var count int

	page, _ := strconv.Atoi(c.Request.FormValue("page"))
	limit, _ := strconv.Atoi(c.Request.FormValue("limit"))
	start := (page-1)*limit
	title := c.Request.FormValue("title")


	db := databases.Connect()
	if title == "" {
		db.Table("merchant").Limit(limit).Offset(start).Find(&merchant).Scan(&merchant)
	}else{
		db.Table("merchant").Limit(limit).Offset(start).Where("title like ?",title).Find(&merchant).Scan(&merchant)
	}

	db.Table("merchant").Count(&count)

	c.JSON(200,gin.H{"code":200,"data":merchant,"total":count})
}

func DelMerchant(c *gin.Context){
	var merchant Merchant

	id := c.Request.FormValue("id")
	id2 ,_ := strconv.Atoi(id)
	merchant.Merchant_id = id2

	db := databases.Connect()
	db.Table("merchant").Delete(&merchant)
	c.JSON(200,gin.H{"code":200})
}

//商户结构体
type Merchant4 struct {
	Merchant_id int
	Name        string
	Mobile      string
	Img1        string
	Img2        string
	Longitude   float32
	Latitude    float32
	Address     string
	Order       int
	Create_time int64
	Qrcode string
}

func CreateMerchant(c *gin.Context){
	var merchant Merchant
	var merchant4 Merchant4

	merchant.Create_time = time.Now().Unix()
	if c.BindJSON(&merchant)==nil {
		db := databases.Connect()
		db.Table("merchant").Create(&merchant)

		id := fmt.Sprintf("%v", merchant.ID)
		img := GetQRcode(id)
		url := oss.UploadFile2(c, img)
		merchant.Qrcode = url

		merchant4.Name = merchant.Name
		merchant4.Mobile = merchant.Mobile
		merchant4.Img1 = merchant.Img1
		merchant4.Img2 = merchant.Img2
		merchant4.Longitude = merchant.Longitude
		merchant4.Latitude = merchant.Latitude
		merchant4.Address = merchant.Address
		merchant4.Order = merchant.Order
		merchant4.Create_time = merchant.Create_time
		merchant4.Qrcode = url

		fmt.Println(id)
		db.Table("merchant").Where("merchant_id=?",id).Update(&merchant4)
		//
		_ = os.Remove(img)

		c.JSON(200,gin.H{"code":200})
	}
}

func FetchOneMerchant(c *gin.Context){
	var merchant Merchant

	id , _ := strconv.Atoi(c.Request.FormValue("id"))
	merchant.Merchant_id = id
	db := databases.Connect()
	db.Table("merchant").Where("merchant_id=?",id).First(&merchant)
	c.JSON(200,gin.H{"code":200,"data":merchant})
}


func UpdateMerchant(c *gin.Context){
	var merchant Merchant

	if c.BindJSON(&merchant)==nil{
		db := databases.Connect()
		db.Table("merchant").Where("merchant_id=?",merchant.Merchant_id).Update(&merchant)
		c.JSON(200,gin.H{"code":200})
	}else{
		c.JSON(200,gin.H{"code":400})
	}
}

//订单结构体
type Order struct{
	Order_id int
	Merchant_id int
	Name string
	Service_id string
	User_id int
	Nick_name string
	Price float32
	Status int
	Stars int
	Conment string
}

func FetchOrderList(c *gin.Context){
	var order[] Order
	var count int

	page, _ := strconv.Atoi(c.Request.FormValue("page"))
	limit, _ := strconv.Atoi(c.Request.FormValue("limit"))
	start := (page-1)*limit

	db := databases.Connect()
	db.Table("order").Select("order.order_id, order.merchant_id, order.user_id, order.price, order.status, order.stars, order.conment, merchant.name, user.nick_name").Joins("join merchant on merchant.merchant_id = order.merchant_id").Joins("join user on user.user_id = order.user_id").Limit(limit).Offset(start).Find(&order).Scan(&order)

	db.Table("order").Count(&count)

	c.JSON(200,gin.H{"code":200,"data":order,"total":count})

}

func DelOrder(c *gin.Context){
	var order Order

	id := c.Request.FormValue("id")
	id2 ,_ := strconv.Atoi(id)
	order.Order_id = id2

	db := databases.Connect()
	db.Table("order").Delete(&order)
	c.JSON(200,gin.H{"code":200})
}

//服务结构体
type Service struct {
	Service_id int
	Service_name string
	Name string
	Icon string
	Origin_price float32
	Now_price float32
	Pid int
	Is_sale string
}

func FetchServicesList(c *gin.Context){
	var service[] Service
	var count int

	page, _ := strconv.Atoi(c.Request.FormValue("page"))
	limit, _ := strconv.Atoi(c.Request.FormValue("limit"))
	start := (page-1)*limit

	db := databases.Connect()
	db.Debug().Table("service").Select("service.service_id, service.service_name, service.icon, service.origin_price, service.now_price, service.pid, service.is_sale, merchant.name").Joins("join merchant on merchant.merchant_id = service.pid").Limit(limit).Offset(start).Find(&service).Scan(&service)

	db.Table("service").Count(&count)

	c.JSON(200,gin.H{"code":200,"data":service,"total":count})

}

func GetQRcode(id string)(string){
	qrc := qrcode.NewQrCode(id, 300, 300, qr.M, qr.Auto)
	path := qrcode.GetQrCodeFullPath()
	one, _, err := qrc.Encode(path)
	if err != nil {
		return one
	}else{
		return one
	}
}

func ChangeStatus(c *gin.Context){
	var service Service

	id , _:= strconv.Atoi(c.Request.FormValue("id"))

	service.Service_id = id
	service.Is_sale = c.Request.FormValue("is_sale")

	db := databases.Connect()
	db.Table("service").Where("service_id=?",id).Update(&service)

	c.JSON(200,gin.H{"code":200})
}

//服务结构体
type Service2 struct {
	Service_id int
	Service_name string
	Icon string
	Origin_price string
	Now_price string
	Pid int
	Is_sale string
}

func CreateService(c *gin.Context){
	var service Service2

	if c.BindJSON(&service)==nil{
		db := databases.Connect()
		db.Table("service").Create(&service)

		c.JSON(200,gin.H{"code":200})
	}else{
		c.JSON(200,gin.H{"code":400})
	}
}

//商户结构体
type Merchant2 struct{
	Merchant_id int
	Name string
}

func GetMerchant(c *gin.Context){
	var merchant[] Merchant2

	db := databases.Connect()
	db.Table("merchant").Select("merchant_id, name").Find(&merchant).Scan(&merchant)
	c.JSON(200,gin.H{"code":200, "data":merchant})
}

//服务结构体
type Service3 struct {
	Service_id int
	Service_name string
	Icon string
	Origin_price string
	Now_price string
	Pid int
	Is_sale string
}

func UpdateService(c *gin.Context){
	var service Service3

	if c.BindJSON(&service)==nil{
		db := databases.Connect()
		db.Table("service").Where("service_id=?",service.Service_id).Update(&service)
		c.JSON(200,gin.H{"code":200})
	}else{
		c.JSON(200,gin.H{"code":400})
	}
}

//消息结构体
type  Question struct{
	Question_id int
	User_id2 int
	Content string
	Time int
	Is_read string
	Nick_name string
	Avatar string
	Phone int
}

func GetNewMsg(c *gin.Context){
	var count int

	db := databases.Connect()
	db.Table("question").Where("is_read = ?", 0).Count(&count)

	if count == 0{
		c.JSON(200,gin.H{"code":300})
	}else{
		c.JSON(200,gin.H{"code":200})
	}

}

func GetQuestionList(c *gin.Context){
	var question[] Question
	var count int

	page, _ := strconv.Atoi(c.Request.FormValue("page"))
	limit, _ := strconv.Atoi(c.Request.FormValue("limit"))
	start := (page-1)*limit

	db := databases.Connect()
	err := db.Debug().Table("question").Select("question.question_id, question.user_id2, question.content, question.time, question.is_read, user.nick_name, user.avatar, user.phone, user.user_id").Joins("join user on user.user_id = question.user_id2").Order("question_id desc").Limit(limit).Offset(start).Find(&question).Scan(&question).Error
	db.Table("question").Count(&count)

	if err == nil{
		c.JSON(200,gin.H{"code":200,"data":question,"count":count})
	}else{
		c.JSON(200,gin.H{"code":300,"msg":"error"})
	}
}

func SetRead(c *gin.Context){
	var question Question

	id, _ := strconv.Atoi(c.Request.FormValue("id"))

	question.Question_id = id
	question.Is_read = "1"
	db := databases.Connect()
	err := db.Debug().Table("question").Where("question_id = ?",id).Update(&question).Error

	if err != nil{
		c.JSON(200,gin.H{"code":300,"msg":"error"})
	}else{
		c.JSON(200,gin.H{"code":200,"msg":"success"})
	}
}

//消息结构体
type  Question2 struct{
	Question_id int
	User_id2 int
	Content string
}

func ArticleCreate(c *gin.Context){
	var question Question2

	if c.BindJSON(&question)==nil{
		db := databases.Connect()
		db.Table("question").Create(&question)

		c.JSON(200,gin.H{"code":200})
	}else{
		c.JSON(200,gin.H{"code":400})
	}
}