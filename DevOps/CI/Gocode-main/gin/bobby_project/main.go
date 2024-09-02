package main

import (
	"EthanCode/gin/bobby_project/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	accountGroup := r.Group("/account")
	{
		accountGroup.POST("/login", AccountLoginHandler)
	}
	r.Run(":9090")
}

func AccountLoginHandler(c *gin.Context) {
	account := model.Account{}
	// 使用 postman 传递 raw 格式的 json 数据时，需要搭配 c.ShouldBindJSON(&account) => gorm:"json:xxx" ，即结构体模型中要写 json 的 tag
	// 使用 postman 传递 form 表单时，可以使用如下示例搭配 c.ShouldBind(&account) => gorm:"form:xxx" ，即结构体模型中要写 form 的 tag
	if err := c.ShouldBind(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":  "Welcome to This Page",
		"account":  account.Name,
		"password": account.Passwd,
	})

}
