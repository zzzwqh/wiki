package main

import (
	"bufio"
	"fmt"
	"os"
)

//
//func main() {
//	// BufIo NewWriter
//	file, err := os.OpenFile("writeTest_bufio.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
//	if err != nil {
//		fmt.Println(err)
//	}
//	defer file.Close()
//	writer := bufio.NewWriter(file)
//	writer.WriteString("abcdefg\n")
//	// 必须要将内存缓冲区的数据，刷入硬盘
//	writer.Flush()
//
//}

func main() {
	file, err := os.OpenFile("writeTest_bufio.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString("sdfsdf\n")
	writer.Flush()
}
