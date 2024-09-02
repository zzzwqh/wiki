---
title: Review-G
date: 2022-08-19 15:38:54
tags: Review
categories: Review
---

> Golang Review

<!--MORE-->

## fmt.Scan && fmt.Scanln && fmt.Scanf

> Scan 和 Scanln 的区别 ===> https://blog.csdn.net/u013792921/article/details/84553192 
>
> 也就是说，Scan 会死等你输入变量的值，如果不输入就一直阻塞，Scanln 不会阻塞

```go
func main() {
	var name string
	var age int
	var married bool
	// 1.1 Scanln 接受控制台输入，换行符就结束，即使所有变量并没有全部 Scan 到值的传入
	//fmt.Scanln(&name, &age, &married)
	// 1.2 Scan 接受控制台输入，即使换行，也会等待所有参数输入才结束
	//fmt.Scan(&name, &age, &married)
	// 1.3 Scan 是根据扫描到的输入结果，自动识别格式，然后把值传给变量
	fmt.Scanf("1:%v\n2:%v\n3:%v", &name, &age, &married)
	// 2. 利用 Fprintln 输出 inputString 内容 到 os.Stdout （输出到控制台）
	fmt.Fprintln(os.Stdout, "个人信息...", name, age, married)
	// 3. Sprintf 拼接字符串、可返回值
	var res01 string = fmt.Sprintf("输入的 name 是: %v \n 输入的 age 是: %v \n 是否结婚: %v \n 来一个随机数: %v", name, age, married, rand.Intn(100))
	fmt.Println(res01)
}
```

> 如下图是 Scanf 方法的测试验证

![image-20220819154221973](image-20220819154221973.png)

## bufio.NewReader

> 想从控制台读取输入，就要用 os.Stdin，想从文件种读取内容，就用 file 类型的对象（他们都是实现了 io.Reader 的接口方法）

```go
func main() {
	// 4. bufio 从控制台读取输入，就要用 os.Stdin
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Input Your Name Please: ")
	str, err := reader.ReadString('\n')
	//str, _, err := reader.ReadLine()
	if err == io.EOF {
		fmt.Println(err)
	}
	fmt.Println(string(str))
}
```

## error.New && fmt.Errorf && 自定义错误对象

> error.New && fmt.Errorf 都会返回一个错误类型对象，区别就是，fmt.Errorf 可以传递变量

```go
// error 是一个接口类型的变量
type error interface {
	Error() string
}
// 看 error.New 方法，返回了 errorString 结构体对象，这个 errorString 又是什么？
func New(text string) error {
	return &errorString{text}	
    // return &errorString{s:text} 这样看清晰一点，将 text 赋予 s 字段 
}
// 看 errorString 结构体，只有一个 s string 字段
type errorString struct {
	s string
}
// 实现了 error 接口类型的 Error 方法（所有方法）
func (e *errorString) Error() string {
	return e.s
}
```

> 所以我们实际上可以自定义错误，只要实现了 error 接口类型中的所有方法

```go
// RadiusErr 自定义的错误类型，实现了 error 接口的结构体
type RadiusErr struct {
	Msg    string
	Radius float64
}

// 实现了 error 接口的 Error 方法
func (radi RadiusErr) Error() string {
	return fmt.Sprintf("半径是 %v，%v", radi.Radius, radi.Msg)
}

// 获取圆面积的函数，可以返回 result 和 error
func getCircleArea(radius float64) (area float64, e error) {
	if radius <= 0 {
		return 0, RadiusErr{Radius: radius, Msg: "半径赋值不符合要求"}
	}
	return radius * radius * 3.14, nil
}

func main() {
	// 半径，从键盘获取输入
	var inputRadius float64
	fmt.Scan(&inputRadius)
	res, err := getCircleArea(inputRadius)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("计算得到的面积为", res)
	}
}
```

## time.Now().Format && time.Parse && && time.Date&& time.Loadlocation

