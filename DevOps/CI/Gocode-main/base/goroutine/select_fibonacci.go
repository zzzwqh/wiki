package main

import "fmt"

func fibonacci(num chan int, quit chan interface{}) {
	x, y := 1, 1
	for {
		select {
		// 如果 channel num 可写，y 就会进入 channel
		case num <- y:
			medium := x
			x = x + y
			y = medium
		// 如果 channel num 不可写，那么就会取出 channel quit 信道中的值，Sub Goroutine 中的 quit <- 1 也不会阻塞
		case <-quit:
			fmt.Println("结束...")
			return
		}
	}
}

func main() {
	var c chan int = make(chan int)

	var quit chan interface{} = make(chan interface{})
	go func() {
		for i := 0; i < 10; i++ {
			// 取出 channel c 中的放入的值，容量是 0，如果不放入值会阻塞
			// 主 Goroutine 调用了 fibonacci() 函数，会放入值
			fmt.Printf("%v ", <-c)
		}
		quit <- 1
	}()

	fibonacci(c, quit)
}
