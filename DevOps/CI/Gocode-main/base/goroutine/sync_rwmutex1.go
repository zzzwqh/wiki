package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// 全局共享变量，用来测试读写锁的效率
var globalY int = 1

// 定义一个互斥锁
var lock sync.Mutex

// 定义一个读写锁
var rwlock sync.RWMutex

func writeY(wg *sync.WaitGroup) {
	defer wg.Done()
	rwlock.Lock()
	globalY += 10
	rwlock.Unlock()
}
func readY(wg *sync.WaitGroup) {
	defer wg.Done()
	// 使用互斥锁（访问同一个共享变量，即使读也是要竞争的）
	lock.Lock()
	lock.Unlock()
	// 使用读锁，访问同一个共享变量，读锁是可以避免竞争的，那么理论上这段代码更快些
	// rwlock.RLock()
	// rwlock.RUnlock()
}

func main() {
	startTime := time.Now()
	var wg sync.WaitGroup
	// 设置 CPU 数，为逻辑 CPU 的数量 - 1
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go writeY(&wg)
	}

	for i := 0; i < 10000000; i++ {
		wg.Add(1)
		go readY(&wg)
	}
	wg.Wait()
	endTime := time.Now()
	// 计算出从 startTime - endTime 度过的时间，以此衡量读锁和互斥锁的效率
	fmt.Println("用时", endTime.Sub(startTime))
}
