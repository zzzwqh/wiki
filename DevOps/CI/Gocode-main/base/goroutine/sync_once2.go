package main

import (
	"fmt"
	"sync"
)

// 小美去淘宝购物，只加载一次这个账号
type Woman struct {
	name string
}

var (
	once         sync.Once
	shoppingUser *Woman
)

func actionOfUser() *Woman {
	// once.Do 包裹的代码，无论调用多少次，只执行一次
	once.Do(func() {
		shoppingUser = new(Woman)
		shoppingUser.name = "小美"
		fmt.Println("Load User 小美.....")
	})
	// 这行代码会一直执行
	fmt.Println("小美 Shopping")
	return shoppingUser
}
func main() {
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			actionOfUser()
			wg.Done()
		}()
	}
	wg.Wait()
}
