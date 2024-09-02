package main

import (
	"awesomeProject/entity"
	"fmt"
)

// 我们在项目路径下，新建一个 entity 的包，entity 包下有一个 Car.go 文件

func main() {
	//var myCar entity.Car
	//fmt.Println(myCar)	只是声明了变量，什么也没干。但是 init() 方法是在引入 package 的时候就依次调用的

	// 通常情况，我们会在结构体文件中，构造一个 New 方法，用来规约赋值
	var myCar entity.Car = entity.NewCar("BM", "fast", map[string]int{"appearance": 90}, []string{"AL alloy"})
	fmt.Println(myCar.Name, myCar.Introduction)
	// 虽然我们不能直接访问 performance 属性，但是依然可以通过方法赋值，道理和上面的 NewCar() 一样
	myCar.AddNPerfomance(map[string]int{"engine": 91, "price": 92})
	// 可以看到新加入的 map[string]int 类型参数
	fmt.Println(myCar)

	var yourCar entity.Car = entity.Car{Name: "BC"}
	yourCar.ADDMaterials("glass")
	yourCar.AddNPerfomance(map[string]int{"a": 1})
	fmt.Println(yourCar)
}
