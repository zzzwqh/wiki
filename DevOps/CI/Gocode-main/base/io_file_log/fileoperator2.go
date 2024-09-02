package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	content, err := ioutil.ReadFile("./log.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(content))
}
