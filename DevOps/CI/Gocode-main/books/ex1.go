package main

import (
	"fmt"
)

func main() {
	var array [10]int

	var slice = array[5:6]

	fmt.Println("lenth of slice: ", len(slice))
	fmt.Println("capacity of slice: ", cap(slice))
	fmt.Println(&slice[0] == &array[5])
	orderLen := 5
	order := make([]uint16, 2*orderLen)
	order = []uint16{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(order)
	pollorder := order[:orderLen:orderLen]
	lockorder := order[orderLen:][:orderLen:orderLen]
	fmt.Println(pollorder)
	fmt.Println(lockorder)
	fmt.Println("len(pollorder) = ", len(pollorder))
	fmt.Println("cap(pollorder) = ", cap(pollorder))
	fmt.Println("len(lockorder) = ", len(lockorder))
	fmt.Println("cap(lockorder) = ", cap(lockorder))

}
