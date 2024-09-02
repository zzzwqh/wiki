package main

import (
	"EthanCode/base/goroutine/entity"
	"fmt"
)

// 信道的长度和容量
func main() {
	// 信道的长度是，目前信道中有几个值
	// 信道的容量是，信道中总共可以放多少，不可以扩容（Slice可以扩容）
	var chanTest09 chan int = make(chan int)
	fmt.Println(len(chanTest09))
	fmt.Println(cap(chanTest09))
	fmt.Println("===============================")
	var chanTest10 chan int = make(chan int, 5)
	fmt.Println(len(chanTest10))
	fmt.Println(cap(chanTest10))
	fmt.Println("===============================")
	chanTest10 <- 1
	chanTest10 <- 2
	chanTest10 <- 3
	fmt.Println(len(chanTest10))
	fmt.Println(cap(chanTest10))

	// 信道 传递接口体类型
	var chanTest11 chan entity.User = make(chan entity.User, 1)
	// Send 一个结构体类型（字段过多，传递一个指针）
	chanTest11 <- *entity.NewUser("ethan", 25, "wqh3456@126.com")
	// Receive 一个结构体类型
	fmt.Println(<-chanTest11)

}
