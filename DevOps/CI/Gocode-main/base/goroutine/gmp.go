package main

import (
	"fmt"
	"time"
)

func main() {
	// 如果设置数量为 1 ，可以保证输出的数字是升序的，如果不设置，默认是 CPU 核心数，不能保证输出的数字是顺序的，因为使用了 Goroutine 并发
	// runtime.GOMAXPROCS(1)
	for i := 1; i < 10000; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(time.Second)
}
