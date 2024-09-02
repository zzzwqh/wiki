package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:123456@tcp(192.168.17.106)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
