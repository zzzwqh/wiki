package main

import (
	"fmt"
	"strings"
)

func main() {
	// 测试数据，一共四个 Ip 地址，我们想把分成四个，怎么分？
	Hosts := "15.227.30.1 15.227.30.2       15.227.30.4    15.227.30.3"
	// 用 strings.Split 指定分隔符，不能匹配多个
	countSplit01 := strings.Split(Hosts, " ")
	fmt.Println(countSplit01)
	// 可以看到切片的长度为 13
	fmt.Println(len(countSplit01))

	// 可以发现并不好用，Split 会将 str 中的 sep 去掉，而 SplitAfter 会保留 sep
	fmt.Println(strings.Split("a,b,c,d,e,d,c,b,a", "c"))           // [a,b, ,d,e,d, ,b,a]
	fmt.Println(len(strings.Split("a,b,c,d,e,d,c,b,a", "c")))      // 3
	fmt.Println(strings.SplitAfter("a,b,c,d,e,d,c,b,a", "c"))      // [a,b,c ,d,e,d,c ,b,a]
	fmt.Println(len(strings.SplitAfter("a,b,c,d,e,d,c,b,a", "c"))) // 3

	//
	// 字符串的拼接
	s := []string{"hello", "world", "ethanz"}
	fmt.Println(strings.Join(s, "-")) // hello-word-ethanz
}
