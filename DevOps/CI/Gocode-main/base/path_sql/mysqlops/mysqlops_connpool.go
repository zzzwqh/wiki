package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"sync"
)

// 映射了 Mysql 中表字段的结构体
type Person struct {
	userId   int
	username string
	sex      string
	email    string
}

var DB *sql.DB

// 此处的 $PASSWORD 和 $IP 需要自行替换哦
const connectUrl = "root:wqh127.0.0.1@tcp(121.89.244.58:3306)/orders?interpolateParams=true"

// CreateDBConn 建立连接
func CreateDBConn() (err error) {
	// Open 函数只会检查 connectUrl 中的语法格式是否正确，不会真正的连接 Mysql
	DB, err = sql.Open("mysql", connectUrl)
	// 设置最大开启连接数，以及空闲连接数
	DB.SetMaxOpenConns(30)
	DB.SetMaxIdleConns(15)
	if err != nil {
		fmt.Println("sql.Open 出错")
		return
	}
	// 没问题就返回 nil 的 error
	return nil
}

// PingEx 连通性测试
func PingEx() {
	var wg sync.WaitGroup
	// 每次 Ping 都会建立一个 Connection 真正的连接，同时存在的最多连接数量由 DB.SetMaxOpenConns(30) 决定
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(num int) {
			if err := DB.Ping(); err != nil {
				fmt.Println("DB.Ping() Goroutine Seq", num, err)
			} else {
				fmt.Println("DB.Ping() Goroutine Seq", num, "Success")
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}

// QueryRowEx 查询功能测试
func QueryRowEx() {
	fmt.Println("======== QueryRowEx() ========")
	var personIns = Person{}
	// 即使可以有多条记录，QueryRow 返回的只能是一个记录
	err := DB.QueryRow("select * from person;").Scan(&personIns.userId, &personIns.username, &personIns.sex, &personIns.email)
	if err != nil {
		fmt.Println("读取单行时出错", err)
		return
	}
	fmt.Println(personIns)
}

func QueryEx() {
	fmt.Println("======== QueryEx() ========")
	// 查询多行，DB.Query 有两个返回值，返回一个 *Rows 类型的值，可以调用其方法 Scan 赋予结构体
	rows, err := DB.Query("select * from person;")
	if err != nil {
		fmt.Print(err)
	}
	// 如何判断 rows 是否读取完，用 rows.Next()
	for rows.Next() {
		var personInsInner = Person{}
		err := rows.Scan(&personInsInner.userId, &personInsInner.username, &personInsInner.sex, &personInsInner.email)
		if err != nil {
			fmt.Print("读取多行时出错", err)
			return
		}
		fmt.Print(personInsInner, "\n")
	}
}

func InsertEx() {
	fmt.Println("======== InsertEx() ========")
	var personIns4Insert = Person{username: "marry", sex: "female", email: "marry@126.com"}
	// 插入数据需要调用 Exec 方法，返回两个对象，一个 Result 类型对象，一个 error 对象
	res, err := DB.Exec("insert into person(username,sex,email) values(?,?,?);", personIns4Insert.username, personIns4Insert.sex, personIns4Insert.email)
	if err != nil {
		fmt.Println("插入数据报错", err)
		return
	}
	// Result 类型对象是什么？我们来看看
	fmt.Println(res.LastInsertId())
	fmt.Println(res.RowsAffected())
}

func DeleteEx() {
	fmt.Println("======== DeleteEx() ========")
	res, err := DB.Exec("delete from person where user_id>=3;")
	if err != nil {
		fmt.Print("删除数据出错", err)
	}
	fmt.Println(res.RowsAffected())
	// 此时的 res.LastInsertId() 是 0
	fmt.Println(res.LastInsertId())
}

func UpdateEx() {
	fmt.Println("======== UpdateEx() ========")
	sqlStr := "update person set email = ? where user_id = ?"
	res, err := DB.Exec(sqlStr, "may3456@126.com", "2")
	if err != nil {
		fmt.Println("更新数据出错", err)
	}
	fmt.Println(res.RowsAffected())
	fmt.Println(res.LastInsertId())
}

func PrepareEx() {
	fmt.Println("======== PrepareEx() ========")
	var personIns4Pre = Person{}
	sqlStr := "select user_id,username,sex,email from person where user_id > ?"
	stmt, err := DB.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预处理功能错误", err)
		return
	}
	// Prepared statements take up server resources and should be closed after use.
	defer stmt.Close()
	// 一定记得加取地址符号 & ,如下 personIns4Pre 的字段取地址符少一个都会有问题
	stmt.QueryRow(0).Scan(&personIns4Pre.userId, &personIns4Pre.username, &personIns4Pre.sex, &personIns4Pre.email)
	fmt.Println(personIns4Pre)
}

func TransactionEx() {
	fmt.Println("======== TransactionEx() ========")

	var personIns4Tx = Person{username: "Tx", sex: "none", email: "Tx3456@126.com"}
	// 开启事务，返回两个对象，用返回的 *Tx 类型对象代替 sql.*DB 类型对象，执行事务（调用方法 Exec()/Query()/QueryRow() ）
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("事务处理 DB.Begin() 方法调用出错", err)
	}
	res, err := tx.Exec("insert into person(username,sex,email) values(?,?,?)", &personIns4Tx.username, &personIns4Tx.sex, &personIns4Tx.email)
	if err != nil {
		fmt.Println("事务处理 tx.Exec() 出错，执行回滚", err)
		if err := tx.Rollback(); err != nil {
			fmt.Println("执行 tx.Rollback() 回滚出错", err)
			return
		}
		return
	}
	fmt.Println(res.RowsAffected())
	// 提交事务
	err = tx.Commit()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.RowsAffected())

}
func main() {
	// 建立连接
	err := CreateDBConn()
	if err != nil {
		fmt.Println("CreateDBConn Error: ", err)
	}
	// Ping 连通性功能测试
	PingEx()
	// Query 单行查询功能测试
	QueryRowEx()
	// Query 多行查询功能测试
	QueryEx()
	// Insert 插入功能测试
	InsertEx()
	// Delete 删除功能测试
	DeleteEx()
	// Update 更新功能测试
	UpdateEx()
	// Prepare 预处理功能测试
	PrepareEx()
	// TransactionEx 事务处理功能测试
	TransactionEx()

	QueryEx()

}
