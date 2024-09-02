package main

import (
	"fmt"
)

func ProducerProgram(c chan interface{}) {
	for i := 1; i < 4; i++ {
		c <- i * i * i
	}
	//
	close(c)
	// 如下这行代码，是在信道关闭后，再次 Send 值到信道中
	// c <- "信道关了能不能再传送进去值？"
	// 不可行，报错内容为 panic: send on closed channel
}

// 有缓冲信道，需要在初始化时，指定 capacity，不指定默认是 0
func main() {
	var chanTest08 chan interface{} = make(chan interface{}, 4)
	// 向信道中输入三个值，不使用 goroutine
	ProducerProgram(chanTest08)
	// 如果是无缓冲信道，下面这句不会运行
	fmt.Println("Producer program Completed...")
	//fmt.Println(<-chanTest08)
	//fmt.Println(<-chanTest08)
	//fmt.Println(<-chanTest08)

	// 如果不关闭信道，使用 range 取值会报错
	// 同时，即使关闭了信道，值一直在，是可以取出来的
	for value := range chanTest08 {
		fmt.Println(value)
	}

}
