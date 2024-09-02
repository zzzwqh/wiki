package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"os/exec"
	"strings"
)

func process(conn net.Conn) {
	defer conn.Close()
	var buf [1024]byte
	for {
		n, _ := conn.Read(buf[:])
		// 当我从 Windows 的命令行输入命令时，换行符是 /r ，传入到 Server 端会报错，所以要将 /r 替换（清洗）
		// 另外这里的 buf[:n] 的 n 必须指定，不然会有很多空格传入 args，我用 Trim 都清洗不了,估计因为是 \r windows 换行符的原因
		cmd := exec.Command("/usr/bin/bash", "-c", strings.ReplaceAll(string(buf[:n]), "\r", ""))
		// 我们将 正确/错误 的执行结果输出，都传递到同一个缓冲区 stdout，如果想分开，那就定义两个 bytes.Buffer，分开接收结果输出
		var stdout bytes.Buffer
		cmd.Stdout = &stdout
		cmd.Stderr = &stdout
		err := cmd.Run()
		if err != nil {
			fmt.Println("cmd.Run Error 命令执行错误：", err)
			fmt.Println("Stderr：", stdout.String())
		}
		conn.Write([]byte("命令回执 => "))
		conn.Write([]byte(stdout.String()))
	}
}

func main() {
	log.Printf("Server Running on 0.0.0.0:12345...")
	listener, err := net.Listen("tcp", "0.0.0.0:12345")
	if err != nil {
		fmt.Println("net.Listen Error:", err)
	}
	defer listener.Close()

	//  for 循环等待连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept Error:", err)
		}
		fmt.Println("当前建立连接的客户端地址", conn.RemoteAddr())
		go process(conn)
	}
}
