package entity

import (
	"fmt"
)

type User struct {
	Name  string
	Age   int
	Email string
}

func init() {
	fmt.Println("Import package entity .......")
}

// 当结构体中的字段过多时，我们可以返回一个指针，减少空间开销
func NewUser(name string, age int, email string) *User {
	return &User{
		Name:  name,
		Age:   age,
		Email: email,
	}
}
