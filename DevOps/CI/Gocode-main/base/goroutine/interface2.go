package main

import (
	"fmt"
)

// 1、任何类型都属于空接口
type Empty interface {
}

func MyPrint(i Empty) {
	fmt.Println(i)
}

// 上述可以简化，并等同于这样写，匿名空接口
// 匿名空接口，没有名字，只使用一次 ====》 fmt.Println() 就是用了多匿名空接口传参数
func YourPrint(i interface{}) {
	fmt.Println(i)
}

type Animal interface {
	Speak()
	Sleep()
}

type Horse struct {
	Name string
	Age  int8
}

func (horse Horse) Speak() {
	fmt.Println("I'm a horse")
}
func (horse Horse) Sleep() {
	fmt.Println("I'm a horse,and Sleeping ....")
}

// 除了实现 interface Animal 的方法以外，我们还有如下 horse 特有的方法
func (horse Horse) Walk() {
	fmt.Println("I'm a horse,Runing ....")
}

type Snake struct {
	Name string
	Age  int8
}

func (snake Snake) Speak() {
	fmt.Println("I'm a snake")
}
func (snake Snake) Sleep() {
	fmt.Println("I'm a snake，and can’t Sleep.....")
}

// 除了实现 interface Animal 的方法以外，我们还有如下 snake 特有的方法

func (snake Snake) Walk() {
	fmt.Println("I'm a snake,Creeping ....")
}

func main() {
	// 1、任何类型对象都属于空接口（因为空接口类型对象中，没有方法，也意味着，任何类型的对象 都直接实现了空接口类型）
	MyPrint("字符串")
	MyPrint(1)
	MyPrint([4]int{1, 2, 3})
	MyPrint([]int{1, 2, 3})
	// 简化的，使用匿名空接口传入参数的函数 YourPrint
	YourPrint("字符串")
	YourPrint(1)
	YourPrint([4]int{1, 2, 3})
	YourPrint([]int{1, 2, 3})

	// 2、Important 类型断言
	// 当把具体类型，当作接口类型来使用时，具体类型中的属性和非接口类型的方法就无法使用了
	var myhorse Horse = Horse{Name: "little-horse", Age: 2}
	var animalTest01 Animal = myhorse
	animalTest01.Speak()
	animalTest01.Sleep()
	//animalTest01.Walk()	当把具体类型，赋值给接口类型对象，再使用时，非接口类型方法是无法调用的
	//fmt.Println(animalTest01.Name,animalTest01.Age)	当把具体类型，赋值给接口类型对象，再使用时，具体类型中的属性也无法获取
	// 如何能获取具体类型中的属性值，以及方法呢，类型断言（断言正确）
	typeDeclare(animalTest01)
	var mysnake Snake = Snake{Name: "little-snake", Age: 5}
	var animalTest02 Animal = mysnake
	animalTest02.Speak()
	animalTest02.Sleep()
	// 如果 typeDeclare 函数中没有判断断言是否正确（去掉 if），如下这行代码，运行的时候会报错
	// typeDeclare(animalTest02) // panic: interface conversion: main.Animal is main.Snake, not main.Horse

	// 但是加了判断，不会报错
	typeDeclare(animalTest02)
	fmt.Println("============================")
	typeDeclareTest01(1)
	typeDeclareTest01("ethanz")
	typeDeclareTest01(animalTest01)
	typeDeclareTest01(animalTest02)
	fmt.Println("============================")

	typeDeclare(animalTest02)
	fmt.Println("============================")
	typeDeclareTest02(1)
	typeDeclareTest02("ethanz")
	var arrayTest = [4]int{1, 2, 3, 4}
	typeDeclareTest02(arrayTest[:])
	typeDeclareTest02(animalTest01)
	typeDeclareTest02(animalTest02)
	fmt.Println("============================")

	//	3、interface 的底层实现，是引用指针，零值是 nil
	// Horse

}

// 2 、Important 类型断言
func typeDeclare(animal Animal) {
	//	类型断言，经过类型断言以后（如果断言正确），我们就可以获得具体类型的属性字段，以及特有方法（不在接口类型中的方法）
	// horse, isHorse := animal.(Horse)
	// snake, isSnake := animal.(Snake)
	// 上面两行注释，放在 if 语句中，两步并作一步
	if horse, isHorse := animal.(Horse); isHorse {
		fmt.Println(horse.Name, horse.Age)
		horse.Walk()
	} else if snake, isSnake := animal.(Snake); isSnake {
		fmt.Println(snake.Name, snake.Age)
		snake.Walk()
	}
	//	思考：那么不正确的情况情况呢？
	//	上述情况中，如果传入一个具体类型为 Snake 的接口类型 Animal，把 else if 后面删掉（即 typeDeclare 是一个只能断言 Horse 具体类型的函数），那么就会报错如下内容
	// panic: interface conversion: main.Animal is main.Snake, not main.Horse
}

// 利用类型断言，我们可以写这样一个函数，判断传入的参数类型
func typeDeclareTest01(i interface{}) {
	if object, ok := i.(int); ok {
		fmt.Printf("%v is Int type\n", object)
	} else if object, ok := i.(Horse); ok {
		fmt.Printf("%v is Horse struct\n", object)
	} else if object, ok := i.(string); ok {
		fmt.Printf("%v is String type\n", object)
	} else {
		fmt.Println("不知道是什么类型 ~")
	}
}

// 利用类型断言，我们可以写这样一个函数，判断传入的参数类型，用 switch type 判断
func typeDeclareTest02(i interface{}) {
	switch object := i.(type) {
	case int:
		fmt.Println("Int 类型")
	case string:
		fmt.Println("String 类型")
	case []int:
		fmt.Println(object[3])
	// 值得注意的是，case Animal 代码段放在这里，和放在下面，如果传入了接口类型参数，得到的输出结果是不一样的
	// case Animal:
	//	 object.Speak()
	//	 object.Sleep()
	case Horse:
		object.Walk()
	case Snake:
		fmt.Println(object.Name, object.Age)
	case Animal:
		object.Speak()
		object.Sleep()
	default:
		fmt.Println("不知道是什么类型 ~")
	}

}
