package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Book struct {
	Num  int32
	Name string
}

func ListBookHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg":  "ListBookHandler Router",
		"data": nil,
	})
}
func AddBookHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "添加 Book 功能研发中...",
	})
}

// LoadBookHandler 通过路由组的方式，返回路由组两个接口
func LoadBookHandler(rg *gin.RouterGroup) {
	rg.GET("/list", ListBookHandler)
	rg.POST("/add", AddBookHandler)
}
