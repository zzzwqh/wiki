package main

import (
	"fmt"
	"sync"
)

func task(i interface{}, wg *sync.WaitGroup) {
	fmt.Println("Other Goroutine for", i)
	wg.Done()

}

func main() {
	// 使用 waitgroup 可以控制主 Goroutine 等待子 Goroutine 的完成
	// waitgroup 是可以定义的，并且累死结构体类型，零值不是 nil
	var wg sync.WaitGroup
	fmt.Println(&wg)
	// 启动 3 个子协程
	// wg.Add() 与 wg.Wait() 一起使用，否则 wg.Wait() 不会等待，Wait() 等待的是 wg.Done()
	// wg.Add(Num) 意味着 wg.Wait() 要等待 Num 个子协程给出 wg.Done 信号

	go task("ethan", &wg)
	go task("noah", &wg)
	go task("maven", &wg)
	wg.Add(3)
	fmt.Println("deng.......................")
	wg.Wait()

	// 优雅
	var wgTest sync.WaitGroup
	for i := 0; i < 10; i++ {
		// wgTest.Add(1) 要在 task() 上面，否则会报错，Add 要在 Done 之前
		wgTest.Add(1)
		task(i, &wgTest)
	}
	wgTest.Wait()
}
