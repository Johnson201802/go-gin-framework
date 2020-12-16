package v1

import (
	"awesomeProject/databases"
	"awesomeProject/extend/wechat"
	"awesomeProject/extend/wechatpay"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/iGoogle-ink/gotil"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Merchant struct{
	Merchant_id int
	Longitude string
	Latitude string
}

func GetMerchant2(c *gin.Context){
	var merchant[] Merchant

	db := databases.Connect()
	db.Table("merchant").Where("status=?","1").Find(&merchant)
	c.JSON(200,gin.H{"data":merchant})
	defer db.Close()
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
	res , _ := wechat.WXLogin(code,"wxb5fb97bbf613fa55","392ff804414225de5f4c7b21ddf0afc6")
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
	defer db.Close()

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
	Stars int
	Sales int
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
	//var count int
	//stars := 0
	id := c.Request.FormValue("id")
	id2 ,_ := strconv.Atoi(id)

	db := databases.Connect()
	db.Table("merchant").Where("merchant_id=?",id2).First(&merchant)

	db.Table("order").Where("merchant_id=? AND status = ?",id2,1).Find(&order)
	//db.Table("order").Where("merchant_id=? AND status = ?",id2,1).Count(&count)

	//for key , _ :=range order{
	//	stars += order[key].Stars
	//}
	//
	//if len(order)!=0{
	//	stars = stars/len(order)
	//	stars := int(math.Floor(float64(stars)))
	//	fmt.Println(stars)
	//}

	c.JSON(200,gin.H{"code":200,"data":merchant})
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
	Stars int
	Sales int
	Distant float32
	Status string
}

func GetMerchantList(c *gin.Context){
	var merchant[] Merchant3
	//var order[] Order

	curPage := c.Request.FormValue("curPage")
	page , _:= strconv.Atoi(curPage)
	start := (page-1)*8
	db := databases.Connect()
	db.Table("merchant").Where("status = ?","1").Limit(8).Offset(start).Find(&merchant).Scan(&merchant)

	//for key,value := range merchant{
	//	db.Table("order").Where("merchant_id=? AND status = ?",value.Merchant_id,1).Find(&order)
	//	db.Table("order").Where("merchant_id=? AND status = ?",value.Merchant_id,1).Count(&merchant[key].Count)
	//
	//	for key2 , _ :=range order{
	//		merchant[key].Star += order[key2].Stars
	//	}
	//
	//	if len(order)!=0{
	//		merchant[key].Star = merchant[key].Star/len(order)
	//		merchant[key].Star = int(math.Floor(float64(merchant[key].Star)))
	//		fmt.Println(merchant[key].Star)
	//	}
	//}
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

func GetPayPreview2(c *gin.Context){
	wechatpay.GetInit2(c)
}

//订单结构体
type Order2 struct{
	Id string
	Merchant_id int
	Service_id string
	User_id int
	Price float32
	Discount int
	Status int
	Services [] Service2
	Name string
}

type Service2 struct{
	Service_id int
	Service_name string
	Origin_price int
}

func GetOrderList(c *gin.Context){
	var order[] Order2
	var service[] Service2

	id, _ := strconv.Atoi(c.Request.FormValue("user_id"))
	curPage := c.Request.FormValue("curPage")
	page , _:= strconv.Atoi(curPage)
	start := (page-1)*6

	db := databases.Connect()
	db.Debug().Table("order").Select("order.Id, order.merchant_id, order.service_id, order.user_id, order.price, order.status, order.discount, merchant.name").Joins("join merchant on merchant.merchant_id = order.merchant_id").Where("order.user_id = ?",id).Offset(start).Limit(6).Find(&order).Scan(&order)

	for key , item := range order{
		arr := strings.Split(item.Service_id,",")
		arr2 := []int{}
		for _ , item2 := range arr{
			tt , _:= strconv.Atoi(item2)
			arr2 = append(arr2,tt)
		}
		db.Debug().Table("service").Where("service_id in (?)",arr2).Find(&service).Scan(&service)
		order[key].Services = service
	}

	c.JSON(200,gin.H{"code":200,"data":order})
}

//会员卡结构体
type Card struct{
	Card_id int
	Title string
	Price int
	Desc string
	Days int
	Status string
}

func GetCardList(c *gin.Context){
	var card[] Card

	db := databases.Connect()
	db.Table("card").Where("status = ?",1).Find(&card).Scan(&card)

	c.JSON(200,gin.H{"code":200,"data":card})
}

type User2 struct{
	User_id int
	Expire_vip_time int64
	Discount int
}

func GetVipInfo(c *gin.Context){
	var user User2

	id, _ := strconv.Atoi(c.Request.FormValue("user_id"))
	user.User_id = id

	db := databases.Connect()
	db.Table("user").Where("user_id = ?",id).First(&user).Scan(&user)

	datetime := time.Unix(user.Expire_vip_time, 0).Format("2006-01-02")

	c.JSON(200,gin.H{"code":200,"data":user,"datetime":datetime})
}

func Notify(c *gin.Context){
	wechatpay.Notify(c)
}

func Notify2(c *gin.Context){
	wechatpay.Notify2(c)
}

type User3 struct{
	User_id int
	Expire_vip_time int64
	Is_vip string
	Discount int
}

type Order22 struct{
	Id string
	Merchant_id int
	Service_id string
	User_id int
	Price int
	Discount int
	Status int
	Time int64
}

type Order33 struct{
	Id string
	Stars int
	Conment string
}

func MakeOrder(c *gin.Context){
	var order Order22

	tt := fmt.Sprintf("%06v", rand.Int31n(100000000))
	order.Id = tt + gotil.GetRandomString(10)
	order.Merchant_id , _ = strconv.Atoi(c.Request.FormValue("merchant_id"))
	order.User_id , _ = strconv.Atoi(c.Request.FormValue("user_id"))
	order.Service_id  = c.Request.FormValue("service_id")
	order.Price ,_ = strconv.Atoi(c.Request.FormValue("price"))
	order.Discount  , _= strconv.Atoi(c.Request.FormValue("discount"))
	order.Time = time.Now().Unix()

	db := databases.Connect()
	if db.Table("order").Create(&order).Error == nil {
		c.JSON(200,gin.H{
			"code" : 200,
			"order_id" : order.Id,
			"user_id" : order.User_id,
		})
	}else{
		c.JSON(200,gin.H{
			"code" : 300,
		})
	}

}

func MakeComment(c *gin.Context){
	var order Order33

	order.Id = c.Request.FormValue("order_id")
	order.Conment = c.Request.FormValue("comment")
	order.Stars ,_ = strconv.Atoi(c.Request.FormValue("count"))

	fmt.Println(order.Id)
	fmt.Println(order.Conment)
	fmt.Println(order.Stars)

	db := databases.Connect()
	tt := db.Debug().Table("order").Where("Id=?",order.Id).Updates(&order).Error
	if tt == nil{
		c.JSON(200,gin.H{
			"code" : 200,
		})
	}else{
		c.JSON(200,gin.H{
			"code" : 300,
		})
	}
}