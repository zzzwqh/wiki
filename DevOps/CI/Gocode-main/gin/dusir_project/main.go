package main

import (
	"EthanCode/gin/dusir_project/middlewares"
	"EthanCode/gin/dusir_project/routers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 新创建一个 *engine 对象，默认会使用两个中间件（日志和错误处理），然后调用 gin.New() 方法
	r := gin.Default()
	// 新增一个自定义的中间件（全局）
	r.Use(middlewares.Requestid())

	// 1. 简单的添加路由，传入 *gin.engine 加载路由
	routers.LoadUserHandler(r)

	// 2. 创建路由分组，比如 http://$domain/book 的请求都会分配到这个路由分组中，那先将 /book 前缀拎出来建组
	bookGroup := r.Group("/book")
	// 加载这个路由分组下的所有路由，传入 *gin.RouterGroup 加载路由
	routers.LoadBookHandler(bookGroup)
	r.POST("/upload1", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		// c.SaveUploadedFile(file, dst)
		c.String(http.StatusOK, "%s uploaded!", file.Filename)
	})
	r.Run(":9090")
}
