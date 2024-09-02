package main

import "fmt"

func main() {
	var strTest01 string = "I'm Ethan ,an Operation Engineer"
	var strTest02 string = "我是一个 Linux 运维工程师"
	// _ 省略的是 Key，也就是字符串的下标
	for _, value := range strTest01 {
		fmt.Println(string(value))
		fmt.Printf("%c\n", value)
	}
	// 将字符串转换成 rune 切片，这样就可以取出并打印完整的中文（4个Bytes）
	runeForStrTest02 := []rune(strTest02)
	// _ 省略 Key，这里是切片的下标
	for _, value := range runeForStrTest02 {
		fmt.Printf("%c\n", value)
	}
}
