package main

import "fmt"

var variableTest01 string
var variableTest02 int = 1

func main() {
	// 变量的定义和声明
	// 1.1 完整定义
	var nameOfEthan string = "ethan"
	var ageOfEthan int = 24
	fmt.Println(nameOfEthan, ageOfEthan)

	// 1.2 类型推导
	var nameOfNoah = "noah"
	var ageOfNoah = 18
	fmt.Println(nameOfNoah, ageOfNoah)

	// 查询变量的类型
	fmt.Printf("type is %T，value is %d .\n", ageOfNoah, ageOfNoah)
	fmt.Printf("type is %T，value is %s .\n", nameOfNoah, nameOfNoah)

	// 1.3 简略声明，必须加 ：号，加冒号的意思是定义变量，不加冒号的意思是修改变量值
	nameOfKevin := "kevin"
	ageOfKevin := 21
	fmt.Println(nameOfKevin)
	fmt.Println(ageOfKevin)
	ageOfKevin = 30 // 不加冒号意味着修改变量值，无法重复定义同一个变量
	fmt.Println(ageOfKevin)

	// 2. 变量的使用
	// 修改变量值，但不能改变其变量类型，变量类型在定义阶段就固定了
	var nameOfSun string = "sun"
	nameOfSun = "moon"
	fmt.Println(nameOfSun)

	// 3. 同时定义多个变量（在一行同时定义多个变量）
	// 3.1 可以这样写
	var nameOfMoon, ageOfMoon, sexOfMoon = "moon", 21, "male"
	// 如果多个变量的类型不一致，只能用上述类型推导方式定义，无法完整定义其所有类型
	// 如果完整定义多个变量，那么这些变量只能是同一个类型，不然会报错
	var nameOfStar, ageOfStar string = "star", "19"
	fmt.Println(nameOfMoon, ageOfMoon, sexOfMoon)
	fmt.Println(nameOfStar, ageOfStar)
	// 3.2 也可以这样写
	var (
		nameOfTest1 string
		ageOfTest1  int
		nameOfTest2 string = "test2"
		ageOfTest2  int    = 15
	)
	fmt.Println(nameOfTest1, ageOfTest1, nameOfTest2, ageOfTest2)

	// 4. 变量要先定义，再使用，并且只能定义一次
	variableTest03 := "ethanz"
	fmt.Println(variableTest03)

	// 5. 小细节，重复定义变量
	var myAge = 19
	myName, myAge := "ethan", 20
	fmt.Println(myName, myAge)

	var a, b, c, d = 1, "a", [3]int{}, []string{}
	fmt.Println(a, b, c, d)
}
