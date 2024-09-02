package main

import (
	"EthanCode/base/goroutine/entity"
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	// 使用 reflect.ValueOf().Set*** 系列方法可以修改值
	strIns := "12345"
	val := reflect.ValueOf(&strIns) // 一定要加取地址符号，修改变量本身
	elem := val.Elem()
	elem.SetString("abcde") // 还有 SetInt 等待方法，不一一列举
	fmt.Println(elem)

	// 如果是结构体怎么修改？
	var dogIns = entity.Dog{Name: "ethan", Age: "3"}
	val4Dog := reflect.ValueOf(&dogIns)
	// 可以用如下方式去改
	val4Dog.Elem().FieldByName("Name").SetString("Noah")
	fmt.Println(val4Dog.Elem().FieldByName("Name"))
	// 如果不想一个个写，那么可以用下面的方式遍历哇，Elem() 是解引用 * 的，传入 reflect.ValueOf 函数中的是指针
	fmt.Println("先试试用 Reflect 遍历结构体字段 ===>")
	for i := 0; i < val4Dog.Elem().NumField(); i++ {
		fmt.Println(val4Dog.Type().Elem().Field(i).Name)
	}
	fmt.Println("有方法可以获取到字段，那么遍历修改 ===>")
	for i := 0; i < val4Dog.Elem().Type().NumField(); i++ {
		fieldName := val4Dog.Elem().Type().Field(i).Name

		val4Dog.Elem().FieldByName(fieldName).SetString("Test...") // 这里图个方便都赋值 Test...
		// 当然字段不一定都是 String 类型，这里可以写个函数用 Kind() 做判断，如果是 struct 就要递归，可参考 reflect_tag 中的代码，62 行
	}
	info, _ := json.MarshalIndent(dogIns, "", "\t")
	fmt.Println(string(info))
}
