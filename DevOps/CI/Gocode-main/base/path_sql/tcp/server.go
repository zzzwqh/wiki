package main

import (
	"bytes"
	"fmt"
	"net"
)

func handleConn(conn net.Conn) {
	defer conn.Close()
	// 这里的 for 循环是循环处理 Client 端给到的数据
	for {
		// 读取客户端数据
		var buf [1024]byte
		n, err := conn.Read(buf[:]) // 返回两个值，n 字节长度，err 错误
		if err != nil {
			fmt.Println("读取客户端数据出错", err)
			return
		}
		fmt.Println("接收到的数据：", string(buf[:n]))
		// 将接收到的数据处理后，写回 TCP Conn 连接返回给客户端
		conn.Write(bytes.ToUpper(buf[:n]))
	}
}

func main() {
	// 启动监听端口
	listener, err := net.Listen("tcp", "0.0.0.0:12345")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Service running on 0.0.0.0:12345 ...")
	for {
		// 用 for 循环等待连接， listener.Accept() 可以返回连接对象 conn
		conn, err := listener.Accept() // 如果没有连接会卡主，如果有人连接，会继续往下走
		if err != nil {
			fmt.Println("建立连接出错", err)
		}
		fmt.Println("当前服务端地址", conn.LocalAddr())
		fmt.Println("当前客户端请求地址", conn.RemoteAddr())
		go handleConn(conn) // 获取到了连接后，使用 goroutine 去处理连接
	}
}
