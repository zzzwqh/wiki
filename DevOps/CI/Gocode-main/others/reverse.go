package main

import (
	"fmt"
	"reflect"
)

func main() {
	var str = "我是一个 Golang 语言爱好者"
	fmt.Println(str)
	var arrayIns = [...]int{1, 2, 3, 4}
	fmt.Println(reflect.TypeOf(&arrayIns))

	runeIns := []rune(str)
	fmt.Println(string(runeIns))
}
