package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "121.89.244.58:12345")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	for {

		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("Input Your Command: ")
		inputCmd, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println(err)
			return
		}
		conn.Write([]byte(inputCmd))
		var bufRec [1024]byte
		conn.Read(bufRec[:])
		fmt.Println(string(bufRec[:]))
	}
}
