package main

import "fmt"

func f1() {
	defer func() {
		// 使用 recover() 恢复程序，即使异常也不中断程序运行，
		if err := recover(); err != nil {
			fmt.Println(err) // exception 打印异常
		}
		fmt.Println("finally") // finally 无论是否异常，都会执行
	}()
	fmt.Println("f1 program completed...")
}
func f2() {
	defer func() {
		// 使用 recover() 恢复程序，即使异常也不中断程序运行，
		if err := recover(); err != nil {
			fmt.Println(err) // exception 打印异常
		}
		fmt.Println("finally") // finally 无论是否异常，都会执行
	}()
	panic("f2 主动抛出异常")
	fmt.Println("f2 program completed...") // 不会执行

}
func f3() {
	fmt.Println("f3 program completed...")

}

func main() {
	f1()
	f2()
	f3()
}
