package main

import (
	"fmt"
)

func main() {

	ReceiveArgsInt(1, 3, 4, 5)
	ReceiveArgsStr("ethan", "zzz")
	// 2. 匿名函数 =====>  只能定义在函数内部，不能有名字（不配）
	// 2.1 基本匿名函数，直接调用
	func() {
		fmt.Println("1")
	}()
	// 2.1 基本匿名函数，将其赋值给变量，然后调用，因为 Golang 里面，函数是一等公民，即函数也是一种类型！
	varFun := func() { fmt.Println("2") }
	varFun()
	// 2.2  有参数的匿名函数
	func(a, b int, c string) {
		fmt.Println(a, b, c)
	}(1, 2, "ethan")
	// 2.3 有返回值的匿名函数
	res := func(a, b int, c string) int {
		return a + b
	}(100, 2, "ethan")
	fmt.Println(res)
	// 3. 函数也是一种类型，可以赋值给变量，然后调用
	fType := func() { fmt.Println("3") }
	fType()            // 调用函数
	fmt.Println(fType) // 获取函数变量地址

	// 4.既然函数可以是变量，那么就可以做参数传递
	resx, resFunc := FuncArgTest(1, 2, "ethanz", fType)
	resFunc()
	fmt.Println(resFunc)
	fmt.Println(resx)

	fmt.Println("===============")
	resFuncz := ReturnFunc()
	// 上述简略声明的 resFuncz 完整定义应该如下
	// var resFuncz func(a int, b int) int = ReturnFunc()
	resFuncOfFuncz := resFuncz(1, 2)
	fmt.Println(resFuncOfFuncz)

	fmt.Println("---------------")
	resFuncx := ReturnFuncTest()
	// 上述简略声明的 resFuncx 完整定义可以如下（使用了 type 命名匿名函数的别名，然后 ReturnFuncTest 中返回值使用了这个别名，所以可以这么写）
	// var resFuncx MyType = ReturnFuncTest()
	resFuncOfFuncx := resFuncx(3, 4)
	fmt.Println(resFuncOfFuncx)
}

// 1.1 可变长参数，可以接收任意个 int 类型的参数
func ReceiveArgsInt(a ...int) (b []int) {
	fmt.Println(a)
	return
}

// 1.2 可变长参数，接收任意个 string 类型的参数
func ReceiveArgsStr(str ...string) {
	fmt.Println(str)

}

// 4. 既然函数可以是变量，那么就可以做参数传递
func FuncArgTest(a, b int, c string, d func()) (x int, y func()) {
	x = a * b
	// 闭包函数
	y = func() {
		fmt.Println(x)
		fmt.Println("yyyyyy")

	}
	return
}

func ReturnFunc() func(a, b int) int {
	z := func(a, b int) int {
		fmt.Println("我是内层函数 1")
		return a + b
	}
	return z
}

// 5. 上面返回值的类型太繁琐，我们可以定义一个别名
type MyType func(a, b int) int

func ReturnFuncTest() MyType {
	z := func(a, b int) int {
		fmt.Println("我是内层函数 2")
		return a + b
	}
	return z
}
