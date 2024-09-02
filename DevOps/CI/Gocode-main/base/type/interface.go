package main

import "fmt"

type Asian interface {
	Locate()
	Speak(someWords string)
}

type Chinese struct {
	Name string
	Age  uint8
}

type Korean struct {
	Name string
	Age  uint8
}

func (chinese Chinese) Locate() {
	fmt.Println("Locate at west of Korea")
}
func (chinese Chinese) Speak(someWords string) {
	fmt.Println("Chinese Say: ", someWords)
}

func (korean Korean) Locate() {
	fmt.Println("Locate at east of China")
}
func (korean Korean) Speak(someWords string) {
	fmt.Println("Korean Say: ", someWords)
}

func main() {
	// 定义的结构体，可以赋值给 interace 接口类型，但是没法获取这个结构体的属性
	// 用接口类型，调用接口中的方法，只可以调用接口类型中声明的方法
	var chinese = Chinese{
		Name: "小王",
		Age:  24,
	}
	var asian1 Asian
	asian1 = chinese
	asian1.Locate()
	asian1.Speak("热血沸腾，所向披靡")

	var korean = Korean{
		Name: "axiba",
		Age:  22,
	}
	var asian2 Asian = korean
	asian2.Locate()
	asian2.Speak("阿西吧阿 ")

}
