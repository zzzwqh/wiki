package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "abcdefg"
	sep := "d"
	i := strings.Index(s, sep)
	fmt.Println(i)
	fmt.Println(s[i+1:])
}
