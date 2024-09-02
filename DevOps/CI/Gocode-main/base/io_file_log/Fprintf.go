package main

import (
	"fmt"
	"os"
)

func main() {
	// 将内容输出到终端
	fmt.Fprintf(os.Stdout, "向终端输出的内容")
	// 打开一个新的文件 ./xx.txt 用 os.OpenFile 函数
	fileObj, err := os.OpenFile("./notes.txt", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("打开文件出错")
		fmt.Printf(err.Error())
	}
	// 可以向 Fprintf 传递一个打开文件对象，也可以传递一个变量值
	fprintVarTest1 := "ethan"
	fprintVarTest2 := "noah"
	// 向打开的文件句柄中写入内容
	fmt.Fprintln(fileObj, "向文件中输入内容", fprintVarTest1)
	fmt.Fprintf(fileObj, "向文件中输入内容 %v", fprintVarTest2)
}
