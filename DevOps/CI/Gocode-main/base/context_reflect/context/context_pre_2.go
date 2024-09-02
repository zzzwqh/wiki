package main

import (
	"fmt"
	"sync"
	"time"
)

// 问题：不适用全局变量，还有什么方法通知 goroutine 优雅的退出?
// 答案：Channel 管道

var wg4B sync.WaitGroup
var chan4B = make(chan struct{})

func processB() {
	defer wg4B.Done()
LOOP:
	for {
		fmt.Println("ProcessB Goroutine Task Running...")
		time.Sleep(time.Second)
		select {
		case chan4B <- struct{}{}:
			fmt.Println("Main routine give notice ~~")
			break LOOP
		default:

		}

	}
}

func main() {
	wg4B.Add(1)
	go processB()
	for i := 0; i < 3; i++ {
		fmt.Println("Main routine Running Step", i)
		time.Sleep(time.Second)
	}
	<-chan4B
	wg4B.Wait()
}
