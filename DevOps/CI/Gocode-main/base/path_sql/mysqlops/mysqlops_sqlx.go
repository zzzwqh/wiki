package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"strconv"
	"strings"
	"time"
)

var DBX *sqlx.DB
var err error

type PersonX struct {
	// 这里字段如果小写，DBX.Select() 就查不到了，那可以用 rows.Scan
	Id    string `db:"user_id"`
	Name  string `db:"username"`
	Sex   string
	Email string
}

const connectUrlX = "root:wqh127.0.0.1@tcp(121.89.244.58:3306)/orders?interpolateParams=true"

func SqlxCreateDB() {
	// Open is the same as sql.Open, but returns an *sqlx.DB instead
	DBX, err = sqlx.Open("mysql", connectUrlX)
	if err != nil {
		fmt.Println("Sqlx.Open() Error: ", err)
	}
	DBX.SetMaxOpenConns(100) // 设置连接池最大连接数
	DBX.SetMaxIdleConns(20)  // 设置连接池最大空闲连接数

}

//

// SqlxQueryEx 测试 Sqlx 包多出来的 Select 方法
func SqlxQueryEx() {
	var personList []PersonX
	err := DBX.Select(&personList, "select * from person")
	if err != nil {
		fmt.Println("SqlxQueryEx() Exec Error: ", err)
	}
	fmt.Println(personList)
}

func SqlxNamedExecEx() {
	var personIns4NamedExec = map[string]interface{}{"name": "sqlxNameExec", "sex": "none", "email": "NameExec@163.com"}

	// 这里需要额外注意，我用 &personIns4NameExec 居然报错
	res, err := DBX.NamedExec("insert into person(username,sex,email) values(:name,:sex,:email)", personIns4NamedExec)
	if err != nil {
		fmt.Println("SqlxNameExecEx() Executing Error", err)
	}
	fmt.Println(res.RowsAffected())
}

func SqlxNameQueryEx() {
	/*
		Sqlx 包中的 NameQuery 会返回 *sql.Rows 对象，这个 *sql.Rows 类型对象有更多的方法
		使用 StructScan 方法时传入结构体指针，对结构体赋值
		使用 SliceScan 方法可以得到单行记录 Row 的切片，每个 Row 元素类型也是 []byte 字节切片类型
		使用 MapScan 方法则要传入 map[string]interface{} 类型，会被赋值，赋值后得到的 Key 是 String 类型，Value 也是 []byte 字节切片类型
	*/

	var personIns4NamedQuery = map[string]interface{}{"name": "sqlxNameExec", "sex": "none", "email": "NameExec@163.com"}
	rows, err := DBX.NamedQuery("select * from person where username!=:name", personIns4NamedQuery) // 如果传入 Map 作为条件参数，这里变量名 :name 必须和 Map 类型中的 Key 一致
	//var personIns4NamedQuery = PersonX{Name: "sqlxNameExec"}
	//rows, err := DBX.NamedQuery("select * from person where username!=:username", personIns4NamedQuery)	// 如果传入 Struct 作为条件参数，这里变量名 :username 就要和结构体定义的 Tag 一致

	// *sqlx.rows 对象类似游标，可以关闭，也可以不关闭（没有影响）
	defer rows.Close()
	if err != nil {
		fmt.Println("SqlNameQueryEx() Executing Error", err)
	}
	for rows.Next() {
		// 方法一，StructScan 是 Sqlx 包提供的，传入结构体，会将查到的结果映射赋值给结构体
		var personList4SqlNameQuery PersonX
		rows.StructScan(&personList4SqlNameQuery)
		fmt.Println("使用 StructScan 得到的值：", personList4SqlNameQuery)
		// 方法二，SliceScan 方法会返回一个包含 byte 切片（也就是元素为 []byte）的切片
		personSliceList, _ := rows.SliceScan()
		fmt.Println("使用 SliceScan 得到的值（没有转换 []byte 类型）：", personSliceList) // 打印效果： [[49] [101 116 ... 122] [109 97 108 101] [119 ... 109]]
		for i, col := range personSliceList {
			switch col.(type) {
			case float64:
				personSliceList[i] = strconv.FormatFloat(col.(float64), 'f', 6, 64)
			case int64:
				personSliceList[i] = strconv.FormatInt(col.(int64), 10)
			case bool:
				personSliceList[i] = strconv.FormatBool(col.(bool))
			case []byte:
				personSliceList[i] = string(col.([]byte))
			case string:
				personSliceList[i] = strings.Trim(col.(string), " ")
			case time.Time:
				personSliceList[i] = col.(time.Time).String()
			case nil:
				personSliceList[i] = "NULL"
			default:
				log.Print(col)
			}
		}
		fmt.Println("使用 SliceScan 得到的值（做了类型断言使 []byte 得到转换）：", personSliceList) // 打印效果：[1 ethan male wqh3456@126.com]
		// 方法三，MapScan 方法和 StructScan 一样需要传入对应类型参数
		var personMapList = make(map[string]interface{})
		rows.MapScan(personMapList)
		fmt.Println("使用 MapScan 得到的值（没有转换 []byte 类型）：", personMapList) // 打印效果：map[email:[119 ... 109] sex:[109 97 108 101] user_id:[49] username:[101 116 104 97 110 122]]
		// 这里的 Key 是 string 类型，Value 是 []byte 字节切片类型
		for key, value := range personMapList {
			switch value.(type) {
			case float64:
				personMapList[key] = strconv.FormatFloat(value.(float64), 'f', 6, 64)
			case int64:
				personMapList[key] = strconv.FormatInt(value.(int64), 10)
			case bool:
				personMapList[key] = strconv.FormatBool(value.(bool))
			case []byte:
				personMapList[key] = string(value.([]byte))
			case string:
				personMapList[key] = strings.Trim(value.(string), " ")
			case time.Time:
				personMapList[key] = value.(time.Time).String()
			case nil:
				personMapList[key] = "NULL"
			default:
				log.Print(value)
			}
		}
		// 经过上面 Switch 接口类型的类型断言 + 类型转换，我们 map 类型中的 Value 就被转换成了 String 类型
		fmt.Println("使用 MapScan 得到的值（做了类型断言使 []byte 得到转换）：", personMapList)
	}

}

func SqlxTransactionEx() {
	var personIns4Tx = PersonX{Name: "Sqlx4Tx", Sex: "none", Email: "Sqlx4Tx@126.com"}
	tx, err := DBX.Beginx()
	if err != nil {
		fmt.Println("事务开启失败", err)
	}
	_, err = tx.NamedExec("insert into person(username,sex,email) values(:username,:sex,:email)", &personIns4Tx)
	if err != nil {
		fmt.Println("执行事务失败", err)
		fmt.Println("执行回滚")
		tx.Rollback()
		return
	}
	fmt.Println("执行事务成功")
	tx.Commit()

}
func main() {
	// SqlxCreateDB 创建数据库连接对象
	SqlxCreateDB()
	defer DBX.Close()
	// NameExec 功能测试
	SqlxNamedExecEx()
	// NameQuery 功能测试
	SqlxNameQueryEx()
	// Sqlx 事务操作功能测试
	SqlxTransactionEx()
	// Select 方法功能测试（Get 方法省略）
	SqlxQueryEx()
}
