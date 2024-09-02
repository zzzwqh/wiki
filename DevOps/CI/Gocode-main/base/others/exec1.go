package main

import (
	"fmt"
	"os"
)

func init() {
	fmt.Println("init....")
}
func main() {
	err := fileClose()
	fmt.Println(err)
}
func fileClose() (err error) {
	file, err := os.Create("test.txt")
	defer file.Close()
	// Do stuff
	return file.Sync()
}
