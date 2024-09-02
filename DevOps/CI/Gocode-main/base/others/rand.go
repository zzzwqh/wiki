package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(100))

	r := rand.New(rand.NewSource(time.Now().Unix()))
	fmt.Println(r.Int()) // 随机生成 Int 类型值
	fmt.Println(r.Int31())
	fmt.Println(r.Uint32())
	fmt.Println(r.Int63())
	fmt.Println(r.Uint64())
	fmt.Println(r.Float32()) // 0.0 - 1.0 的伪随机 float32 值
	fmt.Println(r.Float64()) // 0.0 - 1.0 的伪随机 float64 值

	// 随机生成 100 以内的 Int 值（利用 Float32）
	fmt.Println(int(r.Float32() * 100))

	// 使用 Intn 生成 n 以内的值
	fmt.Println(r.Intn(100)) // 随机生成 100 以内的 Int 类型值
	fmt.Println(r.Int31n(2000))
	fmt.Println(r.Int63n(200000000000))

	fmt.Println("=====================")
	// 生成定长的随机数
	res := rand.Int31()
	fmt.Printf("%d \n", res%10000)

}
