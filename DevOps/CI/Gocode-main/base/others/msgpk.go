package main

import (
	"fmt"
	"github.com/vmihailenco/msgpack"
	"io/ioutil"
)

type Face struct {
	Id   string
	Name string
	Type string
}

func main() {
	// messagepack 序列化
	var faceIns01 = Face{Id: "1", Name: "ethan", Type: "Person"}
	output, err := msgpack.Marshal(&faceIns01)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(output))
	err = ioutil.WriteFile("msg.txt", output, 0666)
	if err != nil {
		fmt.Println(err)
	}
	// messagepack 反序列化
	var faceInsRec Face
	recStr, _ := ioutil.ReadFile("msg.txt")
	err = msgpack.Unmarshal(recStr, &faceInsRec)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(faceInsRec)
}
