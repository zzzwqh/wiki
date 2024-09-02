package main

import (
	"flag"
	"fmt"
)

func main() {

	flagTest := flag.String("path", "/home/admin/monitor.yml", "flag.exe --path [$path]")
	flag.Parse()
	fmt.Println(*flagTest)
}
