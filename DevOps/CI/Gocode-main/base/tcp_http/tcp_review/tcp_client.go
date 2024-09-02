package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9199")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer conn.Close()
	fmt.Println("Server Addr :", conn.RemoteAddr())
	fmt.Println("Client Addr :", conn.LocalAddr())
	for {
		time.Sleep(time.Second)
		conn.Write([]byte("ls"))
	}

}
