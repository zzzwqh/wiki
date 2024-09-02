package main

import "fmt"

type Pet interface {
	Speak()
	Sleep()
}

// Java 中实现接口，需要显式声明
// Go 和 Python 实现接口，无需显式声明，只要实现了 interface 中所有的方法，就叫做实现接口，这种特性属于 鸭子类型语言特性
type Cat struct {
	Name string
	Age  int8
}
type Pig struct {
	Name string
	Age  int8
}

// 指针类型的 receiver 实现的是 指针类型的 interface
func (cat *Cat) Speak() {
}
func (cat *Cat) Sleep() {
	fmt.Println("喵 ~ ~ ~")
}

// 值类型的 receiver 实现的是 值类型的 interface
func (pig Pig) Speak() {
}
func (pig Pig) Sleep() {
	fmt.Println("哼哼哼 ~")
}

func main() {
	// https://blog.csdn.net/timemachine119/article/details/54927121
	// 指针类型的 receiver 方法实现接口时，只有指针类型的 interface 对象实现了该接口，我们在将 struct 类型对象赋值给 interface 类型的变量时，只能用指针类型的对象赋值
	var myCat Cat = Cat{Name: "yuanxiao", Age: 2}
	var myPet Pet = &myCat // 将 struct 对象赋值给 interface 类型的对象（指针类型）
	fmt.Println(myPet)
	// 值类型的 receiver 方法实现接口，只有值类型的 interface 对象实现了该接口
	var myPig Pig = Pig{Name: "huasheng", Age: 5}
	var yourPet Pet = myPig // 将 struct 对象赋值给 interface 类型的对象（值类型）
	fmt.Println(yourPet)
	// =========== Important ==============
	// 如果实现接口的方法，使用的是指针类型的 receiver（只要其中一个 method 使用的是指针类型的 receiver），则必须把指针赋值给 interface 类型
	// 如果实现接口的方法，使用的是值类型的 receiver，则可以把指针、值赋值给 interface 类型

	// 1、接口的实际用途：忽略实现了接口的 struct 具体的类型，传入函数中去
	GotoSleep(&myCat)
	GotoSleep(myPig)

}

// 1、接口的实际用途：忽略实现了接口的 struct 具体的类型，传入函数中去
func GotoSleep(pet Pet) {
	pet.Sleep()
}
