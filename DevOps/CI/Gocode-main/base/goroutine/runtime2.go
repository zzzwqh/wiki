package main

import (
	"fmt"
	"runtime"
	"time"
)

func runtimeTask() {
	fmt.Println("runtimeTask 开始了！")
	defer fmt.Println("Defer 注册的代码块！")
	// 会将 goroutine 退出
	runtime.Goexit()
	// 下面这行不会打印
	fmt.Println("runtimeTask 结束了！")
}

func main() {
	go runtimeTask()
	time.Sleep(time.Second)
	fmt.Println("hello")
}
