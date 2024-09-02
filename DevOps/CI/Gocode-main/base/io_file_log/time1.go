package main

import (
	"fmt"
	"time"
)

func main() {
	// 取当前的时间，time.Time 是一个结构体类型
	var now time.Time = time.Now()
	fmt.Println(now)
	// 取出年、月、日、时、分、秒
	fmt.Println(time.Now().Year())
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Date())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())

	// 时间戳（1970-现在经过的秒数）1658643828
	fmt.Println(now.Unix())
	fmt.Println(now.UnixMilli()) // 经过的毫秒数 1658643828395
	fmt.Println(now.UnixMicro()) // 经过的微秒数 1658643828395155
	fmt.Println(now.UnixNano())  // 经过的纳秒数 1658643828395155800
}
