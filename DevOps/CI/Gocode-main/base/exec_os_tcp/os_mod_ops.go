package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"syscall"
	"time"
)

// osNormalFunc 关于 os 提供的基础操作
func os4NormalFunc() {
	fmt.Println(os.Hostname())                // 打印主机名
	fmt.Println(os.Getuid())                  // 返回调用进程的实际用户 ID
	fmt.Println(os.Getgid())                  // 返回调用进程的实际组 ID
	fmt.Println(os.Getegid())                 // 返回调用进程的有效组 ID
	fmt.Println(os.Getwd())                   // 返回当前目录
	fmt.Println(os.Environ())                 // 返回当前程序中的全部变量（包括从系统继承的，以及程序内自定义的）
	fmt.Println(os.Getpid())                  // 获取当前进程的 ID
	fmt.Println(os.Getppid())                 // 获取当前进程的父进程 ID
	fmt.Println(os.Getpagesize())             // 获取当前系统分页大小
	fmt.Println(os.Getenv("GOPATH"))          // 根据 Key 值获取变量的值
	fmt.Println(os.ExpandEnv("PATH = $PATH")) // 传入 ExpandEnv() 的字符串中 $var 格式的变量将会替换成变量值
	fmt.Println(os.TempDir())                 // 返回用于保管临时文件的默认目录，在 Linux 系统中其实就是 /tmp
	fmt.Println(os.Args)                      // Args 是个 []string 类型，保管了执行程序时，命令行的参数，其中 os.Args[0] 就是该程序文件名
}

// os4PathOrFile 关于 os 提供的文件和目录的操作
func os4PathOrFile() {
	fmt.Println(os.Getwd())
	os.Chdir("/etc/pam.d") //	切换当前程序执行的目录，os.Getwd() 在这句命令执行前后会有变化
	fmt.Println(os.Getwd())

	os.Mkdir("osDir", 0666)                       // 在当前程序的执行目录 /etc/pam.d/ 下创建了一个 osDir，如果文件夹存在，那么不会执行，error 类型返回值是 nil
	os.MkdirAll("/dirTest/ethan/mysql/bin", 0666) // 创建多级目录，权限都是 0666

	err := os.Remove("/dirTest/ethan/mysql") // 这种写法会删除最后一个 [目录]/[文件]，但是最后一个目录若不为空，会报错 remove /dirTest/ethan/mysql: directory not empty
	if err != nil {
		fmt.Println("os.Remove() Exec Error: ", err)
	}
	os.RemoveAll("/dirTest/ethan/mysql")               // 同上，会删除最后一个 [目录]/[文件]，即使最后一个目录不为空，也可以递归删除
	err = os.Rename("/dirTest/ethan", "/dirTest/noah") // 重命名 [文件]/[文件夹]
	if err != nil {
		fmt.Println("os.Rename() Exec Error: ", err)
	}

	// os.Create 创建文件，记得要关闭文件
	file, _ := os.Create("/dirTest/os_mod.log")
	defer file.Close()

	// os.Stat 获取文件信息，返回一个 FileInfo 接口类型，有很多方法
	fileInfo, _ := os.Stat("/dirTest/os_mod.log")
	fmt.Println(fileInfo.Name())    // 获取文件信息：文件名
	fmt.Println(fileInfo.Mode())    // 获取文件信息：文件权限
	fmt.Println(fileInfo.IsDir())   // 获取文件信息：是否是目录
	fmt.Println(fileInfo.Size())    // 获取文件信息：文件大小
	fmt.Println(fileInfo.ModTime()) // 获取文件信息：更改时间

	// os.Chmod 可以改变文件/目录的权限，会立即生效
	if err := os.Chmod("/dirTest/os_mod.log", 0777); err != nil {
		log.Fatal(err)
	}
	//
	os.Chown("/dirTest/os_mod.log", 998, 996)
	os.Chtimes("/dirTest/os_mod.log", time.Now().Add(-time.Hour), time.Now().Add(-time.Hour))
	fileInfo, _ = os.Stat("/dirTest/os_mod.log") // 上面 os.Stat 执行的时候，信息就已经获取固定了，所以如果想看到变化，需要重新执行 os.Stat()
	fmt.Println(fileInfo.Mode())
	fmt.Println(fileInfo.Sys()) // 输出的是一个集合，其中包含了文件的属主和属组信息
	fmt.Println(fileInfo.ModTime())

	res := fileInfo.Sys()
	fmt.Println(reflect.TypeOf(res)) // 可以观察到我们的 fileInfo.Sys 是啥类型，从而决定怎么取出属主的 ID
	/*
		下面代码行 fileInfo.Sys().(*syscall.Stat_t) 在 Windows 中无法编译，因为 *syscall.Stat_t 类型是属于 Linux 的 Golang SDK 的
		https://studygolang.com/topics/287/comment/778 中介绍了 Stat_t 结构体中包含如下字段，都可以打印出
		type Stat_t struct {
			Dev       uint64
			Ino       uint64
			Nlink     uint64
			Mode      uint32
			Uid       uint32
			Gid       uint32
			X__pad0   int32
			Rdev      uint64
			Size      int64
			Blksize   int64
			Blocks    int64
			Atim      Timespec
			Mtim      Timespec
			Ctim      Timespec
			X__unused [3]int64
		}
	*/
	// stat := fileInfo.Sys().(*syscall.Stat_t)
	// fmt.Println(timeSpecToTime(stat.Atim))
	// fmt.Println(timeSpecToTime(stat.Ctim))
	// fmt.Println(timeSpecToTime(stat.Mtim))
	// fmt.Println(stat.Uid) // 打印出属主
	// fmt.Println(stat.Gid) // 打印出属组

}

