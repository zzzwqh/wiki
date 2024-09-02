package main

import "fmt"

func main() {
	// 1、完整定义方式定义数组，数组中存储一系列相同类型的变量
	// 定义的时候，数组的长度大小就确定了，后期不能修改
	var arrayTest01 [3]int = [3]int{1, 2, 3}
	fmt.Println(arrayTest01)

	// 根据下标修改值
	arrayTest01[0] = 4
	fmt.Println(arrayTest01)
	fmt.Println(arrayTest01[0])

	// 2、类型推导方式定义数组，数组中存储的变量个数，可以少于数组的长度，不能超过数组的长度
	// 超过数组的长度会报错，数组越界
	var arrayTest02 = [3]int{5}
	// 如果存储的变量个数少于数组的长度，那么未赋值的变量置为 0 （整型数组）
	fmt.Println(arrayTest02)
	fmt.Println(arrayTest02[0])
	// 字符串数组则会置为空字符串
	var arrayTest03 = [2]string{"a"}
	fmt.Println(arrayTest03)
	fmt.Println(arrayTest03[0])

	// 3、简略声明方式定义数组
	// 指定数组索引位置，赋值变量
	arrayTest04 := [30]int{28: 111, 1: 3}
	fmt.Println(arrayTest04[28])
	fmt.Println(arrayTest04)

	// 4、不指定数组长度的方式，定义数组，根据变量的个数确定数组的长度，并不是定义可变长数组！！！
	var arrayTest05 = [...]int{21, 24, 32}
	fmt.Println(arrayTest05)
	fmt.Printf("这个变量类型是 %T \n", arrayTest05)
	var arrayTest06 = [...]int{49: 100}
	fmt.Println(arrayTest06[49])
	fmt.Printf("这个变量类型是 %T \n", arrayTest06)

	// 5、数组的长度获取
	fmt.Println("arrayTest06 数组的长度是", len(arrayTest06))

	// 6、数组的变量循环顺序输出
	for i := 0; i < len(arrayTest05); i++ {
		fmt.Print(arrayTest05[i], " ")
	}
	fmt.Println()
	// 7、数组的变量循环逆序输出
	for i := len(arrayTest05) - 1; i >= 0; i-- {
		fmt.Print(arrayTest05[i], " ")
	}
	fmt.Println()
	// 7、基于迭代的循环，需要使用 range 关键字，range 关键字可以返回一个值（即索引），也可以返回两个值（即索引和值）
	// 如下使用 range 关键字，i 是索引，下面这种方式还是根据索引遍历数组
	for i := range arrayTest05 {
		fmt.Println(i, arrayTest05[i])
	}
	// 下面才是使用迭代，正常是 for i,value := range arrayName 但不需要用到索引 i ，将其置 _
	for _, value := range arrayTest05 {
		fmt.Println(value)
	}

	// 	8、多维数组
	var arrayTest07 = [3][4]int{{1, 2, 3}, {2, 4}}
	fmt.Println(arrayTest07)

	// 9、多维数组的循环，需要使用 2 层循环
	for _, valueOutter := range arrayTest07 {
		for _, valueInner := range valueOutter {
			fmt.Println(valueInner)
		}
	}
}
