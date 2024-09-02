package main

import (
	"fmt"
	"sync"
)

var globalX int = 1

func taskForOneVar(lock *sync.Mutex, wg *sync.WaitGroup) {
	// 为什么让 wg.Done() 注册并延迟调用？为了防止代码运行异常，导致该 SubGoroutine 没有返回 wg.Done()，继而导致 MainGoroutine 一直在 wg.Wait() 处等待
	defer wg.Done()
	// 上 Mutex 互斥锁
	lock.Lock()
	globalX++ // 这段代码叫做临界区
	// 解 Mutex 互斥锁
	lock.Unlock()
}
func main() {
	// Mutex 是一个结构体（值类型），有加锁（Locker）、解锁（Unlocker）方法，实现了 Locker 接口
	var lockCode sync.Mutex
	// 如果不用 wg.Wait() 等所有 taskForOneVar() 函数执行完，中途打印出的数字，是小于 1000 的，等待所有 Goroutine 执行完毕才结果是 1000
	var wg sync.WaitGroup
	for i := 1; i < 1000; i++ {
		wg.Add(1)
		// sync.WaitGroup 和 sync.Mutex 都是 struct 类型，也就意味着需要传入地址的指针进去，否则是 Copy 值传递的方式传入参数
		go taskForOneVar(&lockCode, &wg)
	}
	wg.Wait()
	fmt.Println(globalX)
}
