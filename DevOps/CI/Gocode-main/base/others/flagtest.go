package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

// 需要在命令行获取的参数
var name string
var age int
var married bool
var delay time.Duration

// 不使用原生的 flag.CommandLine，新构造一个相同类型的 *flag.FlagSet 对象
var cmdLine *flag.FlagSet

func init() {
	// 这里就类似 logger := log.New(out io.Writer, prefix string, flag int)
	// 自定义定制的好处就是，你的定制完全不会影响到全局变量 flag.CommandLine
	cmdLine = flag.NewFlagSet("Question", flag.ExitOnError)
	cmdLine.Usage = func() {
		// os.Args[0] 即命令源代码文件执行时，输入命令行的第一个参数，也就是命令源代码文件本身的名字
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		// 输出默认的用法
		flag.PrintDefaults()
	}
	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&delay, "d", 0, "延迟的时间间隔")
}

func main() {
	// 如果使用全局变量 flag.CommandLine 这里要写成 flag.Parse()
	// 查看 flag.Parse() 的源码，其实也是返回 CommandLine.Parse(os.Args[1:]) 即解析脚本后的参数
	cmdLine.Parse(os.Args[1:])
	fmt.Println(name, age, married, delay)
}
