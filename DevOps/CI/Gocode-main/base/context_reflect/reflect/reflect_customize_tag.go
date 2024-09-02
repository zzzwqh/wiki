package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

type ATest struct {
	AAA string        `json:"AAA" def:"123456"`
	BBB time.Duration `json:"BBB" def:"8000000000"`
	CCC bool          `json:"CCC"  def:"true"`
	DDD uint64        `json:"DDD" def:"3000"`
	EEE int64         `json:"EEE" def:"-6000"`
}

func setValue(v reflect.Value, defVal string) {
	// 通过 reflect.Value.Kind() 方法可以获得具体种类，然后用 reflect.Value.Set*** 系列方法改变值
	switch v.Kind() {
	case reflect.String:
		v.SetString(defVal)
		break
	case reflect.Bool:
		tmp, err := strconv.ParseBool(defVal)
		if err != nil {
			panic(err.Error())
		}
		v.SetBool(tmp)
		break
	case reflect.Int64, reflect.Int32, reflect.Int, reflect.Int8:
		tmp, err := strconv.ParseInt(defVal, 10, 64)
		if err != nil {
			panic(err.Error())
		}
		v.SetInt(tmp)
		break
	case reflect.Uint64, reflect.Uint32, reflect.Uint, reflect.Uint8:
		tmp, err := strconv.ParseUint(defVal, 10, 64)
		if err != nil {
			panic(err.Error())
		}
		v.SetUint(tmp)
		break
	default:
		panic("unsupported type :" + v.Type().Kind().String())
	}
}

func parseStruct(v reflect.Value, s reflect.Type) {
	// 根据传入的结构体字段遍历
	for idx := 0; idx < s.NumField(); idx++ {
		// 获取当前结构体字段
		st := s.Field(idx)
		// fmt.Println(st) // 打印 {AAA  string json:"AAA" def:"123456" 0 [0] false}  .....，，也就是结构体信息

		// reflect.TypeOf().Field(i) 可以根据 tag 中的 def 获取值
		defTag := st.Tag.Get("def")
		// fmt.Println(defTag)	 // 打印 123456 8000000000 true 3000 .... ，也就是 def 的值信息

		// 如果传入的结构体字段，也是个结构体，需要再次调用本函数，即递归
		if st.Type.Kind() == reflect.Struct {
			parseStruct(v.Elem().FieldByName(st.Name), st.Type)
		} else {
			// 如果 def 的值不是空，则赋值，否则不赋值（即为 0 值或者 nil）
			if defTag != "" {
				// 如果 Kind() 是指针，那么要用 Elem()，其实直接都用了 v.Elem().FieldByName() ，这个方法返回的是 relect.Value
				if v.Type().Kind() == reflect.Ptr {
					fmt.Println("v.Elem().FieldByName(st.Name) ===>", v.Elem().FieldByName(st.Name))
					setValue(v.Elem().FieldByName(st.Name), defTag)
				} else {
					fmt.Println("v.FieldByName(st.Name) ===>", v.FieldByName(st.Name))
					setValue(v.FieldByName(st.Name), defTag)
				}
			}
		}
	}
}

func main() {
	aa := ATest{}
	info, _ := json.MarshalIndent(aa, "", "\t") // 和 marshal 函数一样，就是输出格式不一样
	fmt.Println(string(info))
	t := reflect.TypeOf(aa)
	//这里必须指针，为了成功修改原结构体，不加也会 panic
	v := reflect.ValueOf(&aa)
	// 解析结构体
	parseStruct(v, t)
	info, _ = json.MarshalIndent(aa, "", "\t")
	fmt.Println(string(info))
}
