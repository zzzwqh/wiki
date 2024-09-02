package main

import "fmt"

func main() {
	// 切片本身不拥有任何数据，是对于数组数据的引用
	// 1、切片的定义方法之-------------基于数组定义切片
	arrayTest01 := [10]int{23, 42, 56}
	var sliceTest01 []int // 中括号中不带任何东西，就是切片类型
	// sliceTest01 = arrayTest01[0:len(arrayTest01)-1]
	sliceTest01 = arrayTest01[:] //	把数组从头到尾引用给 sliceTest01
	fmt.Println(sliceTest01)
	fmt.Printf("arrayTest01 数据类型是 %T ， 值是 %v \n", arrayTest01, arrayTest01)
	fmt.Printf("sliceTest01 数据类型是 %T ， 值是 %v \n", sliceTest01, sliceTest01)

	// 2、切片数据类型的使用，与数组数据类型没有什么区别
	// 取值
	fmt.Println(sliceTest01[0])
	// fmt.Println(sliceTest01[3])	// 索引越界，编译不报错，运行时会报错
	// 更改切片数据的值
	sliceTest01[0] = 999
	fmt.Println(sliceTest01)
	// 会影响底层的数组
	fmt.Println(arrayTest01)
	// 更改数组数据的值
	arrayTest01[1] = 888
	// 也会影响切片的值
	fmt.Println(sliceTest01)

	// 3、切片的长度和容量
	// 数组有长度、但长度不能变，切片有长度，但长度可以变
	var arrayTest02 = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(arrayTest02)
	var sliceTest02 []int
	// 前闭后开区间，取数组中前三个
	sliceTest02 = arrayTest02[0:3]
	fmt.Println(sliceTest02)
	// 切片的长度 : 切片有多长
	fmt.Println(len(sliceTest02))
	// 切片的容量 : 切片最大存储多少值，基于底层数组决定的！！但不一定是底层数组的大小！
	fmt.Println(cap(sliceTest02))
	fmt.Println(cap(arrayTest02))

	// 4、长度和容量的研究
	//	切片的容量，是取决于，从什么位置切取的数组！！
	//	截取 [2:5] 那么切片的长度是 3
	sliceTest02 = arrayTest02[2:5]
	fmt.Println(len(sliceTest02))
	//	切片的容量，因为从索引 2 开始截取，arrayTest02[0] 和 arrayTest02[1] 没有截取到
	//	剩下数据数组下标是从 2 - 9 ，也就是 8 个，那么切片的容量就是 8
	fmt.Println(cap(sliceTest02))

	// 5、切片的定义方法之-------------通过 make 来创建切片，make 是一个内置函数
	//  == Tips 前面是通过数组，创建切片 ==
	//	== make 内置函数的参数（参数的类型，切片的长度，切片的容量）==
	var sliceTest03 = make([]int, 3, 4) //	make 创建切片，赋值了切片类型数据，但都是 0 值，可以用作初始化切片数据类型
	fmt.Println(sliceTest03)
	fmt.Println(len(sliceTest03))
	fmt.Println(cap(sliceTest03))

	// 6、切片的定义方法之-------------完整定义声明切片
	var sliceTest04 []int = []int{1, 3, 5, 7}
	fmt.Println(sliceTest04)
	fmt.Println(len(sliceTest04))
	fmt.Println(cap(sliceTest04)) // 因为没有特殊声明，切片容量等于切片长度是 4

	// 7、切片的追加，切片的长度是可以变化的！而数组不可以哦！
	// 定义一个新的切片，来接收追加后的结果（也是切片）
	var sliceTest05 []int = []int{}
	fmt.Println(len(sliceTest05))
	fmt.Println(cap(sliceTest05))
	sliceTest05 = append(sliceTest04, 2, 4, 6, 8)
	fmt.Println(sliceTest05)
	fmt.Println(len(sliceTest05))
	fmt.Println(cap(sliceTest05))
	fmt.Printf("%T", sliceTest05)
	// 当我在原有切片 SliceTest05 基础上再加一个 element 的时候，观察这个切片的容量
	sliceTest05 = append(sliceTest05, 100)
	fmt.Println(sliceTest05)
	fmt.Println(len(sliceTest05))
	// 切片的容量翻倍了！！！
	// == Tips : 只要到达切片容量的临界值，那么新加入元素时，不论多少个，切片容量会翻倍！
	fmt.Println(cap(sliceTest05))
	fmt.Printf("%T \n", sliceTest05)

	// 8、一个有意思的 “寄生” 现象，如果基于数组定义切片，切片容量（取决于数组长度）不够用时，底层会新创建一个数组（长度为原来的 2 倍）重新 ”寄生“
	// 如何验证这个现象呢？如下定义一个数组，并基于数组定义一个切片
	var arrayTest06 [5]int = [5]int{1, 2, 3, 4, 5}
	var sliceTest06 []int = arrayTest06[2:4]
	fmt.Println("基于数组定义拿到的切片长度", len(sliceTest06)) // 切片长度
	fmt.Println("基于数组定义拿到的切片容量", cap(sliceTest06)) // 切片容量（取决于定义切片数据时，切割数组时下标的位置，以及数组长度）
	fmt.Println("基于数组定义拿到的切片是", sliceTest06)
	sliceTest06[1] = 5
	sliceTest06 = append(sliceTest06, 6)
	fmt.Println("更改切片中的元素后，切片数据为", sliceTest06)
	fmt.Println("数组也会随之改变，数组数据为", arrayTest06)
	// 此时切片容量已经饱和，再加入元素会怎么样？
	sliceTest06 = append(sliceTest06, 7)
	fmt.Println("撑爆切片容量后，现在切片长度", len(sliceTest06)) // 切片长度是 4
	fmt.Println("撑爆切片容量后，现在切片容量", cap(sliceTest06)) // 切片容量变成了原来的 2 倍！！！不是 4
	fmt.Println("基于数组定义拿到的切片是", sliceTest06)
	fmt.Println("原生的数组不会继续拿到切片的值", arrayTest06)
	// 此时再次修改切片前端索引的值，原生数组（宿主）也不会受影响，这个切片已经转移寄生了 ~ ~
	sliceTest06[1] = 8
	fmt.Println(arrayTest06)
	fmt.Println(sliceTest06)
	// 切片就是个寄生虫，如果找到了新的宿主，就和原来的宿主没有关系了

	// 9、切片的函数传递
	// 数组是值类型、而切片是引用类型
	// go 语言中的参数传递是 copy 传递，把变量赋值一份出入函数中
	// 如果是值类型，那么复制一个新值，传入
	// 如果是引用类型，那么复制一个引用，传入（引用指针的指向没有变）
	// 由此导致一个现象，如果传入一个函数的参数是值类型，传出后不会影响之前的变量值
	// 如果传入一个函数的参数是引用类型，传出后会影响之前的变量值
	var arrayTest07 [3]int = [3]int{1, 2, 3}
	var sliceTest07 []int = arrayTest07[:]
	fmt.Println("arrayTest07 初始化值为", arrayTest07)
	fmt.Println("sliceTest07 初始化值为", sliceTest07)
	ReceiveArray(arrayTest07)                          // 函数体内部是 [11111 2 3]
	fmt.Println("arrayTest07 经过函数调用后，值为", arrayTest07) // 出了函数值不变还是 [1 2 3]
	ReceiveSlice(sliceTest07)                          // 函数体内部是 [22222 2 3]
	fmt.Println("sliceTest07 经过函数调用后，值为", sliceTest07) // 出了函数也是 [22222 2 3]
	// == Tips 结论 数组作为参数传入函数，在函数内部作了修改，再传出时不会影响传入的数据本身 ==
	// == Tips 结论 切片作为参数传入函数，在函数内部作了修改，再传出时影响指针指向的数据本身 ==

	//10、多维切片
	// 两种方式，第一种完整定义方式，第二种通过 make 内置函数声明多维切片
	// var sliceTest08 [][]int = [][]int{{1,2,3},{2,2,},{3,3,3}}
	var sliceTest08 [][]int = make([][]int, 3, 4) // 内层的切片没有初始化
	fmt.Println(sliceTest08[0])                   // 是个切片 没有被初始化
	//	sliceTest08[0][1] = 1 因为内层切片没有被初始化，如此写代码运行时会报错空指针异常
	// 只要初始化内层切片就可以了
	sliceTest08[0] = make([]int, 3, 4)
	sliceTest08[0][1] = 1
	fmt.Println(sliceTest08)

	// 11、代码优化，将基于大数组”寄生“的切片，使用 copy 函数，将其复制到另一个切片上（这个切片“寄生”在长度较小的数组）
	// 可以减少内存使用，注意，要保证接收方的切片容量够大，否则不会像 append(s,elements) 方法执行完一样扩大切片容量，只会接收有限个 elements
	var sliceTest09 []int = make([]int, 2, 2)
	fmt.Println("sliceTest05 的值是", sliceTest05)
	copy(sliceTest09, sliceTest05)
	fmt.Println("sliceTest09 的容量只有 2，使用 copy 函数方法只能接收 2 个值", sliceTest09)
	//
	sliceTest09 = append(sliceTest09, 2)
	fmt.Println(sliceTest09)
	fmt.Println(len(sliceTest09))
	fmt.Println(cap(sliceTest09))
}

func ReceiveArray(a [3]int) {
	a[0] = 11111
	fmt.Println("当前函数传入的是数组（值类型），修改后在函数内部输出得", a)
}

func ReceiveSlice(a []int) {
	a[0] = 22222
	fmt.Println("当前函数传入的是切片（引用类型），修改后在函数内部输出得", a)
}
