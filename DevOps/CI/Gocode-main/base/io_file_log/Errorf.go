package main

import (
	"errors"
	"fmt"
)

func main() {
	var name string = "ethan"
	err1 := fmt.Errorf("person %v error", name)
	fmt.Printf(err1.Error())
	err2 := errors.New("person xxx error")
	fmt.Printf(err2.Error())

}
