package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Egg struct {
	Name string `json:"UserName"`
	Age  int    `json:"UserAge"`
}

func main() {
	file, err := os.OpenFile("./user.json", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	stu := Egg{Name: "Ethan", Age: 16}
	receiveByte, err := json.Marshal(stu)
	if err != nil {
		fmt.Println(err)
	}
	n, err := fmt.Fprintln(file, string(receiveByte))
	if err != nil {
		fmt.Println("zzz", err)
	}
	fmt.Println("输入了多少个字符？", n)

}
