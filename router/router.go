package router

import (
	"awesomeProject/apis"
	v1 "awesomeProject/apis/v1"
	"github.com/gin-gonic/gin"
)

//管理端
//跨域中间件配置
func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("origin")
		if len(origin) == 0 {
			origin = c.Request.Header.Get("Origin")
		}
		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With,x-token")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, DELETE")
		c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}
		c.Next()
	}
}

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors())
	router.POST("/adminLogin", apis.IndexApi)

	//路由组没写加上JWT鉴权
	//系统设置
	router.GET("/getInfo", apis.GetInfo)
	router.GET("/serverInfo", apis.GetServerInfo)
	router.GET("/getBaseconfig", apis.GetBaseconfig)
	router.GET("/getMsmConfig", apis.GetMsmConfig)
	router.GET("/getMchConfig", apis.GetMchConfig)
	router.POST("/imgUploads", apis.ImgUploads)
	router.POST("/saveConfigBase", apis.SaveConfigBase)
	router.POST("/saveConfigSms", apis.SaveConfigSms)
	router.POST("/saveConfigMch", apis.SaveConfigMch)

	//权限管理
	router.GET("/getAuthList", apis.GetAuthList)
	router.DELETE("/delRule", apis.DelRule)
	router.POST("/createRule", apis.CreateRule)
	router.POST("/updateRule", apis.UpdateRule)

	//管理员管理
	router.GET("/getAdminList",apis.GetAdminList)
	router.GET("/getGroupList",apis.GetGroupList)
	router.POST("/createAdmin",apis.CreateAdmin)
	router.DELETE("/delAdmin",apis.DelAdmin)
	router.POST("/updateAdmin",apis.UpdateAdmin)

	//用户组管理
	router.GET("/getRoleList",apis.GetRoleList)
	router.GET("/getRuleList",apis.GetRuleList)
	router.POST("/addRole",apis.AddRole)
	router.DELETE("/delRole",apis.DelRole)
	router.GET("/getOneRole",apis.GetOneRole)
	router.POST("/updateRole",apis.UpdateRole)

	//用户管理
	router.GET("/fetchUserList",apis.FetchUserList)
	router.DELETE("/delUser",apis.DelUser)

	//文章管理
	router.GET("/fetchArticleList",apis.FetchArticleList)
	router.DELETE("/delArticle",apis.DelArticle)
	router.POST("/createArticle",apis.CreateArticle)
	router.POST("/updateArticle",apis.UpdateArticle)
	router.GET("/fetchOneArticle",apis.FetchOneArticle)

	//商家管理
	router.GET("/fetchMerchantList",apis.FetchMerchantList)
	router.POST("/delMerchant",apis.DelMerchant)
	router.POST("/createMerchant",apis.CreateMerchant)
	router.GET("/fetchOneMerchant",apis.FetchOneMerchant)
	router.POST("/updateMerchant",apis.UpdateMerchant)
	router.POST("/changeMerchantStatus",apis.ChangeMerchantStatus)

	//订单管理
	router.GET("/fetchOrderList",apis.FetchOrderList)
	router.POST("/delOrder",apis.DelOrder)

	//服务管理
	router.GET("/fetchServicesList",apis.FetchServicesList)
	router.POST("/changeStatus",apis.ChangeStatus)
	router.POST("/createService",apis.CreateService)
	router.GET("/getMerchant",apis.GetMerchant)
	router.POST("/updateService",apis.UpdateService)
	router.DELETE("/delPresent",apis.DelPresent)

	//是否有新信息
	router.GET("/getNewMsg",apis.GetNewMsg)
	router.GET("/getQuestionList",apis.GetQuestionList)
	router.POST("/setRead",apis.SetRead)
	router.POST("/articleCreate",apis.ArticleCreate)

	//会员卡后管理
	router.GET("/getCardList",apis.GetCardList)
	router.GET("/changeStatus2",apis.ChangeStatus2)
	router.POST("/updateCard",apis.UpdateCard)

	//首页数据请求
	router.GET("/getMerchantList55",apis.GetMerchantList55)
	router.GET("/getAllInfo",apis.GetAllInfo)

	//小程序端fetchOneMerchant
	router.GET("/api/v1/getMerchant",v1.GetMerchant2)
	router.GET("/api/v1/getOpenid",v1.GetOpenid)
	router.GET("/api/v1/getPhoneNumber",v1.GetPhoneNumber)
	router.GET("/api/v1/getAd",v1.GetAd)
	router.GET("/api/v1/getDetail",v1.GetDetail)
	router.GET("/api/v1/getMerchantList",v1.GetMerchantList)
	router.GET("/api/v1/getServiceList",v1.GetServiceList)
	router.GET("/api/v1/getPayPreview",v1.GetPayPreview)
	router.GET("/api/v1/getPayPreview2",v1.GetPayPreview2)
	router.GET("/api/v1/makeOrder",v1.MakeOrder)
	router.GET("/api/v1/getOrderList",v1.GetOrderList)
	router.GET("/api/v1/getCardList",v1.GetCardList)
	router.GET("/api/v1/getVipInfo",v1.GetVipInfo)
	router.POST("/api/v1/notify",v1.Notify)
	router.POST("/api/v1/notify2",v1.Notify2)
	router.GET("/api/v1/makeComment",v1.MakeComment)

	return router
}
