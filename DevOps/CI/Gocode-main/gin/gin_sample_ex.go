package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type UserInfo struct {
	Name   string `json:"name" binding:"required"`
	Passwd string `json:"passwd" binding:"required"`
}

func main() {
	r := gin.Default()
	r.GET("/", DealGetFormHandler)
	r.POST("/login", DealPostJsonHandler)
	// 使用 goroutine 运行，是为了优雅退出
	go func() {
		r.Run(":9090")
	}()
	// 使用 quit 接受 syscall 信号 ... 接收到后在主协程中优雅退出
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	log.Println("接收到 SIGTERM 信号，优雅退出......")
	log.Println("处理后事......")

}
func DealPostJsonHandler(c *gin.Context) {
	user := &UserInfo{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": user,
		"code": 0,
	})
}
func DealGetFormHandler(context *gin.Context) {
	user := &UserInfo{}
	if err := context.ShouldBind(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Msg": err.Error()})
	}
	context.JSON(http.StatusOK, &user)
}
