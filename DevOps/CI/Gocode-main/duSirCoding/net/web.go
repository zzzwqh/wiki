package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Client Get
// Client Post
// Web Server

func main() {
	/*
		 使用 handlerFunc 定义接口路径以及 handler，handler 就是接口对应的处理方法
			参数 1 接口路径
			参数 2 handler
	*/
	http.HandleFunc("/req/post", dealPostReqHandler)
	http.HandleFunc("/req/get", dealGetReqHandler)
	http.ListenAndServe(":10001", nil)
}

func dealPostReqHandler(w http.ResponseWriter, r *http.Request) {
	// 获取请求信息
	data, _ := ioutil.ReadAll(r.Body)
	var param = new(struct {
		Name string `json:"name"`
	})
	json.Unmarshal(data, param)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("接收到的名字是 %v", param.Name)))

}

func dealGetReqHandler(w http.ResponseWriter, r *http.Request) {

}
