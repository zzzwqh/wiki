package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Args 是一个[]string
	if len(os.Args) > 0 {
		// 轮询取出 os.Args 列表中的内容（即命令行参数）
		for index, arg := range os.Args {
			fmt.Printf("args[%d]=%v\n", index, arg)
		}
	}
	// 取出第一个参数
	fmt.Println(os.Args[1])
}
