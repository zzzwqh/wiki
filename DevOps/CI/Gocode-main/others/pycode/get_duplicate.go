package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var strList = make([]rune, 1)
	reader := bufio.NewReader(os.Stdin)
	inputStr, _ := reader.ReadString('\n')
LOOP:
	for _, v := range inputStr {
		for _, j := range strList {
			if v != j {
				strList = append(strList, v)
			} else {
				fmt.Println(string(v))
				break LOOP
			}
		}
	}
}
