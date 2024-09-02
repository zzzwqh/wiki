package main

import (
	"fmt"
	"math/rand"
	"time"
)

var mapTest map[string]interface{}

func main() {
	mapTest = make(map[string]interface{})

	mapTest["age"] = 16
	mapTest["name"] = "ethan"
	mapTest["hobby"] = []string{"sing", "rap", "basketball"}

	// 自定义函数，通过函数赋值，通过函数获取值
	setValue("address", "earth")
	addr := getValue("address")
	fmt.Println(addr)
	// time.Now().UnixNano() 纳秒，time.Now().Unix() 秒
	fmt.Println(time.Now().UnixNano())
	fmt.Println(time.Now().Unix())
	// 利用 timestamp 种下 Seed ，用来生成随机数
	rand.Seed(time.Now().UnixNano())
	// 循环向这个 mapTest 中赋值
	for i := 0; i < 11; i++ {
		// strconv.Itoa(int64) 可以用来将 int64 强制转换成 string 类型，rand.Intn(100) 生成 0-100 的随机数
		// setValue(strconv.Itoa(i), rand.Intn(100))
		// 使用 Sprintf 拼接 string 和 int 类型,并返回拼接值（fmt.Printf() 没有返回值）不建议使用 stu := "student" + strconv.Itoa(i)，如下代码更方便格式化
		stu := fmt.Sprintf("student%03d", i)
		setValue(stu, rand.Intn(100))
	}
	fmt.Println(mapTest)
}

func setValue(key string, value interface{}) {
	mapTest[key] = value
}
func getValue(key string) interface{} {
	return mapTest[key]
}
