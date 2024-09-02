package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Requestid() gin.HandlerFunc {
	return func(context *gin.Context) {
		log.Println("这是 requestId 中间件")
		context.Next()
		log.Println("requestId 中间件结束了")

	}
}
