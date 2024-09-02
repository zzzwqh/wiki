package main

import (
	"EthanCode/base/goroutine/entity"
	"fmt"
	"unicode/utf8"
)

type Persion struct {
	Name string
	age  int64
}

// 4. 匿名属性字段的提升。匿名属性方法的提升
type Dog struct {
	Name string
	Age  int64
	Hobby
}
type Hobby struct {
	HobbyName string
	HobbyId   int
}

func (self *Hobby) PrintName() {
	fmt.Println(self.HobbyName)
}

// 取消下面方法注释后，那么就上述的 PrintName 方法就没有机会提升了
// func (self *Dog) PrintName() {
//	 fmt.Println(self.Name)
// }

func main() {
	// 1、自定义类型（type XXX int/string/[]int...） 和 类型别名（rune，byte）的区别
	// byte 通过 %T 拿到的类型，其实还是 unit8（即别名）
	// rune 通过 %T 拿到的类型，其实还是 int32（即别名）
	// 如果我们定义一个如下的 MyInt，拿到的就不是 int 了（type XXX int 的形式，是自定义了一个新类型而不是别名）
	type MyInt int
	var num1 MyInt = 4
	var num2 byte = 6
	var num3 rune = 8
	fmt.Printf("%T,%T,%T", num1, num2, num3) // MyInt 类型的数据类型是 main.int ，byte 类型的数据类型是 unit8，rune 类型的数据类型是 int32
	fmt.Println()

	// 2、如何更改字符串？
	strTest01 := "我是一个执行者"
	runeTest01 := []rune(strTest01)

	for _, value := range runeTest01 {
		if value == '我' {
			value = '你'
		}
		fmt.Printf("%c", value)
	}
	fmt.Println()
	// 上面只是的例子中，并没有修改 rune 类型数据哦，只是根据 key 获取到 value，判断 value 数据
	// 下面我们将 rune 类型数据修改， runTest02[key] = value 这样就可以修改了，并使用 string() 方法重新赋值 strTest02
	var strTest02 = "小王同志"
	runeTest02 := []rune(strTest02)
	for key, value := range runeTest02 {
		fmt.Printf("%c ", value)
		fmt.Println("的下标是", key)
		// 判断字符若为空，则修改成 '对'
		if value != ' ' {
			value = '对'
		}
		runeTest02[key] = value
	}
	strTest02 = string(runeTest02)
	fmt.Println(strTest02)

	// 字符长度获取
	bytesLenOfStrTest02 := len(strTest02)
	fmt.Println(bytesLenOfStrTest02)
	// 字节长度获取
	runesLenOfStrTest02 := utf8.RuneCountInString(strTest02)
	fmt.Println(runesLenOfStrTest02)

	// 3、论如何用函数（用指针传参）修改值
	var numTest01 int64 = 10
	// 传入值是无法变更本体的，只能用指针传入地址
	AddNum(&numTest01)
	fmt.Println(numTest01)

	// 4、匿名字段的提升
	dogTest01 := Dog{Name: "shaoge", Age: 5, Hobby: Hobby{"网球", 1}}
	fmt.Println(dogTest01.HobbyName) // 匿名字段的提升，HobbyName 是 Hooby Struct 的属性，实际上应该像下面这行这么写
	fmt.Println(dogTest01.Hobby.HobbyName)
	dogTest01.PrintName() // 匿名字段方法提升 (如果解开 func (self *Dog) PrintName() 的注释，那么就会打印 Dog 类的 Name

	// 5、引入 entitiy 中的结构体
	softwareUser := entity.NewUser("小王", 16, "wqh3456@126.com")
	// 上面 entity.NewUser 函数获取到的其实是一个指针，我们反解后可以取到结构体
	fmt.Println(softwareUser.Name)
}
func AddNum(i *int64) {
	*i++
}
