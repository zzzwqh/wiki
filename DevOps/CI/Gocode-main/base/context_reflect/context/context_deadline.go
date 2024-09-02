package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 定义一个截至时间点
	deadLineTime := time.Now().Add(time.Second * 2)
	ctx, cancel := context.WithDeadline(context.Background(), deadLineTime)
	// 尽管 ctx 将过期，但最佳实践是
	// 在任何情况下都调用 Cancel() 函数
	// 否则 ctx 及其父级的生存时间可能会超过必要的时间，这也是为什么 WithDeadLine 函数仍然返回一个 cancel 函数的原因
	defer cancel()
	select {
	case <-time.After(3 * time.Second): // time.After(time time.Duration) 会等待 3 秒后，在返回的信道上发送当前时间
		fmt.Println("overslept")
	case <-ctx.Done(): // 设置了 DeadLine 的 ctx 在 2 秒后，在返回的信道上发送一个 struct{} 类型空值，那么上面的 select 语句就不会被执行
		fmt.Println(ctx.Err())
	}
}