```go
func main() {
	// time.Now().Format
	now := time.Now()
	nowFormat := now.Format("上海时间 2006-01-02 15:04:05 -0700 MST")
	fmt.Println(nowFormat)
	// 本地时区
	datetime01 := time.Date(2022, 12, 24, 17, 12, 30, 0, time.Local)
	loc01 := datetime01.Location()
	fmt.Println(loc01)
	fmt.Println(datetime01)
	// 指定时区
	locSet, _ := time.LoadLocation("America/Los_Angeles")
	datetime02 := time.Date(2022, 12, 24, 1, 12, 30, 0, locSet)
	loc02 := datetime02.Location()
	fmt.Println(loc02)
	fmt.Println(datetime02)
	fmt.Println(datetime01.Equal(datetime02))
	// time.Parse
	timeObj, _ := time.Parse("上海时间 2006-01-02 15:04:05 -0700 MST", "上海时间 2020-08-24 10:09:47 +0800 CST")
	fmt.Printf("类型:  %T \n值:    %v\n", timeObj, timeObj)
	// time.Now().unix
	fmt.Println(time.Now().Unix())
}
```

## strconv.Itoa && strconv Atoi && strconv.Parse\* && strconv.Format\*

```go
func main() {
	// 1. strconv.AtoI
	ins01, _ := strconv.Atoi("65")
	fmt.Printf("%T  %v \n", ins01, ins01)

	// 2. strconv.ItoA
	ins02 := strconv.Itoa(ins01)
	fmt.Printf("%T  %v \n", ins02, ins02)
	// 选择使用 strconv.ItoA ，而不用如下的错误写法，因为 string() 方法中传入的参数是字节的 Utf8 编码
	// fmt.Printf("%T %v \n", string(ins01), string(ins01))

	// 3. Parse 系列
	ins03, _ := strconv.ParseFloat("3.1415926", 64)
	fmt.Printf("%T  %v \n", ins03, ins03)
	ins04, _ := strconv.ParseBool("true")
	fmt.Printf("%T  %v \n", ins04, ins04)

	// 4. Format 系列
	ins05 := strconv.FormatFloat(3.1415926, 'f', -1, 64) // 其中 -1 代表打印小数点后所有数字 fmt=f bitSize=64
	fmt.Printf("%T  %v \n", ins05, ins05)                // string  3.1415926
	ins06 := strconv.FormatFloat(3.1415926, 'E', 3, 64)  // 其中 3 代表着小数点后保留三位数字
	fmt.Printf("%T  %v \n", ins06, ins06)                // string  3.142E+00
	ins07 := strconv.FormatInt(5, 2)                     // 2 代表转换成二进制后，在格式化为字符串
	fmt.Printf("%T  %v \n", ins07, ins07)                // string  101
}
```

## log.set\* && log.Println && log.New

```go
func main() {
	// 普通日志
	log.Println("当前服务正在启动中...")
	// 不适用原生自带的 log 对象，新建一个 log 对象
	logger := log.New(os.Stdout, "Terminal Stdout：", log.Ldate|log.Llongfile|log.LUTC)
	logger.Println("当前服务正在运行中...")
	// 当调用了 log 对象的 Fatalln 方法、Panicln 方法时，后面的代码都不会运行
	logger.Fatalln("当前服务遇到了一个 Fatal 错误")
	// logger.Panicln("当前服务遇到了一个 Panic 错误")
	logger.Println("这行不会运行")
}

func init() {
	logFile, err := os.OpenFile("./review.log", os.O_CREATE|os.O_RDONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	// 设置日志格式
	log.SetFlags(log.Ldate | log.Llongfile | log.LUTC)
	// 设置日志前缀
	log.SetPrefix("Monitor App: ")
	// 设置日志输出，如果是想输出到控制台，就写 os.Stdout
	log.SetOutput(logFile)
}
```

## os.Open && os.Create（本质都是 os.OpenFile） 

## file.Read()

> 当前的 file 对象是由 os.OpenFile() 方法返回的

>  一次性全部读完，代码如下

```go
func main() {
	var once sync.Once
	file, err := os.OpenFile("./review.log", os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println(err)
	}
	// file.Read() 方法需要传入一个 []byte 切片，并且需要做了初始化
	var byteSlice []byte = make([]byte, 5)
	var content []byte
	for {
		bytesNumberAlreadyRead, err := file.Read(byteSlice)
        // 这行代码我放在 for 循环里只想输出一遍，所以用了 sync.Once 
		once.Do(func() {
			fmt.Println("每次接受的 []byte 切片容量是", bytesNumberAlreadyRead)
		})
		content = append(content, byteSlice...)
		if err == io.EOF {
			break
		}
	}
	file.Close()
	fmt.Println(string(content))
}
```

