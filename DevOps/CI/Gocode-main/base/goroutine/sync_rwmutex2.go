package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var count int

var mutex sync.RWMutex

func write(n int) {
	rand.Seed(time.Now().UnixNano())
	fmt.Printf("写 goroutine %d 正在写数据...\n", n)
	mutex.Lock()
	num := rand.Intn(500)
	count = num
	fmt.Printf("写 goroutine %d 写数据结束，写入新值 %d\n", n, num)
	mutex.Unlock()

}
func read(n int) {
	mutex.Lock()
	fmt.Printf("读 goroutine %d 正在读取数据...\n", n)
	num := count
	fmt.Printf("读 goroutine %d 读取数据结束，读到 %d\n", n, num)
	mutex.Unlock()
}
func main() {
	for i := 0; i < 10; i++ {
		go read(i + 1)
	}
	for i := 100; i < 110; i++ {
		go write(i + 1)
	}
	time.Sleep(time.Second * 5)
}
