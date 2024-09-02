package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./bufio.go")
	if err != nil {
		fmt.Println(err)
	}

	reader := bufio.NewReader(file)
	var recByteSlice []byte = make([]byte, 5)
	var content []byte
	for {
		_, err := reader.Read(recByteSlice) // Read() 方法两个返回值，第一个是本次读取的 []byte 长度
		if err == io.EOF {
			break
		}
		// 这里 recByteSlice... 后面三个点的意思是，将 []byte 打散追加给 content，需要 content 接收 append 方法的返回值
		content = append(content, recByteSlice...)
	}
	fmt.Println("文件内容已全部 append 到 content 切片...")
	fmt.Println(string(content))

}
