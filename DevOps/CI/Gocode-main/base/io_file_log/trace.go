package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
)

func task(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Goroutine Task...")
}

// 生成的 trace.out 文件，用 go tool trace trace.out 命令去查看
func main() {
	//创建 trace 文件
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	//启动 trace goroutine
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	// main
	var wg sync.WaitGroup
	fmt.Println("Hello World")
	runtime.GOMAXPROCS(runtime.NumCPU())
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go task(&wg)

	}
	wg.Wait()

}
