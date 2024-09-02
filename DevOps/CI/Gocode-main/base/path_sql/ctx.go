package main

import (
	"fmt"
)

func f2() {
SwitchStatement:
	switch 1 {
	case 1:
		fmt.Println(1)
		for i := 0; i < 10; i++ {
			break SwitchStatement
		}
		fmt.Println(2)
	}
	fmt.Println(3)
}

func f3() {
OuterLoop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("i=%v, j=%v\n", i, j)
			continue OuterLoop // 这里其实写 break 是一个效果，退出当前层循环
		}
	}
}

func f4() {
	i := 0
Start:
	fmt.Println(i)
	if i > 2 {
		goto End // goto 语句必须指定 Label
	} else {
		i += 1
		goto Start
	}
End:
}
func main() {
	f4()
}
