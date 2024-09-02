package main

import (
	"awesomeProject1/entity"
	"fmt"
)

// 比如进入一个后台需要账号，需要获取 Userinfo，先定义，不初始化
var userEz *entity.User

func getUserInfo() *entity.User {
	onceTest.Do(func() {
		// 此刻再定义，放在 onceTest.Do(func()) 中，目的是只让他获取一次，不必多次执行
		userEz = &entity.User{Name: "Ez", Age: 15, Email: "wqh3456@126.com"}
	})
	return userEz
}
func main() {
	res1 := getUserInfo()
	res2 := getUserInfo()
	fmt.Println(res1, res2)
	// 如何判断 res1 和 res2 是同一个 struct 对象？
	res1.Name = "cureForMe"
	// 修改了 res1 对象的值，res2 也跟着改变
	fmt.Println(res1, res2)
	// 可以用 %p 打印出两个对象的地址，其实是指向同一个地址
	fmt.Printf("%p,%p", res1, res2)

}
