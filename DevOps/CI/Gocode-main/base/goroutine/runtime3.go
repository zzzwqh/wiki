package main

import (
	"fmt"
	"runtime"
)

// Goshced 让出 CPU 时间片让其他 Goroutine 先行运行
func main() {

	go func() {
		for i := 0; i < 100; i++ {
			runtime.Gosched()
			fmt.Println("goroutine ①", i)
		}
	}()
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("goroutine ②", i)
		}
	}()
	for i := 0; i < 100; i++ {
		runtime.Gosched()
		fmt.Println("main", i)
	}

}
