package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
)

// Exec4NoWaitRes 功能是，执行命令，不获取结果，不等待命令结束
func Exec4NoWaitRes() {
	// 只执行命令不获取结果，不使用 Wait 等待 Start 结束，命令会执行吗？ 会，&& 前的命令执行执行成功，才会执行 && 后的命令
	cmd := exec.Command("sh", "-c", "sleep 5 && date >> ./cmdText") // 使用 go run 构建执行，程序退出后，cmdText 文件仍然会出现
	err := cmd.Start()
	if err != nil {
		fmt.Println("命令执行失败", err)
		return
	}
	fmt.Println("命令执行成功")
}

// Exec4NoRes 功能是执行命令，不获取结果，等待命令结束
func Exec4NoRes() {
	// 这里 Start() 执行命令， Wait() 等待命令执行完成
	// 这里用 cmd.Run() 替换 Start() 和 Wait() 效果相同
	cmd := exec.Command("sh", "-c", "sleep 5 && date >> ./cmdText")
	err := cmd.Start()
	// err := cmd.Run()
	if err != nil {
		fmt.Println("命令执行失败", err)
		return
	}
	cmd.Wait()
	fmt.Println("命令执行成功")
}

// Exec4CorrectRes 功能是执行命令，获取 [正确] 结果，等待命令执行完成
func Exec4CorrectRes() {
	cmd := exec.Command("sh", "-c", "sleep 3 && cat /etc/redhat-release ; cat /etc/pam.d/")
	// cmd.OutPut() 只会打印出执行命令后输出的正确的内容（如果没有正确的就无输出）
	res, err := cmd.Output()
	if err != nil {
		fmt.Println("Exec4CorrectRes() Error:", err)
		// return	// 如果打开 return ，那么执行命令获得到的 [正确] 内容也不会打印了，只会打印 Exec4CorrectRes() Error: exit status 1
	}
	fmt.Println(string(res))
	/*
		这个函数会打印如下内容
		Exec4CorrectRes() Error: exit status 1
		CentOS Linux release 7.6.1810 (Core)
	*/
}

// Exec4AllRes 功能是执行命令，获取 [正确/错误] 的结果，等待命令执行完成
func Exec4AllRes() {
	cmd := exec.Command("sh", "-c", "sleep 3 && cat /etc/redhat-release ; cat /etc/pam.d/")
	res, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Exec4AllRes() Error:", err)
		//	return	// 如果打开 return ，那么执行命令获得到的 [正确/错误] 内容也不会打印了，只会打印 Exec4AllRes() Error: exit status 1
	}
	fmt.Println(string(res))
	/*
		这个函数会打印如下内容
		Exec4AllRes() Error: exit status 1
		CentOS Linux release 7.6.1810 (Core)
		cat: /etc/pam.d/: Is a directory
	*/
}

// Exec4SplitRes 功能是执行命令，获取 [正确/错误] 的结果并将其拆分，等待命令执行完成
func Exec4SplitRes() {
	cmd := exec.Command("sh", "-c", "sleep 3 && cat /etc/redhat-release ; cat /etc/pam.d/")
	// 1. 定义两个缓冲区（内存开辟 buffer 缓冲区），接收 cmd.Stdout 和 cmd.Stderr
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	// 2. 如果想输出在控制台，那么使用如下代码
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr
	// 上述 Stdout/Stderr 必须在 Run() 调用前配置
	err := cmd.Run()
	if err != nil {
		log.Print("Error running command 【", cmd, "】: ", err, "\n")
	}
	fmt.Println(string(stdout.Bytes())) // bytes.Buffer 对象可以使用 Bytes() 方法获取字节切片，然后用 string 转换类型
	fmt.Println(stderr.String())        // bytes.Buffer 对象也有方法 String() 可以获取 String 类型，相对上面的方法更加方便
}

// Exec4SplitRes2File 功能是执行命令，获取 [正确/错误] 的结果并将其拆分，分别输出到指定的日志文件中，等待命令执行完成
func Exec4SplitRes2File() {
	cmd := exec.Command("sh", "-c", "sleep 3 && cat /etc/redhat-release ; cat /etc/pam.d/")
	// 我们可以使用 log 模块产生一个日志文件，当然自己创建一个 *File 类型对象也没问题
	logger := log.New(os.Stdout, "自定义的 Logger 对象写日志 ===> ", log.LstdFlags)
	logger.SetPrefix("Exec4SplitRes2File 模块日志:")
	logger.SetFlags(log.Lshortfile | log.Lmicroseconds | log.Ldate)
	if logFile, err := os.OpenFile("./Exec4SplitRes2File.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666); err == nil {
		logger.SetOutput(logFile)
	} else {
		fmt.Println("Exec4SplitRes2File 模块日志文件创建/读写错误：", err)
	}
	// 只要保证都实现了  io.Writer 即可
	cmd.Stdout = logger.Writer()
	cmd.Stderr = logger.Writer()
	err := cmd.Run()
	if err != nil {
		fmt.Print("Error running command 【", cmd, "】: ", err, "\n") // 因为命令有错误输出，所以这句被打印出来是避免不了的
	}
}

// Exec4Pipe 功能是，用管道接收另外一个命令的输出，再执行命令处理
func Exec4Pipe() {
	cmd1 := exec.Command("sh", "-c", "ps -ef")
	cmd2 := exec.Command("sh", "-c", "grep mysql")
	cmd2.Stdin, _ = cmd1.StdoutPipe()
	cmd2.Stdout = os.Stdout
	cmd2.Start()
	cmd1.Run()
	cmd2.Wait()
	// 其实在 Linux 命令行直接写管道符就可以，记得写 sh -c
	// cmd := exec.Command("sh", "-c", "ps aux | grep mysql")
	// cmd.Stdout = os.Stdout	// 这里只打印了 Stdout，当然也可以打印 Stderr，不赘述
	// cmd.Run()

}

// Exec4GetEnv 功能是，配置程序级别的环境变量
func Exec4GetEnv() {
	os.Setenv("author", "ethan")
	// 1. 正确写法（推荐写法）,os.ExpandEnv() 会将传入的 String 类型参数中的 $var 全部转换成变量值
	cmd := exec.Command("sh", "-c", os.ExpandEnv("echo $author $USER"))
	// 2. 正确写法，可以直接在命令行中写 $key ，可以识别刚刚 os.Setenv() 配置的变量，当然也可以识别原有系统变量（推荐写法）
	// cmd := exec.Command("sh", "-c", "echo $author; echo $USER")

	// 3. 无效/错误写法，这种写法没有打印出变量值，是有问题的写法
	// cmd := exec.Command("sh", "-c", "echo", os.ExpandEnv("$author"))

	outputStr, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("cmd.CombineOutput Error :", err)
	}
	fmt.Printf("%s", outputStr)
}

func main() {
	// Exec4NoWaitRes()
	// Exec4NoRes()
	// Exec4CorrectRes()
	// Exec4SplitRes2File()
	// Exec4Pipe()
	Exec4GetEnv()
}
