package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	// 加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(loc)
	// 按照指定时区和指定格式解析字符串时间，如果不加 Location 是 2019-08-04 14:15:20 +0000 UTC
	timeObj, err := time.ParseInLocation("2006/01/02 15:04:05", "2019/08/04 14:15:20", loc)
	// 加了 Localtion 后，上面的打印结果是 2019-08-04 14:15:20 +0800 CST
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Sub(now))
}
