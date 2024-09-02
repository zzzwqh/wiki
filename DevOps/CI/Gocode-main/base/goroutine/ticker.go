package main

import (
	"fmt"
	"time"
)

func main() {
	// timer 对象，当时间到了，我们就无法再次使用了，使用 timer.Reset() 才能将 timer 对象重置
	// 定时执行器，ticker 对象，可以让我们实现，即使时间到了，也可以多次执行
	ticker1 := time.NewTicker(time.Second)
	// 执行几次？
	times := 0
	for {
		fmt.Println(<-ticker1.C)
		times++
		if times == 10 {
			// ticker.Stop() 使用后，下次 ticker.C 就无法再次使用了，再次使用会报错 deadlock
			ticker1.Stop()
			fmt.Println("ticker stoped...")
			// 如果不加 break，那么会继续循环，fmt.Println(<-ticker1.C) 这行代码会报错 deadlock
			// break
		}
	}
}
