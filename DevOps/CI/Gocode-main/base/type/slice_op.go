package main

import "fmt"

func main() {
	// append 是否会影响原有 slice 的底层数组，答案：yes
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	k := append(s[:2], s[4:]...)
	z := append(k, s[4:]...)

	fmt.Println(cap(s), len(s), s)
	fmt.Println(cap(k), len(k), k)
	fmt.Println(cap(z), len(z), z)

	// RESULT:
	// 10 10 [1 2 5 6 7 8 9 10 9 10]
	// 10 8 [1 2 5 6 7 8 9 10]
	// 20 14 [1 2 5 6 7 8 9 10 7 8 9 10 9 10]
}
