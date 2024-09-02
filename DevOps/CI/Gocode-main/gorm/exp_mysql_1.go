package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"EthanCode/gorm/model"
)

func main() {

	dsn := "root:wqh127.0.0.1@tcp(42.192.150.241)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}
	// AutoMigrate 作用是如果表不存在则新建，如果表存在，则检查结构是否一致，一致就不涉及表的 DDL ，如果模型结构和表结构不一致 :
	// 1. 模型字段多 => 有新字段就 ALTER TABLE `$TABLE` ADD `$FIELD` 到表上 ，有新增 index 就 CREATE INDEX `$IDX_NAME` ON `$TABLE`(`$FIELD`) 到字段上
	// 2. 模型字段少 => 但是如果原本有 index 索引，但是模型中没有，AutoMigrate ，原本表中有某个字段，模型中没有，AutoMigrate 不会删除该字段
	// 3. 模型字段和表字段有出入 => 即使我使用 tag 修改了模型 Name 字段 column:user_name 他也不会修改原有的表字段 name，只会新增表字段 user_name
	// 【 总结 : 生产禁用 AutoMigrate，瞎加字段或者索引会翻车 】
	err = db.AutoMigrate(&model.Messages{}, &model.User{})
	if err != nil {
		fmt.Println("自动迁移过程报错: ", err.Error())
	}
	err = db.AutoMigrate(&model.InfoUser{})
	if err != nil {
		fmt.Println("xxxx", err.Error())
	}

}
