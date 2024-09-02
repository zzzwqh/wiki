package main

import "fmt"

func main() {

	// 指针存储变量内存地址的变量（就是指向了内存地址）
	// 在类型前放 * , 即表示某个类型的指针
	// 在变量前放 & , 表示取这个值的地址
	// 1、指针的定义
	var intTest01 int = 10
	var intInddicator01 *int
	// 使用取地址符号 &
	intInddicator01 = &intTest01
	fmt.Println("intTest01 的值", intTest01)                                //	输出的是值
	fmt.Println("intInddicator01 的值", intInddicator01, "即 intTest01 的地址") // 输出的是这个指针指向的内存地址

	// 2、指针的指针定义
	var inddicatorOfInd01 **int = &intInddicator01
	fmt.Println("inddicatorOfInd01 的值", inddicatorOfInd01, "即 intInddicator01 的地址")
	var inddicatorOfInd02 ***int = &inddicatorOfInd01
	fmt.Println("inddicatorOfInd02 的值", inddicatorOfInd02, "即 inddicatorOfInd01 的地址")

	// 3、把地址反解成值
	fmt.Println("*intInddicator01 反解应该是 intTest 的值", *intInddicator01)
	fmt.Println("*inddicatorOfInd01 反解应该是 intInddicator01 的值", *inddicatorOfInd01)
	fmt.Println("*inddicatorOfInd02 反解应该是 inddicatorOfInd01 的值", *inddicatorOfInd02)
	fmt.Println("**inddicatorOfInd02 反解应该是 intInddicator01 的值", **inddicatorOfInd02)
	fmt.Println("***inddicatorOfInd02 反解应该是 intTest 的值", ***inddicatorOfInd02)

	// 4、指针零值  => nil
	var stringTest01 string = "ethanz"
	var indOfStringTest01 *string
	fmt.Println(indOfStringTest01)
	indOfStringTest01 = &stringTest01
	fmt.Println(*indOfStringTest01)

	// 5、指针是引用类型 可以当作参数传递
	a := 10
	b := &a
	fmt.Println(b)
	c := FanjieInd(b)
	fmt.Println(c)
	d := ChangeValue(b)
	fmt.Println(d, a)

	// 不要向函数传递数组的指针，应该使用切片
	var s1 = "I'm ethan"
	b1 := []byte(s1)
	fmt.Println(b1)

}

// 反解函数
func FanjieInd(a *int) int {
	return *a
}

// 我们在 go 语言中，如果是值类型变量，传入函数中是无法改变值的，只能借用引用型变量
func ChangeValue(a *int) int {
	*a = 8
	return *a
}
