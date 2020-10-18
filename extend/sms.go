package sms

import (
	"fmt"
	"github.com/zboyco/gosms"
)

//单发短信
func SingleSms(){
	// 创建Sender
	sender := &gosms.QSender{
		AppID:  "",                       // appid
		AppKey: "", // appkey
	}

	// 发送短信
	res, err := sender.SingleSend(
		"约工宝",     // 短信签名，此处应填写审核通过的签名内容，非签名 ID，如果使用默认签名，该字段填 ""
		86,            // 国家号
		"", // 手机号
		598800,         // 短信正文ID
		"123456987678567",      // 参数1
		//"5",           // 参数2，后面可添加多个参数
	)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}

//统一国家码群发短信
func mutilSMS(){
	// 创建Sender
	sender := &gosms.QSender{
		AppID:  "1234567890",                       // appid
		AppKey: "12345678901234567890123456789000", // appkey
	}

	// 统一国家码群发短信
	res, err := sender.MultiSend(
		"短信签名", // 短信签名，此处应填写审核通过的签名内容，非签名 ID，如果使用默认签名，该字段填 ""
		86, // 国家号
		[]string{
			"13800000000", // 手机号
			"13800000000", // 手机号
			"13800000000", // 手机号
		},
		10000,    // 短信正文ID
		"123456", // 参数1
		"5",      // 参数2，后面可添加多个参数
	)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}

//短信结构体
type Telphone struct{
	Phone string
	CC int
}

//各自国家码群发短信
//func DifferCoutrySend(){
//	// 创建Sender
//	sender := &gosms.QSender{
//		AppID:  "1234567890",                       // appid
//		AppKey: "12345678901234567890123456789000", // appkey
//	}
//
//	var tel Telphone
//	// 各自国家码群发短信
//	res, err := sender.MultiSendEachCC(
//		"短信签名", // 短信签名，此处应填写审核通过的签名内容，非签名 ID，如果使用默认签名，该字段填 ""
//		tel{
//			tel{
//				Phone: "13800000000", // 手机号
//				CC:    86,            // 国家号
//			},
//			tel{
//				Phone: "13800000000", // 手机号
//				CC:    86,            // 国家号
//			},
//		},
//		10000,    // 短信正文ID
//		"123456", // 参数1
//		"5",      // 参数2，后面可添加多个参数
//	)
//	if err != nil {
//		fmt.Println(err)
//	} else {
//		fmt.Println(res)
//	}
//}

//拉取单个号码短信下发状态
func PullSingleSMSStatus(){
	// 创建Sender
	sender := &gosms.QSender{
		AppID:  "1234567890",                       // appid
		AppKey: "12345678901234567890123456789000", // appkey
	}

	// 拉取下发状态
	res, err := sender.PullSingleStatus(
		86,                    // 国家码
		"13800000000",         // 号码
		"2019-04-01 00:00:00", // 开始日期，注意格式
		"2019-04-03 00:00:00", // 结束日期，注意格式
		100,                   // 拉取最大条数，最大拉取100条
	)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}

//拉取多个号码短信下发状态
func PullMutilSMSStatus(){
	// 创建Sender
	sender := &gosms.QSender{
		AppID:  "1234567890",                       // appid
		AppKey: "12345678901234567890123456789000", // appkey
	}

	// 拉取下发状态 此功能需要联系 qcloud sms helper 开通。
	res, err := sender.PullStatus(100) // 最大拉取100条
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}

//拉取单个号码短信回复
func PullSingleRecover(){
	// 创建Sender
	sender := &gosms.QSender{
		AppID:  "1234567890",                       // appid
		AppKey: "12345678901234567890123456789000", // appkey
	}

	// 拉取短信回复
	res, err := sender.PullSingleReply(
		86,                    // 国家码
		"13800000000",         // 号码
		"2019-04-01 00:00:00", // 开始日期，注意格式
		"2019-04-03 00:00:00", // 结束日期，注意格式
		100,                   // 拉取最大条数，最大拉取100条
	)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}

//拉取多条短信回复
func PullMutilRecover(){
	// 创建Sender
	sender := &gosms.QSender{
		AppID:  "1234567890",                       // appid
		AppKey: "12345678901234567890123456789000", // appkey
	}

	// 拉取短信回复 此功能需要联系 qcloud sms helper 开通。
	res, err := sender.PullReply(100) // 最大拉取100条
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}

//发送语音验证码
func SendVoiceCaptch(){
	// 创建Sender
	sender := &gosms.QSender{
		AppID:  "1234567890",                       // appid
		AppKey: "12345678901234567890123456789000", // appkey
	}

	// 发送语音验证码
	res, err := sender.VoiceSendCaptcha(
		"13800000000", // 手机号
		2,             // 播报次数,最多3次
		"123456",      // 要播报的验证码，仅支持数字（string类型）
	)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}

//发送语音通知
func SendVoiceNotice(){
	// 创建Sender
	sender := &gosms.QSender{
		AppID:  "1400360532",                       // appid
		AppKey: "f215ace273cdf6da032a1d05921780ef", // appkey
	}

	// 发送语音通知
	res, err := sender.VoiceSendPrompt(
		"17621953521", // 手机号
		2,             // 播报次数,最多3次
		"我是测试语音通知的", // 要播报的语音文本类容
	)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}