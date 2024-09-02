package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

func main() {
	str1 := "abcde"
	str2 := "55555"
	fmt.Println(str1, str2)
	fmt.Println("==========")

	var builder strings.Builder
	fmt.Fprintln(&builder, "qwert")
	fmt.Fprintln(&builder, "12345")
	fmt.Println(builder.String())

	var buf bytes.Buffer
	fmt.Fprintln(&buf, "zxcv")
	fmt.Fprintln(&buf, "56789")
	//fmt.Println(buf.String())
	fmt.Println("===========")
	//reader := bufio.NewReader(&buf)
	//for { // 循环读取所有行
	//	line, _, err := reader.ReadLine() // 按行读取文件
	//	if err == io.EOF {
	//		break
	//	}
	//	fmt.Println(string(line))
	//}
	fmt.Println("===========")
	reader := bufio.NewReader(&buf)
	var recByteSlice = make([]byte, 1)
	_, err := reader.Read(recByteSlice)
	if err != nil {
		fmt.Println("reader.Read(recByteSlice) 错误：", err)
	}
	for {
		reader.Read(recByteSlice)
		fmt.Println(string(recByteSlice))
		if err == io.EOF {
			break
		}
	}
}
