package main

import (
	"log"
	"os"
)

func main() {
	// 如果想自己造出一个 logger 对象
	logger := log.New(os.Stdout, "自定义的 Logger 对象写日志 ===> ", log.LstdFlags)
	logger.Println("[info] xxxxx")
}
