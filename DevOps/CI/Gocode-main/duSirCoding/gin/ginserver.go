package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/api/hello", func(context *gin.Context) {
		context.String(http.StatusOK, "hello")
	})
	router.Run(":10001")
}
