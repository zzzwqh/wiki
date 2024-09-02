package main

import (
	"encoding/json"
	"fmt"
)

// Map 数据类型序列化
func main() {
	// 首先是序列化 MAP => JSON
	var mapJson = map[string]interface{}{"name": "小王", "address": "earth", "Wife": "May"}
	b, err := json.Marshal(mapJson)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))

	// 然后是反序列化 JSON => MAP
	var recMap map[string]interface{}
	var jsonString string = string(b)
	json.Unmarshal([]byte(jsonString), &recMap)
	// json.Unmarshal(b, &recMap)	其实直接传入 b 就行，但是模拟用上述代码
	fmt.Println(recMap)

	for key, value := range recMap {
		fmt.Println(key, "=>", value)
	}
}
