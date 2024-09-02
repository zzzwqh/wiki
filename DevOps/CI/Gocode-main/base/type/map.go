package main

import (
	"fmt"
)

func main() {
	// 实际上就是字典类型
	// GoLang 中 key 值的类型必须一致，统一是 string 或者 int 或者 bool
	// GoLang 中 value 值的类型必须一致

	//	1、map 类型数据的定义，如下定义一个 mapTest01 但是并没有初始化，
	var mapTest01 map[string]int // key => string , value => int
	fmt.Println(mapTest01)
	// 此时 map 数据并没有被初始化，能与 nil 作比较，可知是引用类型
	if mapTest01 == nil {
		fmt.Println("mapTest01 is not init yet")
	} else {
		fmt.Println("mapTest01 is already init")
	}
	// 不论是否初始化，当我们去获取 map 数据类型数据中，不存在的一个键值对（key-value）时，都会获得值（value）类型的 0 值
	fmt.Println(mapTest01[""])
	fmt.Println(mapTest01["name"])
	// 没有初始化的 map 类型可以赋值吗？不行
	//mapTest01["name"] = 24	panic: assignment to entry in nil map
	//fmt.Println(mapTest01)

	// 2、map 类型数据的定义，使用 make 函数初始化
	var mapTest02 map[string]int = make(map[string]int)
	fmt.Println(mapTest02)
	// 此时 map 数据已经被初始化
	if mapTest02 == nil {
		fmt.Println("mapTest02 is not init yet")
	} else {
		fmt.Println("mapTest02 is already init")
	}

	//	3、map 数据类型的几种定义方式
	// 完整定义并使用 make 初始化
	var mapTest03 map[string]int = make(map[string]int)
	mapTest03 = map[string]int{"name_03": 12, "age_03": 21}
	fmt.Println(mapTest03)
	// 类型推导
	var mapTest04 = map[string]int{"name_04": 12, "age_04": 21}
	fmt.Println(mapTest04)
	// 简略声明
	mapTest05 := map[string]int{"name_05": 12, "age_05": 21}
	fmt.Println(mapTest05)

	//	4、map 的使用、取值和赋值
	mapTest06 := map[string]int{"name_06": 06, "age_06": 12}
	// 取值直接指明 key
	fmt.Println(mapTest06["name_06"])
	// 赋值直接写
	mapTest06["sex_06"] = 1
	fmt.Println(mapTest06)
	// 取一个不存在的值，会得到 value 值类型的 0 值（即 int 类型的 0 ，以及 string 类型的空字符串 ""）
	fmt.Println(mapTest06["girlfirend"])
	// 删除一个 map 的元素，只能根据 key 来删除
	delete(mapTest06, "age_06")
	fmt.Println(mapTest06)

	//	5、观察第四个知识点，我们可以发现，取一个不存在的 key 会获取 value 的 0 值，那么如果这个 value 本身就是 0 值，怎么判断 map 数据中是不是有这个 key:value 对？
	// 有方法，我们可以使用变量接收 value 值，也可以接收是否有这个变量的 true / flase ，如下所示，将是否有变量的结果，赋值给 ok
	valueOfMapTest, ok := mapTest06["age_06"]
	fmt.Println(valueOfMapTest, ok)
	valueOfMapTest, ok = mapTest06["name_06"]
	fmt.Println(valueOfMapTest, ok)
	// 当我们不想要这个接收 value 值，直接 _
	_, ok = mapTest06["sex_06"]
	fmt.Println(ok)

	//	6、map 数据类型有长度，但是没有容量
	mapTest07 := map[string]int{"name_07": 7, "age_07": 14}
	fmt.Println(len(mapTest07))
	//  ERROR fmt.Println(cap(mapTest07)) 没有容量的属性概念
	// 添加完一个新的元素到 map 中，长度会增加
	mapTest07["salary_07"] = 15000
	fmt.Println(len(mapTest07))

	// 补充 => 数字、字符串、布尔、数组、切片、map类型的 0 值分别是什么？
	// 数字、字符串、布尔、数组 =========> 值类型变量   有自己的 0 值
	// 数字 => 0 、字符串 => ""、布尔 => false、数组 => 元素类型的 0 值
	// 切片、map 类型	========> 引用类型变量   0 值是 nil

	// 7、map 数据之间不可以使用 == 号比较， == 号只能用来判断 map 是否为 nil || 同理， slice 类型也一样哦，因为本质都是 引用类型！ 引用类型！
	// 引用类型不可以用来做 == 比较！只有值类型的数据才可以做 == 比较！
	mapTest08 := map[string]int{"name": 12, "age": 19}
	mapTest09 := map[string]int{"name": 12, "age": 19}
	//fmt.Println(mapTest08==mapTest09)  这样不行
	fmt.Println(mapTest08, mapTest09)
	sliceTest01 := []int{4, 5, 6}
	var sliceTest02 []int
	//fmt.Println(sliceTest01==sliceTest02）这样也不行
	fmt.Println(sliceTest01 == nil)
	fmt.Println(sliceTest02 == nil)
}
