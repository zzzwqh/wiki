package main

import (
	"errors"
	"fmt"
	"log"
	"reflect"
)

type Student struct {
}

type Option struct {
	Value int
}

func (s *Student) A(o Option) {
	log.Println("A", o.Value)
}

func (s *Student) B(o Option) {
	log.Println("A", o.Value)
}

func (s *Student) C(o Option) (interface{}, error) {
	log.Println("c:", o.Value)
	return "a", errors.New("hello")
}

var methods = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
}

func main() {

	var stu = &Student{}

	reflectType := reflect.TypeOf(stu)
	fmt.Println(reflectType, reflectType.NumMethod())
	for i := 0; i < reflectType.NumMethod(); i++ {
		method := reflectType.Method(i)
		fmt.Println("=============>", method)
		fmt.Println("=============>", method.Name)
		if v, ok := methods[method.Name]; ok {
			var args = make([]reflect.Value, 0)

			o := Option{Value: v}
			// 调用第一个参数会把类本身指针作为第一个参数入栈, 然后再入栈其它参数
			// 因为方法里可能会用到里面的字段 所以需要该结构体的内存首地址
			args = append(args, reflect.ValueOf(stu))
			args = append(args, reflect.Indirect(reflect.ValueOf(o)))
			res := method.Func.Call(args)
			log.Println(method.Name, method.Type.In(0), method.Type.In(1), len(res))
			if len(res) > 0 {
				log.Println(res[0].Type(), res[1].Type())
			}

		}
	}
}
