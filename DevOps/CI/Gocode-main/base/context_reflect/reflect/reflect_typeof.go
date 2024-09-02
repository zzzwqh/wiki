package main

import (
	"EthanCode/base/goroutine/entity"
	"fmt"
	"reflect"
)

type Fish struct {
	Name string
	Age  int
}

func (fish *Fish) Pop() {

}
func (fish *Fish) Speak() { // 这里用指针接收器

}
func (fish Fish) Swim() { // 这里用值接收器

}

// 如上 Fish 结构体一共三个方法，使用反射时，传入 Ptr 才能得到正确的方法数量，也就是这样写 reflect.ValueOf(&fishIns).NumMethod()，不要用 Elem()

func main() {
	var catIns = entity.Cat{}
	res4Ty := reflect.TypeOf(&catIns)                                 // 这里注意，传入了 Pointer，那么后面和 Field 有关的方法，都要有 Elem() 解开引用！！！
	fmt.Printf("%v %T\n", res4Ty, res4Ty)                             // *entity.Cat *reflect.rtype
	fmt.Printf("%v %T\n", res4Ty.Elem().Name(), res4Ty.Elem().Name()) // Cat string
	fmt.Printf("%v %T\n", res4Ty.Elem().Kind(), res4Ty.Elem().Kind()) // struct reflect.Kind 用 Elem() 解开了引用才是 struct，否则是 Ptr

	fmt.Printf("%v %T\n", res4Ty.Elem().NumField(), res4Ty.Elem().NumField()) // 4 int

	// 1. 通过索引，指定从 struct 中取出字段（可以获取字段名字等...）
	fmt.Println("============ reflect.TypeOf().FieldByName(\"xxx\")获取方法信息 ============> ")
	for i := 0; i < res4Ty.Elem().NumField(); i++ {
		field := res4Ty.Elem().Field(i) // 返回 StructField 类型
		fmt.Println(field)
		fmt.Printf("field.Name: %v , field.Type: %v , field.Index: %v , field.Tag: %v , field.Tag.Get(\"json\"): %v \n", field.Name, field.Type, field.Index, field.Tag, field.Tag.Get("json"))
	}

	// 2. 通过字符串，指定从 struct 中取出字段（缺点: 如果不知道字段名，那么就没法循环取出全部字段信息）
	fmt.Println("============ reflect.TypeOf().Elem().FieldByName(\"xxx\")获取方法信息 ============> ")
	field, _ := res4Ty.Elem().FieldByName("Info") // 返回 StructField 类型
	fmt.Printf("field.Name: %v , field.Type: %v , field.Index: %v , field.Tag: %v , field.Tag.Get(\"json\"): %v \n", field.Name, field.Type, field.Index, field.Tag, field.Tag.Get("json"))

	// 3. 如果 Struct 中嵌套了 Struct ，利用索引可以获取子结构体中的字段信息
	fmt.Println("============ reflect.TypeOf().FieldByIndex([]int)获取嵌套结构体信息 ============> ")
	fmt.Println(res4Ty.Elem().FieldByIndex([]int{3, 1})) // 找到索引为 3 的字段，是 Info ,在从 Info Struct 中找到索引为 1 的字段，是 Food
	fmt.Println(res4Ty.Elem().FieldByIndex([]int{3, 1}).Name)

	// 4. 获取方法的名字，这里不要用 Elem()，直接用 ！！！如果用了 Elem()，获取到的 NumMethod() 有问题
	fmt.Println("============ reflect.TypeOf().Method(i)获取方法信息 ============> ")
	fishIns := &Fish{Name: "ethan", Age: 1}
	// 传入地址！！！如果结构体方法使用了指针接收器，这里没有传入地址，那么将得不到正确的效果！！！如果传入的是值副本，那么只能得到值接收器的方法
	// 传入地址，无论值接收器类型方法，还是指针接收器类型方法，NumMethod() 获取到的方法数量都不会有偏差
	res4FishTy := reflect.TypeOf(fishIns)
	fmt.Println(res4FishTy.NumMethod()) // 先获取了方法的数量
	for i := 0; i < res4FishTy.NumMethod(); i++ {
		fmt.Println("=============>", res4FishTy.Method(i))
		fmt.Println("=============>", res4FishTy.Method(i).Name)
	}
	// 再试试获取 entity.Cat 类型的方法，只能获取到可访问的方法
	fmt.Println(res4Ty.NumMethod()) // 这里 CatIns 的方法只能
	for i := 0; i < res4Ty.NumMethod(); i++ {
		method := res4Ty.Method(i)
		fmt.Println(method)
	}

	// 5. 关于 Elem() 使用细节的验证
	fmt.Println("========== 关于 Elem() 使用细节的验证 ===========")
	// 这里传入的都是实例的地址指针 Ptr
	res4FishTy = reflect.TypeOf(fishIns)
	res4FishVal := reflect.ValueOf(fishIns)
	fmt.Println(res4FishTy.Elem().NumField())   // 字段数量 correct print 2
	fmt.Println(res4FishTy.NumMethod())         // 方法数量 correct print 3
	fmt.Println(res4FishTy.Elem().NumMethod())  // 方法数量 wrong print 1	也就是只拿到了值类型接收器方法 Swim，把指针类型接收器的方法忽略了
	fmt.Println(res4FishVal.Elem().NumField())  // 字段数量 correct print 2
	fmt.Println(res4FishVal.NumMethod())        // 方法数量 correct print 3
	fmt.Println(res4FishVal.Elem().NumMethod()) // 方法数量 wrong print 1	也就是只拿到了值类型接收器方法 Swim，把指针类型接收器的方法忽略了

}

/*
	- 关于 NumMethod()
	1. reflect.Type 类型（reflect.TypeOf()返回值）去调用 ===> reflect.TypeOf(&xxx).NumMethod()
	2. reflect.Value 类型（reflect.ValueOf()返回值）去调用 ===> reflect.ValueOf(&xxx).NumMethod()
	记住都如上写法，不要用 Elem() !!! 输出是符合预期的

	- 关于 NumField()
	1. reflect.Type 类型（reflect.TypeOf()返回值）:
	如果是传入了结构体地址的指针（Ptr），那么要用 Elem() 解引用
	如果是传入了结构体副本，那么不需要用 Elem() 解引用
	2. reflect.Value 类型（reflect.ValueOf()返回值）:
	传入方法地址指针！！！而且不要用 Elem()！！！
	示例:
	var catIns = entity.Cat{}
	res4Val := reflect.ValueOf(&catIns)
	fmt.Println(res4Val.NumMethod())
*/
