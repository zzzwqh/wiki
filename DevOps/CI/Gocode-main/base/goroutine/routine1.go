package main

import (
	"fmt"
)

func main() {

	chanForSqrt := make(chan interface{})
	chanForPow := make(chan interface{})

	go calSqrt(123, chanForSqrt)
	go calPow(222, chanForPow)
	sqrtOnMainGoroutine, powOnMainGoroutine := <-chanForSqrt, <-chanForPow
	fmt.Println(sqrtOnMainGoroutine, powOnMainGoroutine)

}

func calSqrt(num int, result chan interface{}) {
	sum := 0
	for num != 0 {
		digit := num % 10
		sum += digit * digit * digit
		num = num / 10
	}
	result <- sum
}
func calPow(num int, result chan interface{}) {
	sum := 0
	for num != 0 {
		digit := num % 10
		sum += digit * digit
		num = num / 10
	}
	result <- sum
}
