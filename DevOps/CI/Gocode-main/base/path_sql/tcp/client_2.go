package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:12345")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	for {
		// 向 Server 发送数据
		time.Sleep(time.Second)
		conn.Write([]byte("Normal Thing"))
		// 接收 Server 端数据
		var buf4Cli [1024]byte
		n, _ := conn.Read(buf4Cli[:])
		fmt.Println("Receive from server :", string(buf4Cli[:n]))
	}
}
