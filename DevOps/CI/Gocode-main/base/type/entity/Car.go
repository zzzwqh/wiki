package entity

import "fmt"

type Car struct {
	Name         string
	Introduction string
	performance  map[string]int
	Materials    []string
}

func init() {
	fmt.Println("init() 是这个包内的初始化方法，只要引入当前包，我就会被执行")
}
func init() {
	fmt.Println("init() 可以被写入当前包下的一个或多个 go 文件中，只要引入当前包，就会依次执行")

}
func NewCar(name string, introduction string, performance map[string]int, materials []string) (car Car) {
	car = Car{
		Name:         name,
		Introduction: introduction,
		performance:  performance,
		Materials:    materials,
	}
	return car
}

func (car *Car) AddNPerfomance(performance map[string]int) {
	// 给 map[string]int 类型数据赋值前，检查是否为 nil，若为空需要做初始化，否则 panic: assignment to entry in nil map
	if car.performance == nil {
		car.performance = make(map[string]int)
	}
	for key, value := range performance {
		car.performance[key] = value
	}
}

func (car *Car) ADDMaterials(material string) {
	// 按照正常的思路，给 []string 类型数据赋值前，也要检查是否为 nil，但是 append 方法也有初始化的效果
	car.Materials = append(car.Materials, material)

	// 或者这样写，逻辑更加清晰，检查是否为空，为空则初始化的代码
	//if car.Materials == nil {
	//	car.Materials = make([]string,1,1)
	//}

	// 或者也可以这样写
	//if car.Materials == nil {
	//	car.Materials = []string{material}
	//} else {
	//	car.Materials = append(car.Materials, material)
	//}

}