## ioutil.ReadFile（内置了 OpenFIle,无需自己 OpenFile）

> 一次性全部读完，代码如下

```go
func main() {
	res, err := ioutil.ReadFile("review.log")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(res))
}
```

## bufio.NewReader（带缓冲的读）

```go
func main() {
	file, err := os.OpenFile("review.log", os.O_RDONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
        // 一行一行的有缓冲的读
		fmt.Println(string(line))
	}
}
```

> 我们也可以像上面 file.Read() 方法一样设置缓冲区，并且还可以将所有缓冲到的内容赋值给一个 content，最后全部输出

```go
func main() {
	file, err := os.OpenFile("review.log", os.O_RDONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	var buf []byte = make([]byte, 5)
	var content []byte
	for {
		_, err := reader.Read(buf)
		if err == io.EOF {
			break
		}
		content = append(content, buf...)
	}
	fmt.Println(string(content))
}
```

## file.Write()

```GO
func main() {
	file, err := os.OpenFile("review.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	file.Write([]byte("Music loading...\n"))	// Write 传入 byte 切片
	file.WriteString("Playing...\n")	// 直接传入字符串
}
```

## ioutil.WriteFile（内置了 OpenFile）

```go
func main() {
	ioutil.WriteFile("./review.log", []byte("IOUtil's Music loading...\n"), 0666)
	res, _ := ioutil.ReadFile("review.log")
	fmt.Println(string(res))
}
// 因为 ioutil.WriteFile 内封装的 OpenFile 是如下权限，有 O_TRUNC
f, err := OpenFile(name, O_WRONLY|O_CREATE|O_TRUNC, perm)
// 所以每次写入都会直接清空
```

## bufio.NewWriter（带缓冲的写）

```go
func main() {
	file, err := os.OpenFile("review.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.Write([]byte("I am EThan\n"))
	writer.Flush()
}
```

## 并发写文件

```go
func ProdContent(channel chan interface{}) {
	var wg sync.WaitGroup
    // 模拟多个 Goroutine 写文件
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			// 如果加了 seed 种子和 UnixNano 时间关联，那么同一时间的数字是一样的，所以如果用了 rand.Seed 需要保证每个 goroutine 执行时间不同
			// rand.Seed(time.Now().UnixNano())
			// time.Sleep(1 * time.Nanosecond)	// 如果用了 Seed 就需要用这一行
			channel <- rand.Intn(100)
			defer wg.Done()
		}()
	}
	wg.Wait()
	close(channel)
}

func FileReceiver(file *os.File, writeChannel chan interface{}, isDoneChannel chan interface{}) {
	for i := range writeChannel {
		fmt.Fprintln(file, i)
	}
	isDoneChannel <- "Done"
}
func main() {
	var writeChan chan interface{} = make(chan interface{}, 3)
	var isDone chan interface{} = make(chan interface{})
	file, err := os.OpenFile("review.log", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	// 我们不需要管 ProdContent 函数是否执行完，因为 FileReceiver 函数中一直取 writeChan 信道中的值，信道不关闭，FileReceiver 函数不会执行完
	// 信道什么时候关闭？当 ProdContent 函数执行完才会关闭
	go ProdContent(writeChan)
	go FileReceiver(file, writeChan, isDone)
	// 用 isDone chan 等待 FileReceiver 函数执行完
	<-isDone
}
```

## rand.Read && rand.Int31() && rand.Int31n(100) && rand.New && rand.NewSource

```go
func main() {
	// 生成随机数（固定 5 位数）
	//rand.Seed(time.Now().UnixNano())
	//fmt.Printf("%.5d \n", rand.Int63()%10111)

	// rand.Read() 随机生成 byte 值赋予 []byte 切片（引入math/rand 和 crypto/rand 都可以使用）
	var byteSlice []byte = make([]byte, 5)
	rand.Read(byteSlice)
	fmt.Println(byteSlice)
	fmt.Println(string(byteSlice))
}
```

