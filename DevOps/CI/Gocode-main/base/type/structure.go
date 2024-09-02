package main

import (
	"EthanCode/base/type/entity"
	"fmt"
)

// 1、结构体定义
type Person struct {
	name  string
	age   int
	sex   string
	hobby []string
}

func main() {
	// 1、结构体使用
	var personTest01 Person
	fmt.Println(personTest01)
	if personTest01.hobby == nil {
		fmt.Println("personTest01 的 hobby 为空，personTest01.hobby 没有进行初始化")
	}
	// 2.1 第一种，定义并初始化了部分字段，未被定义的引用变量为 nil，其他值类型变量是其对应的 0 值
	personTest02 := Person{name: "ethan", age: 25} //	PersonTest02 := Person{} 不传值也可以
	//personTest02_1 := Person{"ethan_1",25}  错误示范，如果不显式指定字段名，按位置传参定义变量，需要将所有的字段全部按顺序写上
	//personTest02_2 := Person{"ethan_2",25,"male",nil} 正确示范，如果不显式指定字段名，按位置传参定义变量，需要将所有的字段全部按顺序写上
	fmt.Println(personTest02, personTest02.hobby, personTest02.sex)
	if personTest02.hobby == nil {
		fmt.Println("personTest02 的 hobby 为空，personTest02.hobby 没有进行初始化")
	}
	// 因为 hobby 没有初始化，所以下行代码无法使用，切片类型只有初始化后，才能按照指定下标方式修改值
	// personTest02.hobby[0] = "girl"   // ERROR  index out of range [0] with length 0
	// 虽然没有初始化，当我们使用 append 函数后，也会将切片初始化
	personTest02.hobby = append(personTest02.hobby, "girl", "boy", "dog", "duck")
	fmt.Println(personTest02)
	fmt.Println(len(personTest02.hobby))
	fmt.Println(cap(personTest02.hobby))
	personTest02.hobby = append(personTest02.hobby, "you")
	fmt.Println(personTest02)
	fmt.Println(len(personTest02.hobby))
	fmt.Println(cap(personTest02.hobby))

	// 2.2 第二种，全部定义并初始化
	personTest03 := Person{name: "noah", age: 32, sex: "男", hobby: []string{"吃", "喝", "玩", "乐"}} // 直接在
	fmt.Println(personTest03, personTest03.hobby)

	// 2.3 第三种，先将 hobbyTest04 做 make 初始化，hobbyTest04 切片的元素对应了其值类型的 0 值，即空字符串
	hobbyTest04 := make([]string, 3, 4)
	personTest04 := Person{hobby: hobbyTest04}
	fmt.Println(personTest04)
	if personTest04.hobby == nil {
		fmt.Println("personTest04 的 hobby 为空")
	} else {
		fmt.Println("personTest04 的 hobby 不为空")
	}
	personTest04.hobby[0] = "love"
	personTest04.hobby[1] = "girl"
	personTest04.hobby[2] = "love"
	// == TIPS: personTest04.hobby[3] = "you" // index out of range[3] with length 3， hobby 切片长度是 3，切片只能根据切片长度赋值 or 取值，如果长度溢出，要使用 append（和容量无关）
	fmt.Println(personTest04)
	// == TIPS: 如果想继续添加元素到 hobby，使用切片的 append 方法
	personTest04.hobby = append(hobbyTest04, "you", "~")
	fmt.Println(personTest04)
	fmt.Println(len(personTest04.hobby))
	fmt.Println(cap(personTest04.hobby))

	// 4、结构体的零值，是一个值类型，不是引用类型，值类型传入函数，无法被改变
	var personTest05 Person
	// fmt.Println(personTest05==nil) 编译报错，由此可知并非引用类型
	fmt.Println(personTest05)

	// 5、访问结构体字段，如果访问的某个字段是引用类型，必须做了初始化，否则无法访问
	var personTest06 Person = Person{
		name: "ethan",
		age:  12}
	// 切片没有初始化，导致无法使用的情况
	// fmt.Println(personTest06.hobby == nil)
	// personTest06.hobby[0] = "girl"
	// fmt.Println(personTest06.hobby[0])
	// 其他数据类型无论是否初始化，都可以用的情况
	personTest06.name = "wqh"
	personTest06.sex = "male"
	fmt.Println(personTest06)

	// 3、匿名结构体（结构体没有名字！），定义在函数内部，只使用一次，不需要 type 关键字和名字  ==> 实现目的是把数据整合到一起（到一个对象的属性）
	example := struct {
		nobodyName  string
		nobodyAge   int
		nobodyhobby []string
	}{nobodyName: "nobody"} // 要有 {} 在 struct 结尾，代表着初始化（定义即初始化）
	fmt.Println("匿名函数初始化，指针类型不指定就是 nil，其他值类型不指定就是对应的 0 值")
	fmt.Println(example.nobodyhobby == nil)
	fmt.Println(example.nobodyAge)
	example.nobodyAge = 18
	fmt.Println(example)

	// 6、匿名字段，字段没有名字，每种数据类型只能有一个，
	// 比如下面这种情况，如果有匿名字段 string 了，那么就不能再有第二个匿名字段 string，必须指定字段的名字
	type Dog struct {
		string
		int
		sex string
	}
	var bianMu Dog = Dog{"mumu", 5, "girl"}
	var jinMao Dog = Dog{string: "maomao", int: 2, sex: "boy"} // 指定字段赋值，字段类型就是字段名
	fmt.Println(bianMu)
	fmt.Println(jinMao.string)
	fmt.Println(jinMao.sex)
	// 思考 匿名字段有什么用？？

	//	7、结构体嵌套
	type Outline struct {
		color  string
		eyes   string
		height float32
	}
	type Cat struct {
		name    string
		age     uint8
		outline Outline
	}
	var myCat = Cat{"cola", 2, Outline{"white", "brown", 20.2}}
	fmt.Println("我家猫的身高是", myCat.outline.height)
	fmt.Println("我家猫的年龄是", myCat.age)
	// 8、结构体嵌套 + 匿名结构体
	type Wolf struct {
		name       string
		age        int
		appearance struct {
			color string
			eyes  string
		}
	}
	// 匿名结构体，我发现必须要指定字段，否则 appearance 没法在下面一行定义阶段赋值
	var myWolf = Wolf{name: "moon", age: 3}
	fmt.Println(myWolf)
	myWolf.appearance.color = "grey"
	myWolf.appearance.eyes = "blue"
	fmt.Println(myWolf)

	// 9、结构体嵌套 + 匿名字段
	type ParentsOfAnimal struct {
		name   string
		age    int
		color  string
		height int
	}
	type Lion struct {
		name            string
		age             int
		color           string
		ParentsOfAnimal // 匿名字段
	}
	var myLion = Lion{"xinba", 10, "yellow", ParentsOfAnimal{"noah", 2, "yellow", 50}}
	fmt.Println(myLion.ParentsOfAnimal.height)
	fmt.Println(myLion.height)                            // 与上面一行同样，可以输出，这种情况叫做字段提示，本来是 ParentsOfAnimal 的字段，提升到了 Lion 对象上
	fmt.Println(myLion.name, myLion.ParentsOfAnimal.name) // 如果两个结构体都有同一个字段，那么不能提升子结构体的字段
	// 类似于面向对象的继承，可以说是 Lion 结构体继承了 ParentsOfAnimal 结构体的字段 height 属性
	// 而因为 Lion 结构体中有 name 字段，所以就重写了这个 name，输出了自己的 name
	var myLionTest = Lion{ParentsOfAnimal: ParentsOfAnimal{"ethan", 3, "white", 70}}
	fmt.Println(myLionTest.name, myLionTest.age) // 此处仍然会发生”重写“，因为字段被赋予了对应的 0 值

	// 10、结构体的导入，结构体的名字、以及结构体的字段名字的字母大小写，控制着其导出和非导出属性
	// 非导出的结构体、非导出的字段只能在包内使用
	// 我在当前项目路径下，创建了一个 entity 的包，并在里面创建 Martian.go 文件内容如下注释行
	//package entity
	//
	//type Martian struct {
	//	Name     string
	//	Age      int
	//	Account  string
	//	password string
	//}
	var myFriend = entity.Martian{
		Name:    "Alvin",
		Age:     25,
		Account: "alvin3456@126.com",
	}
	// 只能获取大写字母开头的字段，无法获取小写字母开头的字段，并且也无法定义其值 myFriend.password
	fmt.Println(myFriend.Name, myFriend.Age, myFriend.Account)

	// 11、结构体的相等性
	// 结构体的字段全部是值类型，可以比较
	var yourFriend = entity.Martian{
		Name:    "Alvin",
		Age:     25,
		Account: "alvin3456@126.com",
	}
	fmt.Println(myFriend == yourFriend)
	// 结构体的字段中有引用类型，不可以比较
	type Children struct {
		name  string
		age   int
		hobby []string
	}
	var kevin = Children{"kevin", 11, []string{"draw", "music"}}
	var may = Children{"may", 9, []string{"dance", "photograph"}}
	// 结构体字段中有引用类型，不可以比较！
	// fmt.Println(kevin==may)
	// 单拎出来值类型可以比较
	fmt.Println(kevin.name == may.name)

	// Tips: 有个特例，如下的示例中，指针引用类型，是可以比较的
	type Adult struct {
		name *string
		age  int
	}
	var zhangsan = Adult{name: nil, age: 0}
	var lisi = Adult{name: nil, age: 0}
	fmt.Println(zhangsan == lisi)
	name1 := "wangwu"
	name2 := "zhaoliu"
	var wangwu = Adult{&name1, 0}
	var zhaoliu = Adult{&name2, 0}
	fmt.Println(wangwu == zhaoliu)
}
