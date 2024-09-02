package main

import (
	"fmt"
	"strings"
)

func isSeparator(s rune) bool {
	if s == 'n' {
		return true
	}
	return false
}

func main() {
	Hosts := "15.227.30.1 15.227.30.2       15.227.30.4    15.227.30.3"
	// 用 strings.Fields 分割 ，是按照 1：n个空格来分割
	Iplist := strings.Fields(Hosts)
	fmt.Println(Iplist)
	// 可以看到切片的长度是 4 符合想要的结果
	fmt.Println(len(Iplist))

	// 指定分隔符，并且按照 1：n个分隔符来分割
	fmt.Println(strings.FieldsFunc("widuunhellonnnword", isSeparator))      // [widuu hello word] 根据 n 字符分割
	fmt.Println(len(strings.FieldsFunc("widuunhellonnnword", isSeparator))) // [widuu hello word] 根据 n 字符分割，长度是 3

}
