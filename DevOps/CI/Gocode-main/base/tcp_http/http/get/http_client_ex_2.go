package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.topgoer.cn/docs/golang//1007")
	if err != nil {
		fmt.Println("Http Get Method Error: ", err)
	}
	defer resp.Body.Close()

	fmt.Println(resp.Status)
	fmt.Println(resp.Header)
	// 和 ex_1 不同之处，只是这里用切片接收 Body 内容
	buf := make([]byte, 1024)
	for {
		//	换用初始化过的切片读内容
		n, err := resp.Body.Read(buf)
		fmt.Println(string(buf[:n]))
		if err != nil || err == io.EOF {
			fmt.Println("Response Read  Error: ", err)
			break
		}
	}
	fmt.Println(cap(buf), len(buf))
}
