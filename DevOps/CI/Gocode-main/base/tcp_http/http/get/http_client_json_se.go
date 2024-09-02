package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type result struct {
	Args    string            `json:"args"`
	Headers map[string]string `json:"headers"`
	Origin  string            `json:"origin"`
	Url     string            `json:"url"`
}

func main() {
	resp, err := http.Get("http://httpbin.org/get")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	// json.Unmarshal  反序列化，映射赋值给 result 结构体
	var res result
	_ = json.Unmarshal(body, &res)
	fmt.Printf("%#v", res)
}
