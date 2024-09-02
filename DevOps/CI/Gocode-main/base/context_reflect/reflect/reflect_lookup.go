// Golang program to illustrate
// reflect.Tag.Lookup() Function

package main

import (
	"fmt"
	"reflect"
)

// Main function
func main() {
	type S struct {
		F0 string `val:"123456"`
		F1 string `val:""`
		F2 string
	}

	s := S{}
	st := reflect.TypeOf(s)
	fmt.Println("============= 下面使用 reflect.Type().Field().Tag.Lookup() 方法 ===============")

	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)

		// use of Lookup method
		// the ok return value reports whether the value was explicitly set in the tag string
		//ok 用来获知该 tag 是否在结构体字段 Tag 中有定义，比如 F2 字段中就会返回 false，打印 None specific
		if value, ok := field.Tag.Lookup("val"); ok {
			if value == "" {
				fmt.Println("(Empty)")
			} else {
				fmt.Println(value)
			}
		} else {
			fmt.Println("(None specific)")
		}

	}
	fmt.Println("============= 代码都一样，下面使用 reflect.Type().Field().Tag.Get() 方法 ===============")

	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)

		value := field.Tag.Get("val")
		if value == "" {
			fmt.Println("(Empty)")
		} else {
			fmt.Println(value)
		}
	}
}
