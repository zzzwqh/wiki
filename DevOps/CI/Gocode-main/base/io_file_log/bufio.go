package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	bufioTest()
}

func bufioTest() {
	reader := bufio.NewReader(os.Stdin) // 从标准输入生成读对象
	fmt.Println("请输入需要做免密的机器地址：")
	argHost, _ := reader.ReadString('\n') // 读到换行，argHost 是一个 String 类型的数据
	fmt.Printf("%T", argHost)
	argHost = strings.TrimSpace(argHost) // string.TrimSpace 删除字符串前后的 \n \t \s 等内容，argHost 是一个 String 类型的数据
	fmt.Printf("正在分配 %v ...... \n", argHost)

}
