package main

import (
	"fmt"
	"os"
)

//func main() {
//	//file, err := os.Create("./writeTest.txt")
//	file, err := os.OpenFile("writeTest.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
//	if err != nil {
//		fmt.Println(err)
//	}
//	defer file.Close()
//	// 使用 file.Write 方法，需要传入一个 []byte 切片
//	var logWriteSlice []byte = []byte("RDS Service Status ====> 2022/07/28 14:40:50.637032 C:/Users/ethan/go/src/awesomeProject2/log1.go:41: 当前日志打印的位置在 log.txt 中...\n")
//	file.Write(logWriteSlice)
//	// 使用 file.Writestring，直接传入字符串
//	file.WriteString("WriteString1 ...\n")
//	file.WriteString("WriteString2 ...\n")
//
//}

func main() {
	file, err := os.OpenFile("fileWrite.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	file.WriteString("abcdefg\n")
}
