package main

import (
	"fmt"
	"time"
)

func main() {
	// 将时间戳转换成时间
	timeObj := time.Unix(1658643828, 0)              // 按照秒级的时间戳
	timeObjMilli := time.UnixMilli(1658643828395)    // 按照毫秒级的时间戳
	timeObjMicro := time.UnixMicro(1658643828395155) // 按照微秒级的时间戳
	fmt.Println(timeObj)
	fmt.Println(timeObjMilli)
	fmt.Println(timeObjMicro)
	fmt.Println(timeObj.YearDay())
	fmt.Println(timeObj.Year())
	fmt.Println(timeObj.Day())
	fmt.Println(timeObj.Date())

	// 时间间隔 time.Second 是 time.Duration 类型，其实底层是 int64 类型，但不同于 int64 类型

	// 时间运算 ==》 time.Time 对象的方法
	var now time.Time = time.Now()
	// 1、time.Time 对象的 Add 方法，时间相加
	lastTime := now.Add(1 * time.Hour)
	fmt.Println(lastTime)
	// 2、time.Time 对象的 Sub 方法，时间相减（两个时间对象）
	fmt.Println(lastTime.Sub(now)) // 1h0m0s

	// 3、判断两个时间是否相等 Equal 方法
	fmt.Println(lastTime.Equal(lastTime))
	fmt.Println(lastTime.Equal(now))
	// 4、判断两个时间先后的 Before、After 方法
	fmt.Println(lastTime.After(now))
	fmt.Println(lastTime.Before(now))
	fmt.Println(lastTime.Before(lastTime))
	fmt.Println(lastTime.After(lastTime))
	fmt.Println(lastTime.Equal(lastTime))
	// 5、时间格式化 ，其他语言用 %Y-%M-%D 但是 Go 不同，是用 Golang 的诞生时间做占位符
	strNowVer1 := now.Format("上海时间 2006-01-02 15:04:05")
	fmt.Println(strNowVer1)
	strNowVer2 := now.Format("15:04:05")
	fmt.Println(strNowVer2)

	// 6、把字符串格式的时间，转换成 time.Time 事件对象
	strTimeParse, err := time.Parse("2006年01月02日，15时04分05秒", "2022年07月25日，12时05分45秒")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(strTimeParse)

	// 7、指定时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(loc)
	// 按照指定时区和指定格式解析字符串时间，如果不加 Location 是 2019-08-04 14:15:20 +0000 UTC
	timeObj, err = time.ParseInLocation("2006/01/02 15:04:05.0000", "2022/08/04 14:15:20.8000", loc)
	// 加了 Localtion 后，上面的打印结果是 2022-08-04 14:15:20 +0800 CST
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Sub(now))

	// 8、时区的概念
	timeInAmerica := time.Date(2020, 12, 25, 12, 11, 42, 12, time.UTC)
	timeInChina := time.Date(2020, 12, 25, 20, 11, 42, 12, loc)
	// 上面两个时间实际上是相等的，一个是 UTC 时区，一个是 CST 时区，虽然差 8 小时，但是是同一个实际
	fmt.Println(timeInAmerica)
	fmt.Println(timeInChina)
	fmt.Println("判断是否相等：", timeInChina.Equal(timeInAmerica))

}
