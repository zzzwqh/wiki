package main // 指定包名，任何 go 文件都需要，

import "fmt" // 导入内置包

func main() { // 定义一个函数名
	fmt.Println("Hello,World") // 想打印字符串必须用双引号
	fmt.Println('a')           // 单引号表示 ascii 的值

}

// Go 的入口，必须是 main 包下的 main 函数
// a... 是 goland 的提示
