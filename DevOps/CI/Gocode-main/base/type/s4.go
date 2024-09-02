package main

import (
	"fmt"
	"strconv"
)

// function 函数知识
func main() {
	PrintTest()
	PrintAB(1, 2)

	resValuez := ReturnOneArg(1, 4, "ethan")
	//var resValuez = ReturnOneArg(1,4,"ethan")
	//var resValuez int = ReturnOneArg(1,4,"ethan")
	fmt.Println(resValuez)

	resValuex, resValuey := ReturnTwoArg(1, 6, "sunflower")
	fmt.Println(resValuex, resValuey)

	resValue1, _, resValue2 := ReturnNotallArg(10, 24, "hello")
	fmt.Println(resValue1, resValue2)

	_, _, _, resValue4 := NameReturnAge(14, 51, "bye")
	fmt.Println(resValue4)
}

// 2. 简单的函数定义，无参数传入
func PrintTest() {
	fmt.Println("test")
}

// 3. 传入参数，只能按位置传，且无法设定参数默认值
func PrintAB(a, b int) {
	// (a int,b int)  (a string,b,c int)
	// (a string,b int)
	fmt.Println(a, b)
}

// 4. 传入参数 有一个返回值
func ReturnOneArg(a, b int, c string) int {
	d := a + b
	return d
}

// 5. 传入参数 有两个或多个返回值，需要括号
// 有几个返回值，调用的时候就需要用几个参数接收
func ReturnTwoArg(a, b int, c string) (int, int) {
	//x := a + b
	//y := a * b
	//return x,y
	return a + b, a * b
}

// 6. 有 N 个返回值时，只需要其中一个或者 < N 个返回值，接收时用 _ 表示忽略这个值
func ReturnNotallArg(a, b int, c string) (int, int, string) {
	x := a + b
	y := a - b
	z := c + c
	return x, y, z
}

// 7. 命名返回值，命名返回值后，无需再定义声明返回值变量，Return 中也无需指定返回值变量
func NameReturnAge(a, b int, c string) (add int, mul int, addStr string, con string) {
	add = a + b
	mul = a * b
	addStr = strconv.Itoa(add)
	con = addStr + c
	return
}
