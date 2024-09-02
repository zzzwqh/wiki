package main

import (
	"fmt"
	"sync"
	"time"
)

// 问题一：为什么需要 context?
// 答案: 为了通知 goroutine ,  优雅的让 goroutine 退出
// 问题二：如果没有 Context，我们如何 [通知] goroutine 运行退出?
// 思考：定义一个全局变量是否可行？

var wg4A sync.WaitGroup // wg 的作用只是让 Main 主协程等待每个 Goroutine 的运行，不是让 Main 协程通知到 Goroutine 退出
var notify bool = false

func processA() {
	defer wg4A.Done()
	for {
		fmt.Println("ProcessA Goroutine Task Running...")
		time.Sleep(time.Second)
		if notify == true {
			break
		}
	}
}

func main() {
	wg4A.Add(1)
	go processA()
	// 主协程等待 3 秒后，将 notify 全局变量设置为 true，即可达到通知 goroutine 退出效果
	for i := 0; i < 3; i++ {
		fmt.Println("Main Routine Running At Step", i)
		time.Sleep(time.Second)
	}
	notify = true
	wg4A.Wait()
}
