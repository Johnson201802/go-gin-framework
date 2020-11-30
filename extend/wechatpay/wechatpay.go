package wechatpay

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/iGoogle-ink/gopay"
	_ "github.com/iGoogle-ink/gopay"
	"github.com/iGoogle-ink/gopay/wechat"
	"github.com/iGoogle-ink/gotil"
	"github.com/iGoogle-ink/gotil/xlog"
	"strconv"
	"time"
)

func GetInit(c *gin.Context){
	// 初始化微信客户端
	//    appId：应用ID
	//    mchId：商户ID
	//    apiKey：API秘钥值
	//    isProd：是否是正式环境
	client := wechat.NewClient("wxb7475adc9f5980d2", "1601340050", "H6OSunp4pNZuMu7EA0IK5ayIp3mKAqsL", true)

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

	// 初始化 BodyMap
	bm := make(gopay.BodyMap)
	bm.Set("nonce_str", gotil.GetRandomString(32))
	bm.Set("body", "小程序测试支付")
	bm.Set("out_trade_no", "56789876545678")
	bm.Set("total_fee", 1)
	bm.Set("spbill_create_ip", "127.0.0.1")
	bm.Set("notify_url", "http://www.gopay.ink")
	bm.Set("trade_type", wechat.TradeType_Mini)
	bm.Set("device_info", "WEB")
	bm.Set("sign_type", wechat.SignType_MD5)
	bm.Set("openid", "oV3TJ5bH0w7ajL3SzOtbDDP_HmDo")

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
	sign := wechat.GetParamSign("wxb7475adc9f5980d2", "1601340050", "H6OSunp4pNZuMu7EA0IK5ayIp3mKAqsL", bm)
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
	paySign := wechat.GetMiniPaySign("wxb7475adc9f5980d2", wxRsp.NonceStr, packages, wechat.SignType_MD5, timeStamp, "H6OSunp4pNZuMu7EA0IK5ayIp3mKAqsL")

	fmt.Println(paySign)


	c.JSON(200,gin.H{"timeStamp":timeStamp,"nonceStr":wxRsp.NonceStr,"package":packages,"paySign":paySign})
}

//异步通知回调
func Notify(c *gin.Context){
	// 初始化微信客户端
	//    appId：应用ID
	//    mchId：商户ID
	//    apiKey：API秘钥值
	//    isProd：是否是正式环境
	client := wechat.NewClient("wxb7475adc9f5980d2", "1601340050", "H6OSunp4pNZuMu7EA0IK5ayIp3mKAqsL", true)

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

	rsp := new(wechat.NotifyResponse)
	// 解析参数
	notifyReq, err := wechat.ParseRefundNotify(c.Request)
	if err != nil {
		xlog.Debug("err:", err)
	}
	xlog.Debug("notifyReq:", *notifyReq)
	// 退款通知无sign，不用验签

	// 解密退款异步通知的加密数据
	refundNotify, err := wechat.DecryptRefundNotifyReqInfo(notifyReq.ReqInfo, "GFDS8j98rewnmgl45wHTt980jg543abc")
	if err != nil {
		xlog.Debug("err:", err)
	}
	xlog.Debug("refundNotify:", *refundNotify)

	// 或者

	bodyMap, err := wechat.ParseNotifyToBodyMap(c.Request)
	if err != nil {
		xlog.Debug("err:", err)
	}
	xlog.Debug("bodyMap:", bodyMap)

	// 解密退款异步通知的加密数据
	refundNotify2, err := wechat.DecryptRefundNotifyReqInfo(bodyMap.Get("req_info"), "GFDS8j98rewnmgl45wHTt980jg543abc")
	if err != nil {
		xlog.Debug("err:", err)
	}
	xlog.Debug("refundNotify:", *refundNotify2)

	// 返回微信
	rsp.ReturnCode = gopay.SUCCESS
	rsp.ReturnMsg = "OK"

	fmt.Println(rsp.ToXmlString())
}

//退款
func Refund(c *gin.Context){
	// 初始化微信客户端
	//    appId：应用ID
	//    mchId：商户ID
	//    apiKey：API秘钥值
	//    isProd：是否是正式环境
	client := wechat.NewClient("wxb7475adc9f5980d2", "1601340050", "H6OSunp4pNZuMu7EA0IK5ayIp3mKAqsL", false)

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


