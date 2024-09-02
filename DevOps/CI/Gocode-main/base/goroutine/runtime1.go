package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	fmt.Println(runtime.NumCPU())  // CPU 逻辑核心数（超线程数）
	fmt.Println(runtime.GOARCH)    // amd 架构
	fmt.Println(runtime.Version()) // go 版本
	fmt.Println(runtime.GOOS)      // Windows 系统
	fmt.Println(runtime.GOROOT())  // Golang 的根目录

	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(wg *sync.WaitGroup) {
			time.Sleep(time.Second * 2)
			wg.Done()
		}(&wg)
	}
	fmt.Println(runtime.NumGoroutine()) // 101 个 Goroutine
	wg.Wait()
	fmt.Println("结束")
}
