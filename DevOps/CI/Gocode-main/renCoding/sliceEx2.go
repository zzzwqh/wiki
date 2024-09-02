package main

import "fmt"

func SliceRiseUsedPointer(s *[]int) {
	*s = append(*s, 0)
}
func SliceRise(s []int) {
	s = append(s, 0)
}
func main() {
	var s1 []int = []int{1, 2}
	SliceRise(s1)
	SliceRiseUsedPointer(&s1)
	fmt.Println(s1)
	var c chan int
	close(c)
	if c == nil {
		fmt.Println("只是声明，。")
	}
}

//====>  prometheus  监控思维导图