// timeSpecToTime 功能是用于 os4PathOrFile 函数中 *syscall.Stat_t 对象里的时间类型 syscall.TimeSpec 的类型转换
func timeSpecToTime(ts syscall.Timespec) time.Time {
	return time.Unix(int64(ts.Sec), int64(ts.Nsec))

}

// os4EnvFunc 关于 os 模块的设置/获取变量的操作
func os4EnvFunc() {
	// os.ExpandEnv() 会将传入的字符串中的 $var 自动传换成变量值（包括程序内设置的变量值）
	str := "Hello , $USER ,The server hostname is $HOSTNAME , Your Email is $EMAIL"
	fmt.Println(os.ExpandEnv(str)) // 输出结果: Hello , root ,The server hostname is master01 , Your Email is
	// os.SetEnv() 一般放在 init 函数执行
	os.Setenv("EMAIL", "wqh3456@126.com")
	fmt.Println(os.ExpandEnv(str)) // 输出结果: Hello , root ,The server hostname is master01 , Your Email is wqh3456@126.com

	fmt.Println(os.Getenv("$USER")) // 至于 GetEnv,只是传入 Key，获取 Value 的功能

	// 删除当前程序所有环境变量，包括从系统继承过来的环境变量（不会影响系统中的系统变量）
	os.Clearenv()

}

// getAllFileInPath 遍历当前路径的所有文件
func getAllFileInPath() {
	// 用到了 os 标准库中的 Getwd 方法，也就是获取当前路径
	dir, _ := os.Getwd()
	fmt.Println(dir)
	// 遍历其实是用了 path 标准库的 Walk 方法
	filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		fmt.Println(path)
		return err
	})
}

func main() {
	// 用来梳理 os 模块提供常用功能
	os4NormalFunc()
	// 用来测试 os 模块文件/目录相关功能
	// os4PathOrFile()
	// 用来测试 os 模块环境变量相关功能
	// os4EnvFunc()

	// 实践：遍历当前路径的所有文件
	// getAllFileInPath()

	defer fmt.Println("当执行 os.Exit() 时，即使被 defer 注册过的代码也不会被执行")
	os.Exit(0)
	// os.exit() 和 panic都能退出程序，但是使用上也是有区别的
	// os.Exit 函数可以让当前程序以给出的状态码 code 退出。一般来说，状态码 0 表示成功，非 0 表示出错。程序会立刻终止，并且 defer 的函数不会被执行
	// panic 可以触发 defer 延迟语句，panic 还可以被 recover 捕获处理

}
