package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

type CarsList struct {
	Cars []Car
}

type Car struct {
	Id   int
	Name string
}

// gob 序列化
func se() {
	file, err := os.OpenFile("./gob.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		closeErr := file.Close()
		if err == nil {
			err = closeErr
		}
	}()
	var carIns = Car{Id: 3, Name: "kkk"}
	encoder := gob.NewEncoder(file)
	err = encoder.Encode(carIns)
	if err != nil {
		fmt.Println(err)
	}
}

// gob 反序列化
func xse() {
	file, err := os.OpenFile("./gob.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var carListIns4Rec = new(CarsList)
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&carListIns4Rec)
	if err != nil {
		fmt.Println("反序列化 ERROR：", err)
	}
	fmt.Println(carListIns4Rec)
}

func main() {
	se()
	xse()
}
