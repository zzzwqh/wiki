package main

import (
	"errors"
	"fmt"
)

type fmtTestType struct {
	name  string
	age   int
	hobby []string
}

func main() {
	s1 := fmt.Sprint("枯藤")
	name := "枯藤"
	age := 18
	s2 := fmt.Sprintf("name:%s,age:%d", name, age)
	s3 := fmt.Sprintln("枯藤")
	fmt.Println(s1, s2, s3)

	// 直接打印，没有换行
	fmt.Print("her")
	fmt.Println()
	// 打印并换行
	fmt.Println(1222)
	// 格式化输出
	fmt.Printf("%v\n", 1)
	fmt.Printf("%v\n", "hello")
	fmt.Printf("%v\n", []int{1, 2, 3})
	fmt.Printf("%v\n", map[string]interface{}{"name": "ethan", "age": 18})
	// 打印结构体变量，使用通用占位符 %v
	var obj fmtTestType = fmtTestType{"noah", 19, []string{"play", "code"}}

	fmt.Printf("%v\n", obj)
	// %+v 将 struct 打印，可以看到字段名（ Key:value 的形式 ）
	fmt.Printf("%+v\n", obj)
	// %#v 可以打印出本身的数据类型，并打印出其数据
	fmt.Printf("%#v\n", obj)
	fmt.Printf("%#v\n", []string{"a", "b"})
	// 打印数据类型 %T
	fmt.Printf("%T\n", obj)
	// 当我们想使用 % 这个百分号符号时，怎么办？使用 %% 两个百分号，就可以将后面的 % 转义
	fmt.Printf("100%%\n")

	id, name := 1, "wqh"
	// 通过 fmt.Errorf 方式定义 和 errors.New 方式定义的 Error 是一样的，前者可以传递变量 args
	err1 := fmt.Errorf("id %v ,name %v not found\n", id, name)
	err2 := errors.New("errors.New can't receive any args...\n")
	fmt.Println("=========================")
	fmt.Printf(err1.Error())
	fmt.Printf(err2.Error())
	fmt.Println("=========================")
	fmt.Printf("err1 打印：%v", err1)
	fmt.Printf("err1 类型：%T\n", err1)
	fmt.Printf("err2 打印：%v", err2)
	fmt.Printf("err2 类型：%T\n", err2)

	// 关于 %f 占位符表示的数字，如下示例 %9.2f  9 是数字位数宽度，2是小数后数字位数精度
	n := 88.88
	fmt.Printf("%f\n", n)
	fmt.Printf("%9f\n", n)
	fmt.Printf("%.2f\n", n)
	fmt.Printf("%9.2f\n", n)
	fmt.Printf("%9.f\n", n)

	s := "枯藤"
	s0 := "老树"
	fmt.Printf("%s%s\n", s, s0)
	fmt.Printf("%5s%5s\n", s, s0)
	fmt.Printf("%-5s%-5s\n", s, s0)
	fmt.Printf("%5.7s%5.7s\n", s, s0)
	fmt.Printf("%-5.7s%-5.7s\n", s, s0)
	fmt.Printf("%5.2s%5.2s\n", s, s0)
	fmt.Printf("%05s%05s\n", s, s0)
}
