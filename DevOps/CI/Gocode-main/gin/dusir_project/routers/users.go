package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListUserHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg":  "ListUserHandler Router",
		"data": nil,
	})
}

func AddUserHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "添加 User 功能未开放...",
	})
}
func LoadUserHandler(r *gin.Engine) {
	r.GET("/user/list", ListUserHandler)
	r.POST("/user/add", AddUserHandler)

}
