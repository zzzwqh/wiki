package main

import (
	"fmt"
	"path"
)

func main() {
	// path.IsAbs 判断是否时绝对路径，即使在 Linux 系统下，~/ethan.txt 也不算绝对路径
	fmt.Println("======== path.IsAbs ========>")
	fmt.Println(path.IsAbs("~/ethan.txt")) // false

	// path.Dir 获取路径（不带尾部元素）
	fmt.Println("======== path.Dir ========>")
	fmt.Println(path.Dir("https://www.itsky.tech/2022/11/16/Golang-SQL-%E6%93%8D%E4%BD%9C/image-20221120214244373.png"))
	fmt.Println(path.Dir("/usr/local/mysql/my.cnf"))
	fmt.Println(path.Dir("/usr/local/mysql/")) // 打印 /usr/local/mysql/
	fmt.Println(path.Dir("/usr/local/mysql"))  // 这里和上面的结果不一样，打印 /usr/local
	// path.Base 获取尾部元素
	fmt.Println("======== path.Base ========>")
	fmt.Println(path.Base("https://www.itsky.tech/2022/11/16/Golang-SQL-%E6%93%8D%E4%BD%9C/image-20221120214244373.png"))
	fmt.Println(path.Base("/usr/local/mysql/my.cnf"))
	fmt.Println(path.Base("/usr/local/mysql/"))
	fmt.Println(path.Base("/usr/local/mysql"))

	// path.Join 拼接路径
	fmt.Println("======== path.Join ========>")
	fmt.Println(path.Join("/usr/", "local", "mysql", "my.cnf"))

	// path.Ext 返回最后元素的扩展名
	fmt.Println("======== path.Ext ========>")
	fmt.Println(path.Ext("/etc/nginx.d"))

	// path.Split 将文件路径和文件名分开，返回两个 string 变量
	fmt.Println("======== path.Split ========>")
	dirPath, fileName := path.Split("/etc/nginx/conf.d/server.cnf")
	fmt.Println(dirPath, fileName)

	// path.Clean 将复杂的路径，缩减成最简路径
	fmt.Println("======== path.Clean ========>")
	fmt.Println(path.Clean("../usr/../abs/../ethan/server.conf"))

	fmt.Println("======== path.Match ========>")
	fmt.Println(path.Match("abc", "abc"))
	fmt.Println(path.Match("a*", "abc"))
	fmt.Println(path.Match("[/*]*	b", "a/c/d/b"))

}
