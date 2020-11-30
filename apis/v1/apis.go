package v1

import (
	"awesomeProject/databases"
	"awesomeProject/extend/wechat"
	"awesomeProject/extend/wechatpay"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"strconv"
)

type Merchant struct{
	Merchant_id int
	Longitude string
	Latitude string
}

func GetMerchant2(c *gin.Context){
	var merchant[] Merchant

	db := databases.Connect()
	db.Table("merchant").Find(&merchant)
	c.JSON(200,gin.H{"data":merchant})
}

type User struct{
	User_id int
	Nick_name string
	Avatar string
	Phone int
	Open_id string
	Merchant_id int
	Is_vip string
}

func GetOpenid(c *gin.Context){
	var user User
	code := c.Request.FormValue("code")
	nick_name := c.Request.FormValue("nick_name")
	avatar := c.Request.FormValue("avatar")

	//fmt.Println(code+","+nick_name+","+avatar)
	res , _ := wechat.WXLogin(code,"wxb7475adc9f5980d2","2fcbbe27e765628bf3c61f8ff384ccac")
	db := databases.Connect()
	err := db.Table("user").Where("open_id = ?",res.OpenId).First(&user).Scan(&user).Error
	if err != nil {
		user.Avatar = avatar
		user.Nick_name = nick_name
		user.Open_id = res.OpenId
		db.Table("user").Create(&user)
		c.JSON(200,gin.H{"code":200})
	}else{
		c.JSON(200,gin.H{"code":200,"user":user,"session_id":res.SessionKey,"openid":res.OpenId})
	}

}

type Result struct{
	PhoneNumber string
	PurePhoneNumber string
	CountryCode string
}

func GetPhoneNumber(c *gin.Context){
	var result Result
	var user User
	encryptedData := c.Request.FormValue("encryptedData")
	iv := c.Request.FormValue("iv")
	session_id := c.Request.FormValue("session_id")
	user_id := c.Request.FormValue("user_id")

	src, _ := wechat.Dncrypt(encryptedData,session_id,iv)
	json.Unmarshal([]byte(src), &result)

	num, _ := strconv.Atoi(result.PurePhoneNumber)
	user.Phone = num
	db := databases.Connect()
	err := db.Table("user").Where("user_id=?",user_id).Update(&user).Error

	if err == nil{
		c.JSON(200,gin.H{
			"code":200,
			"msg":"ok",
		})
	}else{
		c.JSON(200,gin.H{
			"code":400,
			"msg":"error",
		})
	}
}

type Config struct{
	Config_value string
}

func GetAd(c *gin.Context){
	var config Config
	db := databases.Connect()
	err := db.Table("config").Where("config_id=?",13).First(&config).Error
	if err==nil{
		c.JSON(200,gin.H{
			"code":200,
			"content":config.Config_value,
		})
	}else{
		c.JSON(200,gin.H{
			"code":400,
		})
	}
}

//商户结构体
type Merchant2 struct {
	Merchant_id int
	Name        string
	Img1        string
	Img2        string
	Longitude   float32
	Latitude    float32
	Address     string
}

type Order struct{
	Order_id int
	Merchant_id int
	Status int
	Stars int
}

func GetDetail(c *gin.Context){
	var merchant Merchant2
	var order[] Order
	var count int
	stars := 0
	id := c.Request.FormValue("id")
	id2 ,_ := strconv.Atoi(id)

	db := databases.Connect()
	db.Table("merchant").Where("merchant_id=?",id2).First(&merchant)

	db.Table("order").Where("merchant_id=? AND status = ?",id2,1).Find(&order)
	db.Table("order").Where("merchant_id=? AND status = ?",id2,1).Count(&count)

	for key , _ :=range order{
		stars += order[key].Stars
	}

	if len(order)!=0{
		stars = stars/len(order)
		stars := int(math.Floor(float64(stars)))
		fmt.Println(stars)
	}

	c.JSON(200,gin.H{"code":200,"data":merchant,"count":count,"stars":stars})
}

//商户结构体
type Merchant3 struct {
	Merchant_id int
	Name        string
	Img1        string
	Img2        string
	Longitude   float32
	Latitude    float32
	Address     string
	Star int
	Count int
	Distant float32
}

func GetMerchantList(c *gin.Context){
	var merchant[] Merchant3
	var order[] Order

	curPage := c.Request.FormValue("curPage")
	page , _:= strconv.Atoi(curPage)
	start := (page-1)*8
	db := databases.Connect()
	db.Table("merchant").Limit(8).Offset(start).Find(&merchant).Scan(&merchant)

	for key,value := range merchant{
		db.Table("order").Where("merchant_id=? AND status = ?",value.Merchant_id,1).Find(&order)
		db.Table("order").Where("merchant_id=? AND status = ?",value.Merchant_id,1).Count(&merchant[key].Count)

		for key2 , _ :=range order{
			merchant[key].Star += order[key2].Stars
		}

		if len(order)!=0{
			merchant[key].Star = merchant[key].Star/len(order)
			merchant[key].Star = int(math.Floor(float64(merchant[key].Star)))
			fmt.Println(merchant[key].Star)
		}
	}
	c.JSON(200,gin.H{"code":200,"data":merchant})
}

type Service struct{
	Service_id int
	Service_name string
	Icon string
	Origin_price float32
	Now_price float32
	Pid int
	Is_sale string
	Checked bool
}

func GetServiceList(c *gin.Context){
	var service[] Service

	pid := c.Request.FormValue("pid")
	pid2, _ := strconv.Atoi(pid)

	db := databases.Connect()
	db.Table("service").Where("pid=?",pid2).Find(&service).Scan(&service)

	c.JSON(200,gin.H{"code":200,"data":service})
}

func GetPayPreview(c *gin.Context){
	wechatpay.GetInit(c)
}

//订单结构体
type Order2 struct{
	Order_id int
	Merchant_id int
	Service_id string
	User_id int
	Price float32
	Discount int
	Status int
	Services [] Service2
}

type Service2 struct{
	Service_id int
	Service_name string
	Now_price float32
}

func GetOrderList(c *gin.Context){
	var order Order2

	id, _ := strconv.Atoi(c.Request.FormValue("user_id"))

	db := databases.Connect()
	db.Table("order").Where("user_id=?",id).Find(&order).Scan(&order)

	c.JSON(200,gin.H{"code":200,"data":order})
}