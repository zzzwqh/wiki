package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	// 只读方式打开当前文件
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}
	// 定义一个接受内容的 content
	var content []byte
	for {
		var receiveByteSliceBuf = make([]byte, 3)
		_, err := file.Read(receiveByteSliceBuf)
		if err == io.EOF {
			break
		} else if err != nil {
			println("读取文件错误：", err)
		}
		// 这里的是把读取到的长度为 3 的 receiveByteSliceBuf []byte 切片，追加到 content []byte 切片中
		//需要将子切片打散，所以写作 receiveByteSliceBuf...
		content = append(content, receiveByteSliceBuf...)
	}
	fmt.Printf("读取文件 %v 结束... 文件内容如下：\n", file.Name())
	fmt.Println(string(content))
}
