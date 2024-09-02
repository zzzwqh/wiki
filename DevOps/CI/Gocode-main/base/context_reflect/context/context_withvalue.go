package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker4WithValue(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	res, ok := ctx.Value(MachineInfo("SN")).(string)
	if !ok {
		fmt.Println("Invalid SN ...")
	}
LOOP:
	for {
		select {
		case <-ctx.Done(): // timeout 时间到，就会退出循环
			break LOOP
		default:
			fmt.Println("worker4WithValue Worker Get MachineInfo SN :", res)
			time.Sleep(time.Second * 1)
		}
	}
	// 我们再向下传递试试，看看在子协程中，再次定义传递 SN 到孙子协程中，是否会被覆盖
	type PersonInfo string // 为了避免上下文在协程之间传递时，被覆盖的问题，我们用 type 定义新的类型，不同类型的 Key 值无法覆盖
	ctx = context.WithValue(ctx, MachineInfo("SN"), "0000000")
	ctx = context.WithValue(ctx, MachineInfo("factory"), "英特尔")
	ctx = context.WithValue(ctx, PersonInfo("name"), "乔丹")
	wg.Add(1)
	go func(ctx context.Context, wg *sync.WaitGroup) {
		fmt.Println("=========> Grandson routine :")
		fmt.Println(ctx.Value(MachineInfo("SN"))) // 打印 0000000 也就是说会被覆盖
		fmt.Println(ctx.Value(MachineInfo("factory")))
		fmt.Println(ctx.Value(PersonInfo("name")))
		wg.Done()
	}(ctx, wg)
	// 这里不用再加 wg.Wait() 了哈，只需要主协程那里同一等待即可
}

// worker4WithValueBro 和 worker4WithValue 同为主协程的子协程，看看拿 MachineInfo 的值是否有问题（没有）
func worker4WithValueBro(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	res, ok := ctx.Value(MachineInfo("SN")).(string)
	if !ok {
		fmt.Println("Invalid SN ...")
	}
LOOP:
	for {
		select {
		case <-ctx.Done(): // timeout 时间到，就会退出循环，然后执行 wg.Done() 退出子 goroutine
			break LOOP
		default:
			fmt.Println("worker4WithValueBro Worker Get MachineInfo SN :", res)
			time.Sleep(time.Second * 1)
		}
	}

}

type MachineInfo string

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	// 传递一个键值对
	ctx = context.WithValue(ctx, MachineInfo("SN"), "1688888")
	wg.Add(2)
	go worker4WithValue(ctx, &wg)
	go worker4WithValueBro(ctx, &wg)
	wg.Wait()
}
