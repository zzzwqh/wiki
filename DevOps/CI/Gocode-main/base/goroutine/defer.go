package main

import (
	"fmt"
)

func ReceiveInd(s string, i *int) {
	fmt.Println(s, *i)
}
func main() {
	fmt.Println("Main Goroutine Running......")
	var i int = 0

	defer func() {
		fmt.Println("此时 i 变量的值", i)
	}()
	defer func() {
		// 这里还算是闭包函数，引用了外部的变量 i，并非 COPY 传递进来一个参数
		for ; i < 10; i++ {
			fmt.Println("====", i)
		}
	}()

	i++
	fmt.Println("Main Goroutine Completed......")
}
