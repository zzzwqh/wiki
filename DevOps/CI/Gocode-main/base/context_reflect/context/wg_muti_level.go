package main

import (
	"fmt"
	"sync"
)

// 本例子尝试验证，多层函数嵌套中，用一个 *sync.WaitGroup 变量，让主协程等待子协程的运行，是否可行（可行）
func process4WG2(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("WG2 ....") // 这行会打印出来
}

func process4WG1(wg *sync.WaitGroup) {
	defer wg.Done()
	wg.Add(1)          // 在子协程中写入 Add
	go process4WG2(wg) // 传入子协程的子协程
	fmt.Println("WG1 ....")
}
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go process4WG1(&wg)
	wg.Wait()
}
