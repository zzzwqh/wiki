package main

import "fmt"

type AnimalTmper interface {
	Eat()
	Sleep()
}

// Ducker 接口嵌套了 AnimalTmper 接口
type Ducker interface {
	AnimalTmper
	Gaga()
}

type Duck struct {
	Color string
	Age   int8
}

func (duck Duck) Eat() {
	fmt.Println("Eating")
}

func (duck Duck) Sleep() {
	fmt.Println("Sleeping")
}

func (duck Duck) Gaga() {
	fmt.Println("Gagaing")
}

func main() {
	var myduck Duck = Duck{
		Color: "Green",
		Age:   2,
	}
	// 将具体类型的变量赋值给 AnimalTmper 接口类型变量时，只能调用 AnimalTmper 接口中的方法
	var myanimaltmper AnimalTmper = myduck
	myanimaltmper.Eat()
	myanimaltmper.Sleep()
	var myducker Ducker = myduck
	// 将具体类型的变量赋值给 AnimalTmper 接口类型变量时，只能调用 AnimalTmper 接口中的方法
	myducker.Gaga()
	myducker.Eat()
	myducker.Sleep()
	// 调用方法，还是需要类型断言，然后再传入方法
	CallMethod(myanimaltmper.(Duck))
	CallMethod(myducker.(Duck))

}

func CallMethod(duck Duck) {
	duck.Eat()
	duck.Gaga()
	duck.Sleep()
}
