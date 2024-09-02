package entity

import (
	"fmt"
)

type Animal interface {
	run()
	eat()
}

func init() {
	fmt.Println("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
}
func init() {
	fmt.Println("bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb")
}

// 上述 Animal 是一个接口类型，如下 Dog 和 Cat 都实现了 eat 和 run 两个方法，就都是 Animal 类型
type Dog struct {
	Name string
	Age  string
}

func (self *Dog) run() {
}
func (self *Dog) eat() {
}
func (self *Dog) wolf() {
}

type Cat struct {
	Alias string `json:"alias_name"`
	Age   uint8  `json:"age"`
	Kind  string `json:"kind"`
	Info  `json:"info"`
}
type Info struct {
	Hobby []string `json:"hobby"`
	Food  string   `json:"food"`
}

func (self *Cat) run() {
}
func (self *Cat) eat() {
}
func (self *Cat) Miao() {
}
func (self *Cat) Speak() {
}
