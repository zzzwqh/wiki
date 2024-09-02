package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.topgoer.cn/docs/golang//1007")
	if err != nil {
		fmt.Println("Http Get Error: ", err)
	}
	defer resp.Body.Close()

	var buf [1024]byte
	for {
		n, err := resp.Body.Read(buf[:])
		// 打印 buf[:n] 需要放在卫戍语句的前面，否则最后一部分可能打印不全哦
		fmt.Println(string(buf[:n]))
		// 卫戍语句做判断，如果读到结尾就 Break 循环
		if err == io.EOF {
			fmt.Println("读取完了", err)
			break
		}
	}
	fmt.Println(cap(buf), len(buf))

}
