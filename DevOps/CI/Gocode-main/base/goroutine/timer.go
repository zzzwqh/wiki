package main

import (
	"fmt"
	"time"
)

func main() {
	// t0 <- timer.C 获取得到的 t0 是延迟后的时间
	fmt.Println("Main 程序执行开始了 ~")
	// 获取并打印当前时间 t
	t := time.Now()
	fmt.Println(t)
	// 传入一个时间，延迟多长时间，t0 是延迟后的时间（t 时刻 + 3s = t0 时刻）
	timer := time.NewTimer(time.Second * 3)
	t0 := <-timer.C
	fmt.Println("Main 程序执行完了 ~")
	fmt.Println(t0)

	// 1.timer基本使用
	timer1 := time.NewTimer(2 * time.Second)
	t1 := time.Now()
	fmt.Printf("t1:%v\n", t1)
	t2 := <-timer1.C
	fmt.Printf("t2:%v\n", t2)

	// 2.timer只能使用一次
	//timer2 := time.NewTimer(time.Second)
	//for {
	//	//timer2.Reset(time.Second)
	//	<-timer2.C
	//	fmt.Println("时间到")
	//}

	// 3.重置定时器
	timer5 := time.NewTimer(3 * time.Second)
	timer5.Reset(1 * time.Second)
	fmt.Println(time.Now())
	fmt.Println(<-timer5.C)

}
