package main

import (
	"fmt"
	"os"
)

func main() {
	// 利用 fmt.Fprintf 函数写文件
	file, err := os.OpenFile("./writeTest_Fprintf.txt", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println(err)
	}
	serviceName := "RDS"
	fprintTest01 := "%v writeTest For Fprintf Func\n"
	fmt.Fprintf(file, fprintTest01, serviceName)

	fprintTest02 := []string{"OKay you are the hero", "Noah fang boat", "Go out the house"}
	for _, value := range fprintTest02 {
		fmt.Fprintln(file, value)
	}
	defer file.Close()
}
