package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   []byte("0.0.0.0"),
		Port: 12345,
		Zone: "",
	})
	if err != nil {
		fmt.Println("UDP 服务启动失败", err)
		return
	}
	defer listener.Close()
}
