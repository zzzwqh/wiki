package main

import (
	"fmt"
)

func main() {
	var variableTest01 bool = true
	fmt.Println(variableTest01)

	// Float 32 表示到小数点后 6 位 , Float 64 表示到小数点后 15 位
	var salaryTest1 float32 = 10.23125123123123123123
	var salaryTest2 float64 = 10.1231215215124123123123123123123123123123
	fmt.Println(salaryTest1, salaryTest2)
	// string
	var strTest1 = "str1" + "str2" + "st3"
	var strTest2 = `str1
str2
					str3`
	fmt.Println(strTest1)
	fmt.Println(strTest2)
	// byte & rune
	var b1 byte = 0   // 只能表示 0 - 255 assci 码
	var r1 rune = 123 // 可以表示任意四个字节的字符
	fmt.Println(b1, r1)

	// 常量定义格式
	const constName string = "constWqh"
	var varName string = "varWqh"
	const constAge = 19
	fmt.Println(constName, constAge, varName)
	// 常量一次定义赋值，就无法改变值，变量可以
	//constName = "zzz"  // ERROR
	varName = "zzz"

	// 多个常量的定义、iota 自增常量
	// 赋值 iota 每一行自增
	const (
		a = iota
		b = iota
		c
		d = 100
		e
		f = iota
		g
		h
	)
	fmt.Println(a, b, c, d, e, f, g, h)

	// 放到一起赋值 iota 却不自增
	const (
		j, k, l = iota, iota, iota
	)
	fmt.Println(j, k, l)

	// 只要不换行，就不自增
	const (
		m, o, v = iota, iota, iota
		n       = iota
		z       = iota
	)
	fmt.Println(m, o, v, n, z)

	//  iota + 1
	const (
		p = iota + 1
		q
	)
	fmt.Println(p, q)

}
