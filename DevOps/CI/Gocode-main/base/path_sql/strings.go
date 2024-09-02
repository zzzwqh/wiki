package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	// 常规功能
	fmt.Println("======= 常规方法 =======")

	fmt.Println(strings.Compare("ab", "ab"))
	fmt.Println(strings.Contains("Hello I'm ethan", "a"))
	fmt.Println(strings.Count("Hello I'm ethan", "e"))
	fmt.Println(strings.ToTitle("Title"))
	fmt.Println(strings.ToUpper("Title"))
	fmt.Println(strings.ToLower("Title"))
	fmt.Println(strings.HasPrefix("Hello,Ethan", "He"))
	fmt.Println(strings.HasSuffix("Hello,Ethan", "an"))

	// 字符串切割
	fmt.Println("======= Split 系列方法 =======")
	fmt.Println(strings.Split("/usr/local/bin", "/"))
	fmt.Println(strings.SplitAfter("/usr/local/bin/abs/sdk", "/"))
	fmt.Println(strings.SplitAfterN("/usr/local/bin/abs/sdk", "/", -1))
	fmt.Println(strings.SplitAfterN("/usr/local/bin/abs/sdk", "/", 2))      // 最后一个参数 n 就是返回的切片元素个数
	fmt.Println(len(strings.SplitAfterN("/usr/local/bin/abs/sdk", "/", 3))) // 用 len 可以验证，其实就是将这段 str 分割成多少"段"

	// 获取 substr 索引
	fmt.Println("======= Index 系列方法 =======")
	fmt.Println(strings.Index("ethan,thanks", "than"))     // 返回的是 ethan 中 t 的 index（不论 substr 多长）
	fmt.Println(strings.LastIndex("ethan,thanks", "than")) // 返回的是 thanks 中 t 的 index（不论 substr 多长）

	// 字符串拼接，可以指定分割符
	fmt.Println("======= Join 系列方法 =======")
	fmt.Println(strings.Join([]string{"usr", "local", "bin"}, "/"))

	// 清洗 Trim 系列函数
	fmt.Println("======= Trim 系列方法 =======")
	// 指定需要清洗的字符，Trim 会从前/后一个个遍历，只要在 cutset 里面出现的字符就会被清洗（true），直到 false 出现，也就是说中间的不会被清洗
	fmt.Println(strings.Trim("  Select * from xxx   ", "Se lt"))
	fmt.Println(strings.TrimSpace("  mysql  ")) // cutset 只有空格
	fmt.Println(strings.TrimLeft("!!!Ethan!!!", "!"))
	fmt.Println(strings.TrimRight("!!!Ethan!!!", "!"))
	fmt.Print(strings.TrimFunc("¡¡,¡Hello, Gophers!!!", func(r rune) bool {
		/*
			TrimFunc() 会将 s 的字符的前/后缀传递给 func(r rune) bool 返回的值，决定了 TrimFunc() 是否删除 r 这个字符
			比如传入一个 H，那么 !unicode.IsLetter(H) 返回 false  !unicode.IsNumber(r) 返回 true
			false && ture ==> 返回 false ，那么不删除
			比如传入一个 !，那么 !unicode.IsLetter(!) 返回 true  !unicode.IsNumber(r) 返回 true
			true && true ==> 返回 true ，那么删除
		*/
		// 由如下代码可以观察其执行逻辑
		// fmt.Println("当前的 rune 字符：", string(r))
		// fmt.Println(!unicode.IsLetter(r), "&&", !unicode.IsNumber(r), "=", !unicode.IsLetter(r) && !unicode.IsNumber(r))
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	}))

	// 替换 Replace 函数
	fmt.Println("======= Replace 系列方法 =======")
	// 可以直接在 strings.Replace 函数中指定替换
	fmt.Println(strings.ReplaceAll("/usr/local/mysql/", "/", "\\"))
	fmt.Println(strings.Replace("/usr/local/mysql/", "/", "\\", 2)) // Replace 多了一个传入参数（替换 n 次）
	// 或者创建一个 *strings.Replacer 类型对象（是个结构体），需要传入替换内容
	Replacer := strings.NewReplacer("/", ":")
	splitStr := strings.SplitAfterN("/usr/local/mysql", "/", 2)[1] // 这里我想把根去掉，用 SplitAfterN 分割成 2 个元素，取后面的就行了
	res := Replacer.Replace(splitStr)
	fmt.Println(res)

	// ReaderAt 接口使得可以从指定偏移量处开始读取数据。
	reader := strings.NewReader("Golang 语言是最好的语言")
	p := make([]byte, 6)
	n, err := reader.ReadAt(p, 2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s, %d\n", p, n)
}
