package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 多级 Goroutine 能不能都被 Main routine 通知退出？
func processE(wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
LOOP:
	for {
		fmt.Println("ProcessE Goroutine Task Running...")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			fmt.Println("ProcessE has received Exist Signal gave by ProcessD Goroutine ~")
			break LOOP
		default:

		}
	}
}

func processD(wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	wg.Add(1)
	go processE(wg, ctx)
LOOP:
	for {
		fmt.Println("ProcessD Goroutine Task Running...")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			fmt.Println("ProcessD has received Exist Signal gave by Main routine ~")
			break LOOP
		default:

		}
	}
}

func main() {
	var wg4D sync.WaitGroup                                 // 为了让主协程等待子协程运行完
	ctx, cancel := context.WithCancel(context.Background()) // 为了让主协程通知（控制）子协程退出
	wg4D.Add(1)
	go processD(&wg4D, ctx)
	for i := 0; i < 3; i++ {
		fmt.Println("Main Routine Running Step", i)
		time.Sleep(time.Second)
	}
	cancel()
	wg4D.Wait()
}

/* 输出结果如下，挺规矩的，如果是自己控制 channel 会比较麻烦
比如此例子中，主协程到子协程用的是一个 context 封装的 Channel，子协程和孙子协程用的其实是另外一个 context 封装的 Channel
Main Routine Running Step 0
ProcessD Goroutine Task Running...
ProcessE Goroutine Task Running...
ProcessE Goroutine Task Running...
ProcessD Goroutine Task Running...
Main Routine Running Step 1
ProcessD Goroutine Task Running...
ProcessE Goroutine Task Running...
Main Routine Running Step 2
ProcessE Goroutine Task Running...
ProcessD Goroutine Task Running...
processD has received Exist Signal gave by Main routine ~	// 收到信号的两行都打印出来了，这两行打印没有一定的先后顺序
processE has received Exist Signal gave by ProcessD Goroutine ~
*/
