package main

import (
	"fmt"
	"reflect"
)

type Order struct {
	orderId string
	price   uint8
}

func (order *Order) PriceChange(price uint8) {
	order.price = price

}
func (order *Order) PrintPrice() {
	fmt.Println("Order Price is", order.price)
}
func (order *Order) GetPrice() (price uint8) {
	return order.price
}
func main() {
	var orderIns = &Order{orderId: "1688", price: 68}
	res4OrderTy := reflect.TypeOf(orderIns)
	res4OrderVal := reflect.ValueOf(orderIns) // 这里注意，传入了 Pointer，那么后面和 Field 有关的方法，都要有 Elem() 解开引用！！！
	fmt.Println(res4OrderVal.NumMethod())
	for i := 0; i < res4OrderVal.NumMethod(); i++ {
		methodByTy := res4OrderTy.Method(i)
		methodByVal := res4OrderVal.Method(i)
		fmt.Println("通过 reflect.TypeOf().Method(i) 返回 Method 类型名字/接收器等信息：", methodByTy)
		fmt.Println("通过 reflect.ValueOf().Method(i) 返回地址：", methodByVal)
	}

	//  通过方法名字，获取到方法 ==> 调用方法
	fn4PriceChange := res4OrderVal.MethodByName("PriceChange")
	var args4PriceChange = []reflect.Value{reflect.ValueOf(uint8(28))}
	fn4PriceChange.Call(args4PriceChange) // 调用方法，需要传入参数
	fn4PrintPrice := res4OrderVal.MethodByName("PrintPrice")
	var args4PrintPrice = []reflect.Value{}
	fn4PrintPrice.Call(args4PrintPrice) // 如果不需要传入参数，传入空参数即可
	fn4GetPrice := reflect.ValueOf(orderIns).MethodByName("GetPrice")
	var args4GetPrice = []reflect.Value{}
	res := fn4GetPrice.Call(args4GetPrice)              // 返回的是什么对象？  []reflect.Value{}
	fmt.Printf("%v %T %v %T", res[0], res[0], res, res) // 返回的也是个 Slice，即一堆返回值的话，可以按索引获取
}
