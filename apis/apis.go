package apis

import (
	"awesomeProject/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexApi(c *gin.Context){
	c.String(http.StatusOK, "It works")
	user := models.GetAllUser()
	c.JSON(200,user)
}
