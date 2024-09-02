package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./log1.go")
	if err != nil {
		fmt.Println(file)
	}
	defer file.Close()
	// 需要先打开文件，才能使用 bufio.NewReader，因为 bufio.NewReader 需要接受一个 file 对象
	reader := bufio.NewReader(file) // 新建一个 Reader 对象，可以用 reader.ReadLine 方法按行读取文件
	for {                           // 循环读取所有行
		line, _, err := reader.ReadLine() // 按行读取文件
		if err == io.EOF {
			break
		}
		fmt.Println(string(line))
	}
}
