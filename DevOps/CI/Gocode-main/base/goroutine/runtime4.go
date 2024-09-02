/**
GOMAXPROCE
调用runtime.GOMAXPROCS()用来设置可以并行计算cpu核数的最大值
并返回之前的值。默认值所有机器核数
*/
package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 获取机器的默认所有核心数
	fmt.Println("real GOMAXPROCS", runtime.GOMAXPROCS(-1))
	// 设置cpu最大核数
	n := runtime.GOMAXPROCS(1)
	fmt.Println("n 的返回值，一直是机器的所有核数", n)
	for {
		//两个协程抢着输出 0，1；观察01交替密度来观察；核数越大，交替越密
		go fmt.Print(0)
		fmt.Print(1)
	}
}
