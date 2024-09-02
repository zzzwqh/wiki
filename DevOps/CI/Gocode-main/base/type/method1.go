package main

import (
	"fmt"
)

// 方法
// 方法其实就是一个函数
// 有属性，有方法 => 对象
// 1. 2. 方法与函数的区别、值类型接收器和引用类型接收器
type Animal struct {
	Species string
	Name    string
	Age     int
}

// 1.1、这个是 Animal 结构体的方法
func (animal Animal) SingASong() {
	if animal.Species == "Dog" {
		fmt.Println(animal.Name, "say: wolf ~")
	} else if animal.Species == "Person" {
		fmt.Println(animal.Name, "say: lala ~")
	}
}

// 1.2、在这里写一个函数
func SingASongFunc(animal Animal) {
	if animal.Species == "Dog" {
		fmt.Println(animal.Name, "say: wolf ~ ===== Func")
	} else if animal.Species == "Person" {
		fmt.Println(animal.Name, "say: lala ~ ===== Func")
	}
}

// 2.1 使用值类型接收器的方法，无法修改结构体的属性值，原因是：和函数传参一样，使用 copy 传递值，不会影响原有变量值
func (animal Animal) ValueReceiveTest() {
	animal.Name = "orange"
}

// 2.2 使用引用类型接收器的方法
func (animal *Animal) IndReceiveTest() {
	(*animal).Name = "noah"
	// animal.Name = "noah" 其实可以直接使用这种方式，也  可以修改成功，go 语言内部实现了此简化写法
}
func (animal *Animal) AddAgeTest() {
	(*animal).Age += 1
}

// 3. 匿名字段的方法
type SonOfDog struct {
	name  string
	age   int8
	color string
}
type Dog struct {
	name     string
	age      int8
	SonOfDog // 结构体嵌套 + 匿名字段
}

// 3.1 匿名字段的方法
func (sonOfDog *SonOfDog) IntroduceMyself() {
	fmt.Println("Hey,my name is", sonOfDog.name)
}

// 3.2 如果 Dog 结构体也有这个方法，那么打印的就是 Dog 结构体实例的名字啦
//func (Dog *Dog) IntroduceMyself() {
//	fmt.Println("Hey,my name is", Dog.name)
//}

func (Dog Dog) ChangeDogName() {
	Dog.name = "valueRecive_yours"
	//fmt.Println(Dog.name)
}

func (Dog *Dog) ChangeDogAge() {
	Dog.age = 100
	//fmt.Println(Dog.age)
}

func main() {
	animalTest01 := Animal{"Dog", "honny", 6}
	animalTest02 := Animal{"Person", "ethan", 21}
	fmt.Println(animalTest01)
	fmt.Println(animalTest02)
	// 1、方法和函数的区别，就是一个调用要传参，一个直接拿到对象调用方法
	SingASongFunc(animalTest01)
	SingASongFunc(animalTest02)
	animalTest01.SingASong()
	animalTest02.SingASong()

	// 2、值类型接收器和指针类型接收器
	animalTest01.ValueReceiveTest()
	fmt.Println(animalTest01) // 值类型接收器方法，名字修改不成功
	animalTest02.IndReceiveTest()
	fmt.Println(animalTest02) // 指针类型接收器方法，名字修改成功了
	// 想要修改原对象的值，就必须要用指针类型接收器

	// 3、匿名字段的方法
	var myDog = Dog{"Ethan", 5, SonOfDog{"noah", 2, "black"}}
	// 我们知道字段可以提升
	fmt.Println(myDog.color)
	fmt.Println(myDog.SonOfDog.color)
	// 其实方法也可以提升，但是如果 Dog 类中有这个方法，那么会重写
	myDog.IntroduceMyself() // 如果将上面 3.2 处的方法解开注释，那么会打印出 Ethan
	myDog.SonOfDog.IntroduceMyself()

	// 4、在方法中使用值接收器，在函数中使用值参数
	var yourDog = Dog{name: "yours", age: 2}
	var othersDog = Dog{name: "others", age: 3}
	yourDog.ChangeDogName() //	更改名字的方法用值接收器，所以无法影响结构体中的名字
	fmt.Println(yourDog.name)
	changeDogNameMethod(othersDog) // 更改名字的函数用值类型传递参数，所以无法影响结构体中的名字
	fmt.Println(othersDog.name)
	fmt.Println("================================================")
	// 5、在方法中使用指针接收器。在函数中使用指针参数
	yourDog.ChangeDogAge() // 更改年龄的方法用指针接收器，可以影响结构体中的年龄
	fmt.Println(yourDog.age)
	changeDogAgeMethod(&othersDog) // 使用指针类型传递了参数，所以可以影响结构体中的名字
	fmt.Println(othersDog.age)

	// 6、那么函数的方式，和结构体使用值/指针接收器的方式有什么不同呢？
	// == important == 结构体的方法， 无论是值、还是指针类型都可以调用方法
	// == important == 函数不行，函数传参时，对于传入的参数有严格的要求，是什么形参，必须传入什么类型的参数
	var herDog = Dog{name: "herDog", age: 10}
	var hisDog = &Dog{name: "hisDog", age: 11}
	herDog.ChangeDogName() // herDog 是值类型，没有修改成功结构体的值，是因为方法的接收器是值类型
	hisDog.ChangeDogName() // hisDog 是指针类型，没有修改成功结构体的值，是因为方法的接收器是值类型
	fmt.Println(herDog.name, hisDog.name)
	herDog.ChangeDogAge() // herDog 是值类型，可以修改原结构体的值，是因为方法的接收器是指针类型
	hisDog.ChangeDogAge() // hisDog 是指针类型，可以修改原结构体的值，是因为方法的接收器是指针类型
	fmt.Println(herDog.age, hisDog.age)
	// ========= important =========
	// 方法优势在于，接收器对于接收到的参数类型更加宽容，接收器的类型决定了是否可以影响原结构体中的值
	// 代码优化：我们定义时尽量按照上面 var hisDog = &Dog{} 的方式，这样调用方法时，传入参数是指针类型的引用参数，占用空间小，不需要另外 COPY 一个值参数
	// ============ end ============

	// 7、在非结构体上使用方法
	// 给 int8 绑定一个自定义 add 方法，每调用一次，就自增 +1
	var num1 uint8 = 10
	var num2 Myunit8 = 20
	num2.addNumMethod()
	fmt.Println(num1, num2)
}

// 7、方法这样写，不能绑定
//func (num int8) addNumMethod() {
//	num = num + 1
//}
// 但是给类型重命名后，是可以绑定的，并且我们可以看到，我传入的 num2 是值参数类型，但接收器是指针类型，也是可以修改原类型变量的
type Myunit8 uint8

func (num *Myunit8) addNumMethod() {
	*(num)++
}

// 4、方法 =》值参数传递
func changeDogNameMethod(dog Dog) {
	dog.name = "ValueReceive_method_others"
	fmt.Println(dog.name)
}

// 5、方法 =》指针参数传递
func changeDogAgeMethod(dog *Dog) {
	dog.age = 20
	fmt.Println(dog.age)
}
