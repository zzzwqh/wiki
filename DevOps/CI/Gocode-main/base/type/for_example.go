package main

import (
	"EthanCode/base/type/user"
	"fmt"
)

func main() {
	user.ShowAge()
	// 1、使用方式一  for 条件 三段都写，日常写法
	for i := 0; i < 10; i++ {
		fmt.Println("I 循环此次索引值", i)
	}

	// 2、 使用方式二	for 条件 省略第一段，索引的作用域变大
	var j = 0
	for ; j < 10; j++ {
		fmt.Println("J 循环此次索引值", j)
	}
	fmt.Println(j)

	// 3、 使用方式三	for 条件 省略第三段，自增写在循环体内
	for k := 0; k < 10; {
		fmt.Println("K 循环此次索引值", k)
		k++
	}

	// 4. 使用方式四 省略第一段和第三段
	// 等同于其他语言的 while 循环
	var m = 0
	//  for ;m<10; 省略了 ; 符号，就很像 while 循环
	for m < 10 {
		fmt.Println("M 循环此次索引值", m)
		m++
	}

	// 5. 省略三段
	// for ;; 	省略了 ; 符号
	// 死循环也可以写成  for true {}
	for {
		fmt.Println("死循环啊喂...")
	}

}
