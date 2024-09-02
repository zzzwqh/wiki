package main

import (
	"fmt"
	"os"
)

func main() {
	//3 Fscan系列---》 Fsacnln Fscanf---》实现io.Reader接口的变量中读-->不仅仅能从控制台读，还能从文件中读
	var name string
	fmt.Fscanln(os.Stdin, &name)
	fmt.Println("名字是：", name)
	//从文件中读
	var s = "aaa"
	file, _ := os.Open("./notes.txt") // 打开文件
	fmt.Fscanln(file, &s)
	fmt.Println(s)
	file.Close()
}
