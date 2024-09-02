package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker4Timeout(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Timeout Signal Received by Main routine ....")
			return // return 语句执行之前会执行 defer 哦，但是如果有返回值，会先赋值返回值（可复习 defer 的正确使用姿势）
		default:
			time.Sleep(time.Second)
			fmt.Println("Connecting Mysql Databases ...")
		}
	}
}
func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	// 和 deadline 例子一样，在任何情况下都调用 cancel，避免上下文以及其父级生命周期过长，超过必要时间
	// 实际上，即使不调用 cancel，也会因为 Timeout 超时时间而结束 worker 携程
	defer cancel()
	wg.Add(1)
	go worker4Timeout(ctx, &wg)
	wg.Wait()
}
