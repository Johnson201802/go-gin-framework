package wechatpay

import (
	"fmt"
	"github.com/objcoding/wxpay"
)

func NewClient(){
	// 创建支付账户
	account := wxpay.NewAccount("wxb7475adc9f5980d2", "1601340050", "H6OSunp4pNZuMu7EA0IK5ayIp3mKAqsL", false)

	// 新建微信支付客户端
	client := wxpay.NewClient(account)

	// 设置证书
	//account.SetCertData("证书地址")

	// 设置http请求超时时间
	client.SetHttpConnectTimeoutMs(2000)

	// 设置http读取信息流超时时间
	client.SetHttpReadTimeoutMs(1000)

	// 更改签名类型
	client.SetSignType("MD5")

	// 统一下单
	params := make(wxpay.Params)
	params.SetString("body", "test").
		SetString("out_trade_no", "436577857").
		SetInt64("total_fee", 1).
		SetString("spbill_create_ip", "127.0.0.1").
		SetString("notify_url", "http://notify.objcoding.com/notify").
		SetString("trade_type", "APP")
	p, _ := client.UnifiedOrder(params)
	fmt.Println(p)
}