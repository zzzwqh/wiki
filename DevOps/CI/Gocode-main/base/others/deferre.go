package main

import (
	"fmt"
)

func main() {
	fmt.Println("return:", Demo2()) // 打印结果为 return: 2
}

func Demo2() int {
	var i int
	defer func() {
		i++
		fmt.Println("defer2:", i) // 打印结果为 defer: 2
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i) // 打印结果为 defer: 1
	}()
	return i // 或者直接 return 效果相同
}

// output：
//在 Demo2() 函数中，先是定义了两个 defer 语句，它们的执行顺序与定义顺序相反。接着，函数返回变量 i 的值，也就是 0。
//
//当函数返回时，先执行最后一个 defer 语句，也就是输出 defer2: 1。随后执行第一个 defer 语句，
//输出 defer1: 2。需要注意的是，在 defer 中修改 i 的值对函数的返回值没有影响，因为在 defer 执行前已经确定了返回值。
//
//因此，虽然 Demo2() 函数中有 return i 语句，但实际上函数返回的值永远是 0。
