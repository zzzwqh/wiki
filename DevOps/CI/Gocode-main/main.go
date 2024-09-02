package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	ginEngine := gin.Default()
	ginEngine.GET("/", func(context *gin.Context) {
		fmt.Println("This is Root Page")
		context.String(200, "hex")
	})
	ginEngine.GET("/index", handleFuncIndex)
	err := ginEngine.Run("0.0.0.0:1234")
	if err != nil {
		fmt.Println(err)
	}
}
func handle() {
	fmt.Println("abc")
	fmt.Println("asdf ")
	context.TODO()
	
}
func handleFuncIndex(context *gin.Context) {
	index, err := os.Open("D:\\Users\\ethan\\GolandProjects\\GoCode\\gin\\index.html")
	if err != nil {
		fmt.Println("Open index.html File Error :", err)
	}
	defer index.Close()
	var buf = make([]byte, 1)
	reader := bufio.NewReader(index)
	for {

		line, _, err := reader.ReadLine()
		fmt.Println(string(line))
		if err != nil || err == io.EOF {

			fmt.Println("reader.ReadLine() Error:", err)
			break
		}
		buf = append(buf, line...)
	}
	fmt.Println("=====================================")
	fmt.Println(string(buf))
	context.String(200, string(buf))
}
