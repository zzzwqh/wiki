package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i int = 1
	fmt.Printf("%T\n", i) // 底层使用反射查询出对应类型你

	var typ reflect.Type = reflect.TypeOf(i) // 实现反射底层，查询类型
	fmt.Println(typ)
}
