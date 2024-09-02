package main

import (
	"fmt"
	"reflect"
)

// 接口中主要定义行为和方法
type Sender interface { // 定义接口类型
	Send(to string, msg string) error
	SendAll(tos []string, msg string) error
}

type EmailSender struct {
}

func (s EmailSender) Send(to, msg string) error {
	fmt.Println("发送邮件给：", to, "消息内容是：", msg)
	return nil
}

func (s EmailSender) SendAll(tos []string, msg string) error {
	for _, to := range tos {
		s.Send(to, msg)
	}
	return nil
}

func do(send Sender) {
	send.Send("领导", "工作日志")
}

func main() {
	var sender Sender = EmailSender{}     // 将结构体赋值给接口
	fmt.Printf("%T %v\n", sender, sender) // main.EmailSender {}

	sender.Send("kk", "早上好") // 通过接口调用结构体的方法
	do(sender)               // 将接口传入函数作为实参，后续函数调用可以使用接口的方法
	a := sender.(EmailSender)
	fmt.Println(reflect.TypeOf(a)) // 复习下类型断言
	var emailSender EmailSender = EmailSender{}
	do(emailSender) // 当然，直接使用这个结构体也可以调用方法，因为他实现了 Sender 接口
}
