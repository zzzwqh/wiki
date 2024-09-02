package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// 第一种，打印出的是字符 unnicode 编码
	var stringTest01 string = "hello"
	for i := 0; i < len(stringTest01); i++ {
		fmt.Println(stringTest01[i])
	}
	// 第二种，可以打出字符（根据 unnicode 编码强制转换成 string 类型），但是是按照字节的方式循环打印
	var stringTest02 string = "hello，中国"
	for i := 0; i < len(stringTest02); i++ {
		fmt.Println(string(stringTest02[i]))
	}

	fmt.Println("rune=================")
	// 1、通过 rune 切片构建字符串，把字符串做成 rune 切片，就可以将所有字符根据 unnicode 编码强制转换成对应的 string 类型字符
	var runeSlice01 []rune = []rune(stringTest02)
	for i := 0; i < len(runeSlice01); i++ {
		fmt.Printf("%c ", runeSlice01[i])
	}
	fmt.Println()

	fmt.Println("byte=================")

	// 2、通过 byte 切片构建字符串，byte 切片输出的一个字符 8 字节，所以无法正常输出中文
	var byteSlice01 []byte = []byte{104, 101, 108, 108, 111, 32}
	for i := 0; i < len(byteSlice01); i++ {
		fmt.Printf("%c ", byteSlice01[i])
	}
	strTest01 := string(byteSlice01)
	fmt.Printf(strTest01)
	fmt.Println()

	fmt.Println("count of string==========")

	// 3、字符串长度统计（utf8字符统计）
	var strTest02 string = string(runeSlice01)
	fmt.Println(strTest02)
	fmt.Println(utf8.RuneCountInString(strTest02))
	fmt.Println()

	fmt.Println("如何修改字符串=============")
	// 4、字符串变量定义赋值后，不可改变！！！
	strTest03 := "我是小王,ethanz"
	fmt.Println(strTest03)
	runeSlice02 := []rune(strTest03)
	runeSlice02[0] = '你'        // 要用 单引号 ！！ 就可以自动赋值为其 uft8 的编码值
	fmt.Println(runeSlice02[0]) // 看一下刚刚赋值的多少
	fmt.Println(string(runeSlice02))
	// 利用 rune 切片更改字符串后，再将其赋值回去
	strTest03 = string(runeSlice02)
	fmt.Println(strTest03)
}
