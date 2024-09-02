package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{Addr: "121.89.244.58:6379", Password: "wqh127.0.0.1", DB: 1})
	res, err := client.Ping().Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
