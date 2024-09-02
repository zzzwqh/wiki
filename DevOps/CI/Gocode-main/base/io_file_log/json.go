package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string   `json:"name"`
	Age   int      `json:"age"`
	Wife  string   `json:"-"`          // 这样表示，不将该字段做序列化输出
	Hobby []string `json:",omitempty"` // 这样表示，如果该字段有值则做序列化输出，如果没有则不做序列化输出
	// Hobby []string `json:"hobby,omitempty"` // 如果想要序列化后的 Hobby 小写，需要这样写
}

func main() {
	// Hobby 没有赋值的结构体，序列化不会有 Hobby 字段
	var one = Person{Name: "ethan", Age: 12, Wife: "may"}
	// Hobby 有赋值的结构体
	var two = Person{Name: "may", Age: 13, Wife: "ethan", Hobby: []string{"zzz", "ddd"}}

	// 1、Marshal 把对象转成 Json 字符串
	res01, err := json.Marshal(one)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(res01))
	res02, err := json.Marshal(two)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(res02))
	// 2、Unmarshal 把 Json 字符串转成对象
	var three Person
	stringOfJson := "{\"name\":\"ethan\",\"age\":12}"
	json.Unmarshal([]byte(stringOfJson), &three)
	fmt.Println(three)
}
