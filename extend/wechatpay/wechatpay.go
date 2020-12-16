package wechatpay

import (
	"awesomeProject/databases"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/iGoogle-ink/gopay"
	_ "github.com/iGoogle-ink/gopay"
	"github.com/iGoogle-ink/gopay/wechat"
	"github.com/iGoogle-ink/gotil"
	"github.com/iGoogle-ink/gotil/xlog"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

//定义用户结构体
type User struct{
	User_id int
	Open_id string
}

func GetInit(c *gin.Context){
	// 初始化微信客户端
	//    appId：应用ID
	//    mchId：商户ID
	//    apiKey：API秘钥值
	//    isProd：是否是正式环境
	client := wechat.NewClient("22", "22", "22", true)

	// 打开Debug开关，输出日志，默认关闭
	client.DebugSwitch = gopay.DebugOn

	// 设置国家：不设置默认 中国国内
	//    wechat.China：中国国内
	//    wechat.China2：中国国内备用
	//    wechat.SoutheastAsia：东南亚
	//    wechat.Other：其他国家
	client.SetCountry(wechat.China)

	// 添加微信证书 Path 路径
	//    certFilePath：apiclient_cert.pem 路径
	//    keyFilePath：apiclient_key.pem 路径
	//    pkcs12FilePath：apiclient_cert.p12 路径
	//    返回err
	//client.AddCertFilePath()

	mouths :=c.Request.FormValue("mouths")
	user_id  := c.Request.FormValue("user_id")
	price , _ := strconv.Atoi(c.Request.FormValue("price"))

	var user User

	db := databases.Connect()
	db.Table("user").Where("user_id = ?",user_id).Find(&user).Scan(&user)

	// 初始化 BodyMap
	bm := make(gopay.BodyMap)
	bm.Set("nonce_str", gotil.GetRandomString(32))
	bm.Set("body", "小程序会员支付")
	bm.Set("out_trade_no", mouths+"-"+gotil.GetRandomString(10))
	bm.Set("total_fee", price*100)
	bm.Set("spbill_create_ip", "127.0.0.1")
	bm.Set("notify_url", "https://api.piduopi.com/api/v1/notify")
	bm.Set("trade_type", wechat.TradeType_Mini)
	bm.Set("device_info", "WEB")
	bm.Set("sign_type", wechat.SignType_MD5)
	bm.Set("openid", user.Open_id)


	// 嵌套json格式数据（例如：H5支付的 scene_info 参数）
	//h5Info := make(map[string]string)
	//h5Info["type"] = "Wap"
	//h5Info["wap_url"] = "http://www.gopay.ink"
	//h5Info["wap_name"] = "H5测试支付"
	//
	//sceneInfo := make(map[string]map[string]string)
	//sceneInfo["h5_info"] = h5Info
	//
	//bm.Set("scene_info", sceneInfo)

	// 参数 sign ，可单独生成赋值到BodyMap中；也可不传sign参数，client内部会自动获取
	// 如需单独赋值 sign 参数，需通过下面方法，最后获取sign值并在最后赋值此参数
	sign := wechat.GetParamSign("22", "22", "22", bm)
	// sign, _ := wechat.GetSanBoxParamSign("wxdaa2ab9ef87b5497", mchId, apiKey, body)
	bm.Set("sign", sign)

	wxRsp, _ := client.UnifiedOrder(bm)

	fmt.Println(wxRsp)

	// ====微信小程序 paySign====
	timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
	packages := "prepay_id=" + wxRsp.PrepayId   // 此处的 wxRsp.PrepayId ,统一下单成功后得到
	// 获取微信小程序支付的 paySign
	//    appId：AppID
	//    nonceStr：随机字符串
	//    packages：统一下单成功后拼接得到的值
	//    signType：签名方式，务必与统一下单时用的签名方式一致
	//    timeStamp：时间
	//    apiKey：API秘钥值
	paySign := wechat.GetMiniPaySign("22", wxRsp.NonceStr, packages, wechat.SignType_MD5, timeStamp, "22")

	fmt.Println(paySign)


	c.JSON(200,gin.H{"timeStamp":timeStamp,"nonceStr":wxRsp.NonceStr,"package":packages,"paySign":paySign})
	defer db.Close()
}

type Order struct{
	Id string
	Merchant_id int
	Service_id string
	User_id int
	Price int
	Discount int
	Status int
}

func GetInit2(c *gin.Context){
	// 初始化微信客户端
	//    appId：应用ID
	//    mchId：商户ID
	//    apiKey：API秘钥值
	//    isProd：是否是正式环境
	client := wechat.NewClient("22", "22", "22", true)

	// 打开Debug开关，输出日志，默认关闭
	client.DebugSwitch = gopay.DebugOn

	// 设置国家：不设置默认 中国国内
	//    wechat.China：中国国内
	//    wechat.China2：中国国内备用
	//    wechat.SoutheastAsia：东南亚
	//    wechat.Other：其他国家
	client.SetCountry(wechat.China)

	// 添加微信证书 Path 路径
	//    certFilePath：apiclient_cert.pem 路径
	//    keyFilePath：apiclient_key.pem 路径
	//    pkcs12FilePath：apiclient_cert.p12 路径
	//    返回err
	//client.AddCertFilePath()

	var order Order

	tt := fmt.Sprintf("%06v", rand.Int31n(100000000))
	order.Id = tt + gotil.GetRandomString(10)
	order.Merchant_id , _ = strconv.Atoi(c.Request.FormValue("merchant_id"))
	order.User_id , _ = strconv.Atoi(c.Request.FormValue("user_id"))
	order.Service_id  = c.Request.FormValue("service_id")
	order.Price ,_ = strconv.Atoi(c.Request.FormValue("price"))
	order.Discount  , _= strconv.Atoi(c.Request.FormValue("discount"))

	db := databases.Connect()
	db.Table("order").Create(&order)

	fmt.Println(order.Id)

	var user User

	db.Table("user").Where("user_id = ?",order.User_id).Find(&user).Scan(&user)

	// 初始化 BodyMap
	bm := make(gopay.BodyMap)
	bm.Set("nonce_str", gotil.GetRandomString(32))
	bm.Set("body", "小程序服务支付")
	bm.Set("out_trade_no", order.Id)
	bm.Set("total_fee", order.Price*100) //order.Price*100
	bm.Set("spbill_create_ip", "127.0.0.1")
	bm.Set("notify_url", "https://api.piduopi.com/api/v1/notify2")
	bm.Set("trade_type", wechat.TradeType_Mini)
	bm.Set("device_info", "WEB")
	bm.Set("sign_type", wechat.SignType_MD5)
	bm.Set("openid", user.Open_id)


	// 嵌套json格式数据（例如：H5支付的 scene_info 参数）
	//h5Info := make(map[string]string)
	//h5Info["type"] = "Wap"
	//h5Info["wap_url"] = "http://www.gopay.ink"
	//h5Info["wap_name"] = "H5测试支付"
	//
	//sceneInfo := make(map[string]map[string]string)
	//sceneInfo["h5_info"] = h5Info
	//
	//bm.Set("scene_info", sceneInfo)

	// 参数 sign ，可单独生成赋值到BodyMap中；也可不传sign参数，client内部会自动获取
	// 如需单独赋值 sign 参数，需通过下面方法，最后获取sign值并在最后赋值此参数
	sign := wechat.GetParamSign("22", "22", "22", bm)
	// sign, _ := wechat.GetSanBoxParamSign("wxdaa2ab9ef87b5497", mchId, apiKey, body)
	bm.Set("sign", sign)

	wxRsp, _ := client.UnifiedOrder(bm)

	fmt.Println(wxRsp)

	// ====微信小程序 paySign====
	timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
	packages := "prepay_id=" + wxRsp.PrepayId   // 此处的 wxRsp.PrepayId ,统一下单成功后得到
	// 获取微信小程序支付的 paySign
	//    appId：AppID
	//    nonceStr：随机字符串
	//    packages：统一下单成功后拼接得到的值
	//    signType：签名方式，务必与统一下单时用的签名方式一致
	//    timeStamp：时间
	//    apiKey：API秘钥值
	paySign := wechat.GetMiniPaySign("22", wxRsp.NonceStr, packages, wechat.SignType_MD5, timeStamp, "22")

	fmt.Println(paySign)


	c.JSON(200,gin.H{"timeStamp":timeStamp,"nonceStr":wxRsp.NonceStr,"package":packages,"paySign":paySign,"order_id":order.Id})
	defer db.Close()
}

type User3 struct{
	Open_id string
	Expire_vip_time int64
	Is_vip string
	Discount int
}

//异步通知回调
func Notify(c *gin.Context) string{
	var user User3
	// 初始化微信客户端
	//    appId：应用ID
	//    mchId：商户ID
	//    apiKey：API秘钥值
	//    isProd：是否是正式环境
	client := wechat.NewClient("22", "22", "22", true)

	// 打开Debug开关，输出日志，默认关闭
	client.DebugSwitch = gopay.DebugOn

	// 设置国家：不设置默认 中国国内
	//    wechat.China：中国国内
	//    wechat.China2：中国国内备用
	//    wechat.SoutheastAsia：东南亚
	//    wechat.Other：其他国家
	client.SetCountry(wechat.China)

	// 添加微信证书 Path 路径 返回err
	//err := client.AddCertFilePath("./apiclient_cert.pem","./apiclient_key.pem","./apiclient_cert.p12")
	// 解析notify参数、验签、返回数据到微信
	   rsp := new(wechat.NotifyResponse)

		// 解析参数
		bodyMap, err := wechat.ParseNotifyToBodyMap(c.Request)
		if err != nil {
		xlog.Debug("err:", err)
	}
		db := databases.Connect()
		xlog.Debug("bodyMap:", bodyMap)
		arr := strings.Split(bodyMap.Get("out_trade_no"), "-")
		mounths ,_ := strconv.Atoi(arr[0])
		user.Open_id = bodyMap.Get("openid")

		ok, err := wechat.VerifySign("83748374873894tyeruigyfhdsefuyHU", wechat.SignType_MD5, bodyMap)
		if err != nil {
			xlog.Debug("err:", err)
		}else{

			db.Table("user").Where("open_id=?",user.Open_id).First(&user).Scan(&user)

			if user.Is_vip != "1"{
				time2 := time.Now().AddDate(0,mounths,0).Unix()
				user.Expire_vip_time = time2
				user.Is_vip = "1"

				db.Table("user").Where("open_id=?",user.Open_id).Update(&user)
			}
		}
		xlog.Debug("微信验签是否通过:", ok)

		rsp.ReturnCode = gopay.SUCCESS
		rsp.ReturnMsg = "OK"
		defer db.Close()
		return rsp.ToXmlString()

}

//异步通知回调2
func Notify2(c *gin.Context) string{
	var user User3
	var order Order
	// 初始化微信客户端
	//    appId：应用ID
	//    mchId：商户ID
	//    apiKey：API秘钥值
	//    isProd：是否是正式环境
	client := wechat.NewClient("22", "22", "22", true)

	// 打开Debug开关，输出日志，默认关闭
	client.DebugSwitch = gopay.DebugOn

	// 设置国家：不设置默认 中国国内
	//    wechat.China：中国国内
	//    wechat.China2：中国国内备用
	//    wechat.SoutheastAsia：东南亚
	//    wechat.Other：其他国家
	client.SetCountry(wechat.China)

	// 添加微信证书 Path 路径 返回err
	//err := client.AddCertFilePath("./apiclient_cert.pem","./apiclient_key.pem","./apiclient_cert.p12")
	// 解析notify参数、验签、返回数据到微信
	rsp := new(wechat.NotifyResponse)

	// 解析参数
	bodyMap, err := wechat.ParseNotifyToBodyMap(c.Request)
	if err != nil {
		xlog.Debug("err:", err)
	}

	db := databases.Connect()
	xlog.Debug("bodyMap:", bodyMap)
	out_trade_no := bodyMap.Get("out_trade_no")
	openid := bodyMap.Get("openid")

	ok, err := wechat.VerifySign("83748374873894tyeruigyfhdsefuyHU", wechat.SignType_MD5, bodyMap)
	if err != nil {
		xlog.Debug("err:", err)
	}else{

		res := db.Table("order").Where("Id=?",out_trade_no).First(&order).Scan(&order).Error
		res2 := db.Table("user").Where("open_id=?",openid).First(&user).Scan(&user).Error

		if res == nil && res2==nil && order.Status == 0{
			order.Status = 1
			user.Discount = user.Discount + order.Discount
			db.Table("order").Where("Id = ?",out_trade_no).Update(&order)
			db.Debug().Table("user").Where("open_id = ?",openid).Update(&user)
		}
	}
	xlog.Debug("微信验签是否通过:", ok)

	rsp.ReturnCode = gopay.SUCCESS
	rsp.ReturnMsg = "OK"
	defer db.Close()
	return rsp.ToXmlString()
}

//退款
func Refund(c *gin.Context){
	// 初始化微信客户端
	//    appId：应用ID
	//    mchId：商户ID
	//    apiKey：API秘钥值
	//    isProd：是否是正式环境
	client := wechat.NewClient("22", "22", "22", false)

	// 初始化参数结构体
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", "SdZBAqJHBQGKVwb7aMR2mUwC588NG2Sd")
	bm.Set("nonce_str", gotil.GetRandomString(32))
	bm.Set("sign_type", wechat.SignType_MD5)
	s := gotil.GetRandomString(64)
	xlog.Debug("out_refund_no:", s)
	bm.Set("out_refund_no", s)
	bm.Set("total_fee", 1)
	bm.Set("refund_fee", 1)
	bm.Set("notify_url", "https://www.gopay.ink")

	//请求申请退款（沙箱环境下，证书路径参数可传空）
	//    body：参数Body
	//    certFilePath：cert证书路径
	//    keyFilePath：Key证书路径
	//    pkcs12FilePath：p12证书路径
	wxRsp, resBm, err := client.Refund(bm, "iguiyu_cert/apiclient_cert.pem", "iguiyu_cert/apiclient_key.pem", "iguiyu_cert/apiclient_cert.p12")
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("wxRsp：", wxRsp)
	xlog.Debug("resBm:", resBm)
}


