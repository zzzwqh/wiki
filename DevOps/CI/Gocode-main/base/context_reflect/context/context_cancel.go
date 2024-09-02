package main

import (
	"context"
	"fmt"
	"time"
)

func generate(ctx context.Context) (ch <-chan int) {
	// 这里需要重新定义一个 dst 信道，不直接使用 ch 的原因是，ch 已经被定义成 only-receive 的信道，而该 gen 函数需要生成数据传入信道
	dst := make(chan int)
	var num int
	// 这里 gen 直接再次单独启用一个 Goroutine ，使用了刚刚定义的 dst 信道
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("cancel() 被调用了，ctx.Done() 的信道中传出值了...")
				return
			case dst <- num:
				num++

			}
		}
	}()

	// 让上面的 Goroutine 单独运行去，gen 函数先将 dst 信道返回给主协程
	return dst
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	// TIPS: 这里写 defer cancel() 是看不到 "cancel() 被调用了，ctx.Done() 的信道中传出值了..." 被打印的，因为主协程率先结束了
	defer cancel()

	// 循环从信道中取出元素（只要信道不关闭就不结束 for 循环）
	for elem := range generate(ctx) {
		fmt.Println(elem)
		if elem == 5 {
			break
		}
	}
	// TIPS: 想看到  "cancel() 被调用了，ctx.Done() 的信道中传出值了..." 被打印，就先调用 cancel()，再让主协程 sleep 一秒
	cancel()
	time.Sleep(time.Second)

}
