package apis

import (
	"awesomeProject/databases"
	"awesomeProject/extend/oss"
	"awesomeProject/extend/sms"
	"awesomeProject/models"
	"awesomeProject/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
)

var roles = [2]string{"admin", "edit"}

func IndexApi(c *gin.Context) {
	name := c.Request.FormValue("username")
	pwd := tools.MD5(c.Request.FormValue("password"))
	var uid2 string
	flag, uid := models.GetAdmin(name, pwd)
	uid2 = strconv.Itoa(uid)

	//生成TOKEN
	token, _ := tools.CreateToken(uid2, "johnson2018092789728376273672364wey")
	red := databases.Connect_redis()
	defer red.Close()
	var err error
	_, err = red.Do("SET", uid, token)
	_, err = red.Do("expire", uid, 7200) //两个小时过期
	if err != nil {
		fmt.Println("set expire error: ", err)
		return
	}
	if flag {
		c.JSON(200, gin.H{
			"code":         200,
			"roles":        roles,
			"name":         name,
			"avatar":       "234567",
			"introduction": "",
			"token":        token,
		})
	} else {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  "",
		})
	}
}

func GetInfo(c *gin.Context) {
	token := c.Request.Header.Get("X-Token")
	uid, _ := tools.ParseToken(token, "johnson2018092789728376273672364wey")
	red := databases.Connect_redis()
	v, _ := red.Do("GET", uid)
	if v == nil {
		c.JSON(200, gin.H{
			"code": 600,
		})
	} else {
		c.JSON(200, gin.H{
			"code":         200,
			"roles":        "admin",
			"name":         "jason",
			"avatar":       "12345",
			"introduction": "234567",
		})
	}
	defer red.Close()
}

func GetServerInfo(c *gin.Context) {
	tools.ServerInfo(c)
}

func GetBaseconfig(c *gin.Context) {
	models.GetBaseconfig(c)
}

func GetMsmConfig(c *gin.Context) {
	models.GetMsmConfig(c)
}

func GetMchConfig(c *gin.Context) {
	models.GetMchConfig(c)
}

func ImgUploads(c *gin.Context) {
	f, err := c.FormFile("file")
	if err != nil {
		c.JSON(200, gin.H{
			"error": err,
		})
		return
	} else {
		err := c.SaveUploadedFile(f, f.Filename)
		if err != nil {
			c.JSON(200, gin.H{
				"error": err,
			})
			return
		} else {
			UploadFile(c, f.Filename)
			err := os.Remove(f.Filename)
			db := databases.Connect()
			db.Table("config").Where("config_id = ?", 5).Update("config_value","https://img-c-jason.oss-accelerate.aliyuncs.com/"+f.Filename)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func SendMSM(c *gin.Context) {
	sms.SingleSms()
}

func GetOss(c *gin.Context) {
	oss.ListFile(c)
}

func UploadFile(c *gin.Context, filename string) {
	oss.UploadFile(c, filename)
}

func GetStorageList(c *gin.Context) {
	oss.GetStorageList(c)
}

func SaveConfigBase(c *gin.Context) {
	models.SaveConfigBase(c)
}

func SaveConfigSms(c *gin.Context){
	models.SaveConfigSms(c)
}

func SaveConfigMch(c *gin.Context){
	models.SaveConfigSms(c)
}

func GetAuthList(c *gin.Context){
	models.GetAuthList(c)
}

func DelRule(c *gin.Context){
	models.DelRule(c)
}

func CreateRule(c *gin.Context){
	models.CreateRule(c)
}

func UpdateRule(c *gin.Context){
	models.UpdateRule(c)
}

func GetAdminList(c *gin.Context){
	models.GetAdminList(c)
}

func GetGroupList(c *gin.Context){
	models.GetGroupList(c)
}

func CreateAdmin(c *gin.Context){
	models.CreateAdmin(c)
}

func DelAdmin(c *gin.Context){
	models.DelAdmin(c)
}

func UpdateAdmin(c *gin.Context){
	models.UpdateAdmin(c)
}

func GetRoleList(c *gin.Context){
	models.GetRoleList(c)
}

func GetRuleList(c *gin.Context){
	models.GetRuleList(c)
}

func AddRole(c *gin.Context){
	models.AddRole(c)
}

func DelRole(c *gin.Context){
	models.DelRole(c)
}

func GetOneRole(c *gin.Context){
	models.GetOneRole(c)
}

func UpdateRole(c *gin.Context){
	models.UpdateRole(c)
}

func FetchUserList(c *gin.Context){
	models.FetchUserList(c)
}

func DelUser(c *gin.Context){
	models.DelUser(c)
}

func FetchArticleList(c *gin.Context){
	models.FetchArticleList(c)
}

func DelArticle(c *gin.Context){
	models.DelArticle(c)
}

func CreateArticle(c *gin.Context){
	models.CreateArticle(c)
}

func FetchOneArticle(c *gin.Context){
	models.FetchOneArticle(c)
}

func UpdateArticle(c *gin.Context){
	models.UpdateArticle(c)
}

func FetchMerchantList(c *gin.Context){
	models.FetchMerchantList(c)
}

func DelMerchant(c *gin.Context){
	models.DelMerchant(c)
}

func CreateMerchant(c *gin.Context){
	models.CreateMerchant(c)
}

func FetchOneMerchant(c *gin.Context){
	models.FetchOneMerchant(c)
}

func UpdateMerchant(c *gin.Context){
	models.UpdateMerchant(c)
}

func FetchOrderList(c *gin.Context){
	models.FetchOrderList(c)
}

func DelOrder(c *gin.Context){
	models.DelOrder(c)
}

func FetchServicesList(c *gin.Context){
	models.FetchServicesList(c)
}

func ChangeStatus(c *gin.Context){
	models.ChangeStatus(c)
}

func CreateService(c *gin.Context){
	models.CreateService(c)
}

func GetMerchant(c *gin.Context){
	models.GetMerchant(c)
}

func UpdateService(c *gin.Context){
	models.UpdateService(c)
}

func GetNewMsg(c *gin.Context){
	models.GetNewMsg(c)
}

func GetQuestionList(c *gin.Context){
	models.GetQuestionList(c)
}

func SetRead(c *gin.Context){
	models.SetRead(c)
}

func ArticleCreate(c *gin.Context){
	models.ArticleCreate(c)
}