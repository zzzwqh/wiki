package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Info
}
type Info struct {
	Account int
	Addr    string
}

func main() {
	// 1. 使用 reflect.ValueOf 可以获取变量值，传递给 ValueOf 的参数确保是一个地址
	user := User{Id: 1, Name: "root", Info: Info{Account: 500000, Addr: "Earth"}}
	val := reflect.ValueOf(&user)
	fmt.Println(val.Kind()) // Ptr
	// 虽然能够获取变量值，但是这个变量值的类型是 reflect.Value 类型，并不是 User 类型
	fmt.Println("================== reflect.ValueOf() ========================")
	fmt.Printf("%T, %v\n", val, val) // reflect.Value  &{1 root {500000 Earth}}

	// 2. 那么怎么把 Value 的 Type 恢复成 User？ 用 reflect.ValueOf().Interface() 先将 value 转成接口类型，然后使用类型断言
	// 如果不是结构体，将会方便很多，例如字符串类型，reflect.ValueOf().String()
	correctTy4UserIns := val.Interface().(*User)
	fmt.Println("=========== reflect.ValueOf().Interface() ===================")
	fmt.Printf("%T %v\n", correctTy4UserIns, correctTy4UserIns) // *main.User &{1 root {500000 Earth}}

	// 3. 获取运行时变量的信息
	// Elem returns the value that the interface v contains or that the pointer v points to.
	// It panics if v's Kind is not Interface or Pointer. It returns the zero Value if v is nil.
	// 可以清楚，使用 Elem() 的前提 v.Kind() 要么是 Interface 要么是指针（ptr），否则会 Elem() 引起 panic
	// 至于 val.Kind() 可以看 22 行代码，如果不取 &user 地址，val 将会是 struct 类型，从而导致 panic
	elem := val.Elem()
	fmt.Printf("%v\n", elem) // {1 root {500000 Earth}}

	fmt.Println("=========== reflect.ValueOf().elem().Type() ===================")
	elemType := elem.Type()
	fmt.Printf("%v\n", elemType) // main.User

	fmt.Println("=========== reflect.ValueOf().elem().Kind() ===================")
	elemKind := elem.Kind()
	fmt.Printf("%v\n", elemKind) // struct

	fmt.Println("=========== reflect.ValueOf().elem().NumField() ===================")
	numField := elem.NumField()
	fmt.Printf("%v\n", numField) // 3
	fmt.Println("=========== reflect.ValueOf().elem().Field(i) ==== 当前 struct 的字段值（返回类型是 reflect.Value） ===============")
	fmt.Println("=========== reflect.ValueOf().elem().Field(i).Interface() ==== 当前 struct 的字段值（返回类型不是 reflect.Value） ===============")
	fmt.Println("=========== reflect.ValueOf().elem().Field(i).Type() ==== 当前 struct 字段类型（返回 reflect.Type 类型）===============")
	fmt.Println("=========== reflect.ValueOf().elem().Type().Field(i) ==== 当前 struct 字段信息（返回 StructField 类型） ===============")
	fmt.Println("=========== reflect.ValueOf().elem().Type().Field(i).Name ==== 当前 struct 字段名字（String 类型）===============")
	fmt.Println("=========== reflect.ValueOf().elem().Type().Field(i).Type ==== 当前 struct 字段类型（reflect.Type 类型）===============")
	for i := 0; i < numField; i++ {
		fmt.Println("第", i, "个字段获取 elem.Field()")
		field := elem.Field(i)
		fmt.Printf("%T,%v\n", field, field)
		fieldName := elem.Type().Field(i).Name // reflect.ValueOf().elem().Type().Field(i).Name 获取字段名字
		fieldType := field.Type()
		fieldValue := field.Interface() // reflect.ValueOf().elem().field(i).Interface() 将值的类型返回，并转成了具体类型
		fmt.Printf("%d: %s %s = %v (类型:%T)\n\n", i, fieldName, fieldType, fieldValue, fieldValue)
		fmt.Println(reflect.TypeOf(fieldValue)) // 可以观察 fieldValue 也并非 Interface 类型，而是具体类型
	}

	fmt.Println("=========== reflect.ValueOf().elem().FieldByIndex() ==== 结构体嵌套结构体，深度遍历 ===============")
	fmt.Println(elem.FieldByIndex([]int{2, 1})) // 打印 Earth，打印了第 2 个字段（Info）的第 1 个字段（Addr）的值，也就是 Earth
}
