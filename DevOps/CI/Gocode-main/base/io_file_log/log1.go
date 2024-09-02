package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// 1、打印普通类型的日志
	log.Println("this line is info log")
	// log.Panic 类型方法，会触发 Panic 导致程序直接结束，后面的不会运行
	// log.Panicln("this line is panic error log")
	// log.Fatal 类型方法，会触发 Fatal 导致程序直接结束，后面的不会运行
	// log.Fatalln("this line is fatal log")

	// 默认是 3 log.Flags()  log.Ldate | log.Ltime
	fmt.Println(log.Flags())

	// 2、重新设定 log.Flags log.Llongfile 是打印当前绝对路径 ， log.Lmicroseconds 是显示毫秒级的时间，log.Ldate 是显示日期
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	fmt.Println(log.Flags())
	log.Println("this log.Flags is 13")
	fmt.Println("===============")

	// 3、prefix  配置日志的前缀
	fmt.Println(log.Prefix())
	log.SetPrefix("RDS Service Status ====> ")
	log.Println("[info] starting...")

	// 4、log.SetOutput 配置日志的输出位置
	log.SetOutput(os.Stdout)
	log.Println("设置日志打印的位置...")

	// 5、文件实现了 io.Writer 接口，放一个文件对象
	file01, err := os.Create("./log.txt")
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(file01) // 需要放入一个实现了 io.Writer 接口的对象
	log.Println("当前日志打印的位置在 log.txt 中...")

	// 6、用 os.Create 每次都是创建一个新的 log.txt 文件
	file02, err := os.OpenFile("./log_OpenFile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(file02)
	log.Println("用 os.OpenFile 函数打开文件并写入日志...")

}

// 7、通常我们将日志的配置，写到 init() 函数中
func init() {
	log.SetPrefix("ServiceTest 服务日志:")
	log.SetFlags(log.Lshortfile | log.Lmicroseconds | log.Ldate)
	if logFile, err := os.OpenFile("servicetest.log ", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666); err == nil {
		log.SetOutput(logFile)
	} else {
		fmt.Println("ServiceTest 日志文件创建/读写错误：", err)
	}

}
