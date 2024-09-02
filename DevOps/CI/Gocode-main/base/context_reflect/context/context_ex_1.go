package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 如何使用 Context 优雅的通知 Goroutine 退出
// 如果 wg 不定义为全局变量，而作为参数传入函数时，一定要记得将指针传入，否则会传入 wg 的副本！导致 Deadlock！
func processC(wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
LOOP:
	for {
		fmt.Println("ProcessC Goroutine Task Running...")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			fmt.Println("Exist Signal gave by Main routine ~")
			break LOOP
		default:

		}
	}

}

func main() {
	var wg4C sync.WaitGroup                                 // 为了让主协程等待子协程运行完
	ctx, cancel := context.WithCancel(context.Background()) // 为了让主协程通知（控制）子协程退出
	wg4C.Add(1)
	go processC(&wg4C, ctx)
	for i := 0; i < 3; i++ {
		fmt.Println("Main Routine Running Step", i)
		time.Sleep(time.Second)
	}
	cancel()
	wg4C.Wait()

}
