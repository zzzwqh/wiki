package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 常规的 map 类型，在 Goroutine 并发操作时，不安全
var stuList map[string]interface{} = make(map[string]interface{})

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 1; i < 3; i++ {
		go func(num int) {
			key := fmt.Sprintf("students%02d", num)
			stuList[key] = rand.Intn(100)
		}(i)
	}
	time.Sleep(time.Second * 3)
	fmt.Println(stuList)
}
