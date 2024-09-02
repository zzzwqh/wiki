package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 浮点型转整型，整型转浮点型
	var a float64 = 10.11
	var b int = 5
	fmt.Println(int(a) + b)
	fmt.Println(a + float64(b))

	// int(float) 无法和 string 类型互转，如果想转换，要用 strconv 包
	// 1. string => int	====== Atoi
	strTest01 := "1001"
	fmt.Printf("Type is %T,Value is %v\n", strTest01, strTest01)
	revStrTest01, err := strconv.Atoi(strTest01)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Type is %T,Value is %v\n", revStrTest01, revStrTest01)
	fmt.Println("=====================")

	// 2. int => string	====== Itoa
	intTest01 := 10
	revIntTest01 := strconv.Itoa(intTest01)
	fmt.Printf("Type is %T,Value is %v\n", revIntTest01, revIntTest01)
	fmt.Println("=====================")

	// 3. string => float ====== ParseFloat
	strTest02 := "10.242415223"
	if revStrTest02, err := strconv.ParseFloat(strTest02, 64); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(revStrTest02)
		fmt.Printf("Type is %T,Value is %v\n", revStrTest02, revStrTest02)
	}
	fmt.Println("=====================")

	// 4.1 ParseBool 类型，把字符串转成对应的类型，字符串可以是 true/false/1/0
	boolTest01 := "0"
	if revBoolTest01, err := strconv.ParseBool(boolTest01); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(revBoolTest01)
		fmt.Printf("Type is %T,Value is %v\n", revBoolTest01, revBoolTest01)
	}

	// 4.2 ParseInt 类型，把字符串转成相对应的类型，base 是进制、bitSize 精度，转换成功后的数据类型都是 int64 ！
	i, _ := strconv.ParseInt("-110", 2, 64)
	fmt.Printf("i的类型是%T，值是：%v\n", i, i)

	// 4.3 ParseUnit 类型，不接受正负号
	z, _ := strconv.ParseUint("16", 10, 0)
	fmt.Printf("i的类型是%T，值是：%v\n", z, z)
	fmt.Println("=====================")

	// 5.1 FormatBool 类型，将 Bool 转成 String 类型
	fmt.Printf("类型是 %T，值是 %v\n", strconv.FormatBool(true), strconv.FormatBool(true))

	// 5.2 FormatInt 类型
	fmt.Printf("类型是 %T，值是 %v\n", strconv.FormatInt(16, 16), strconv.FormatInt(16, 16))

	// 6. IsPrint 方法，判断一个 rune 字符是否可以打印
	resBool01 := strconv.IsPrint('\n')
	resBool02 := strconv.IsPrint('中')
	fmt.Println(resBool01)
	fmt.Println(resBool02)
	fmt.Println("=====================")

	// 7. CanBackquote 判断是否有换行，如果这个字符串没有换行，就返回 true
	fmt.Println(strconv.CanBackquote(`I'm bluedusk `))
	fmt.Println(strconv.CanBackquote("I'm bluedusk \n and you?"))

	// 8. Append 类型方法，没啥用，只能 append byte[] 切片
	byteSliceTest01 := []byte{1, 'z'}
	// 需要有接收的 byte[] 切片，不会改变原有 byte[] 切片
	byteSliceAfterAppend := strconv.AppendInt(byteSliceTest01, 5, 10)
	fmt.Println(byteSliceTest01)
	fmt.Println(byteSliceAfterAppend)
}
