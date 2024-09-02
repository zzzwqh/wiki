package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 多个 Goroutine 并发访问/修改同一个 Map 类型数据时，会 fatal
// 定义一个并发安全的 sync.Map
var safeMap sync.Map

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 1; i < 11; i++ {
		go func(num int) {
			key := fmt.Sprintf("student%02d", num)
			// 用 Store 方法定义 Key:Value
			safeMap.Store(key, rand.Intn(100))
			// 用 Load 方法根据 Key 获取 Value，一共有两个返回值，ok 代表是否有这个 Key
			score, ok := safeMap.Load(key)
			fmt.Println(score, ok)
		}(i)
	}
	time.Sleep(time.Second * 1)
}
