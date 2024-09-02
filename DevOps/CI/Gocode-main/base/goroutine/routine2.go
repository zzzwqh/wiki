package main

import "fmt"

func ProducerTask(c chan interface{}) {
	for i := 0; i < 10; i++ {
		c <- i
		fmt.Println("副 Goroutine 已发送", i)
	}
	// 使用 close() 函数可以关闭信道
	close(c)
}

// 循环的从信道输入、取出值
func main() {
	//
	var chanTest07 chan interface{} = make(chan interface{})
	go ProducerTask(chanTest07)

	// 主线程等待取出
	// for {
	// 	value, ok := <-chanTest07
	// 	// 判断，如果 ok == false ，意味着信道已经被关闭，那么就结束循环
	// 	if ok == false {
	// 		break
	// 	}
	// 	fmt.Println("主 Goroutine 接收到", value)
	// }

	// 注释掉上方代码，更优雅的写法，这里 range 会判断信道是否关闭
	for value := range chanTest07 {
		fmt.Println("主 Goroutine 接收到", value)
	}
}
