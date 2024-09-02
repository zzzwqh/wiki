package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	a := adder()
	fmt.Println(a(1)) // 输出：1
	fmt.Println(a(2)) // 输出：3
	fmt.Println(a(3)) // 输出：6
}
