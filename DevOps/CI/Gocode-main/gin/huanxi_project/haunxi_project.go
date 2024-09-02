package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/slice", SliceHandler)
	router.GET("/var", VariableHandler)
	router.Run()
}

func VariableHandler(c *gin.Context) {
	res := c.Query("id")
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": res,
	})
}

func SliceHandler(c *gin.Context) {
	res := c.QueryArray("my_array")
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": res,
	})
}
