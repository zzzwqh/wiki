---
title: Context && Reflect
date: 2022-12-06 09:32:08
tags: Golang
categories: Golang
---

> Context && Reflect 

<!--MORE-->

## 思考：WaitGroup 的功能是什么？

```go
// 本例子尝试验证，多层函数嵌套中，用一个 *sync.WaitGroup 变量，让主协程等待子/孙子协程的运行，是否可行（可行）
func process4WG2(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("WG2 ....") // 这行会打印出来
}

func process4WG1(wg *sync.WaitGroup) {
	defer wg.Done()
	wg.Add(1)          // 在子协程中写入 Add
	go process4WG2(wg) // 传入子协程的子协程
	fmt.Println("WG1 ....")
}
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go process4WG1(&wg)
	wg.Wait()
}
```

> 上面的例子中，我们用 WaitGroup 一直等待协程退出，要是 Main routine 想通知协程退出怎么办？比如协程工作太久，Main routine 想通知它超时结束，该用什么方法呢？

## 利用 Channel / 全局变量 控制子协程退出

> 全局变量的方式，通知和控制子协程退出

```GO
// 问题一：为什么需要 context?
// 答案: 为了通知 goroutine ,  优雅的让 goroutine 退出
// 问题二：如果没有 Context，我们如何 [通知] goroutine 运行退出?
// 思考：定义一个全局变量是否可行？
var wg4A sync.WaitGroup // wg 的作用只是让 Main 主协程等待每个 Goroutine 的运行，不是让 Main 协程通知到 Goroutine 退出
var notify bool = false

func processA() {
	defer wg4A.Done()
	for {
		fmt.Println("ProcessA Goroutine Task Running...")
		time.Sleep(time.Second)
		if notify == true {
			break
		}
	}
}

func main() {
	wg4A.Add(1)
	go processA()
	// 主协程等待 3 秒后，将 notify 全局变量设置为 true，即可达到通知 goroutine 退出效果
	for i := 0; i < 3; i++ {
		fmt.Println("Main Routine Running At Step", i)
		time.Sleep(time.Second)
	}
	notify = true
	wg4A.Wait()
}
```

> 利用管道，通知子协程优雅的退出

```go
// 问题：不适用全局变量，还有什么方法通知 goroutine 优雅的退出?
// 答案：Channel 管道

var wg4B sync.WaitGroup
var chan4B = make(chan struct{})

func processB() {
	defer wg4B.Done()
LOOP:
	for {
		fmt.Println("ProcessB Goroutine Task Running...")
		time.Sleep(time.Second)
		select {
		case chan4B <- struct{}{}:
			fmt.Println("Main routine give notice ~~")
			break LOOP
		default:

		}

	}
}

func main() {
	wg4B.Add(1)
	go processB()
	for i := 0; i < 3; i++ {
		fmt.Println("Main routine Running Step", i)
		time.Sleep(time.Second)
	}
	<-chan4B
	wg4B.Wait()
}
```

> 问题来了，如果是更复杂的 , 多层级的 Goroutine，用一个管道去控制一堆子协程是不行的，要用很多管道（使用起来相当复杂），所以引申出了 Context
>
> Context 会在主协程中生成一个根节点，可以将通知传递给每一个叶子节点（子协程/孙子协程） 

## Context 控制子协程退出

> 注意区分 Context 和 WaitGroup 的作用，Context 是让主协程通知子协程退出，WaitGroup 是让主协程等待协程运行完

```go
// 如何使用 Context 优雅的通知 Goroutine 退出
// 如果 wg 不定义为全局变量，而作为参数传入函数时，一定要记得将指针传入，否则会传入 wg 的副本！导致 Deadlock！
func processC(wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
LOOP:
	for {
		fmt.Println("ProcessC Goroutine Task Running...")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			fmt.Println("Exist Signal gave by Main routine ~")
			break LOOP
		default:

		}
	}

}

func main() {
	var wg4C sync.WaitGroup                                 // 为了让主协程等待子协程运行完
	ctx, cancel := context.WithCancel(context.Background()) // 为了让主协程通知（控制）子协程退出
	wg4C.Add(1)
	go processC(&wg4C, ctx)
	for i := 0; i < 3; i++ {
		fmt.Println("Main Routine Running Step", i)
		time.Sleep(time.Second)
	}
	cancel()
	wg4C.Wait()
}
```

> 多级 Goroutine 能不能都被 Main routine 通知退出？

```GO
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 多级 Goroutine 能不能都被 Main routine 通知退出？
func processE(wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
LOOP:
	for {
		fmt.Println("ProcessE Goroutine Task Running...")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			fmt.Println("ProcessE has received Exist Signal gave by ProcessD Goroutine ~")
			break LOOP
		default:

		}
	}
}

func processD(wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	wg.Add(1)
	go processE(wg, ctx)
LOOP:
	for {
		fmt.Println("ProcessD Goroutine Task Running...")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			fmt.Println("ProcessD has received Exist Signal gave by Main routine ~")
			break LOOP
		default:

		}
	}
}

func main() {
	var wg4D sync.WaitGroup                                 // 为了让主协程等待子协程运行完
	ctx, cancel := context.WithCancel(context.Background()) // 为了让主协程通知（控制）子协程退出
	wg4D.Add(1)
	go processD(&wg4D, ctx)
	for i := 0; i < 3; i++ {
		fmt.Println("Main Routine Running Step", i)
		time.Sleep(time.Second)
	}
	cancel()
	wg4D.Wait()
}

/* 输出结果如下，挺规矩的，如果是自己控制 channel 会比较麻烦
比如此例子中，主协程到子协程用的是一个 context 封装的 Channel，子协程和孙子协程用的其实是另外一个 context 封装的 Channel
Main Routine Running Step 0
ProcessD Goroutine Task Running...
ProcessE Goroutine Task Running...
ProcessE Goroutine Task Running...
ProcessD Goroutine Task Running...
Main Routine Running Step 1
ProcessD Goroutine Task Running...
ProcessE Goroutine Task Running...
Main Routine Running Step 2
ProcessE Goroutine Task Running...
ProcessD Goroutine Task Running...
processD has received Exist Signal gave by Main routine ~	// 收到信号的两行都打印出来了，这两行打印没有一定的先后顺序
processE has received Exist Signal gave by ProcessD Goroutine ~
*/
```

## Context.WithCancel()

```go
func generate(ctx context.Context) (ch <-chan int) {
	// 这里需要重新定义一个 dst 信道，不直接使用 ch 的原因是，ch 已经被定义成 only-receive 的信道，而该 gen 函数需要生成数据传入信道
	dst := make(chan int)
	var num int
	// 这里 gen 直接再次单独启用一个 Goroutine ，使用了刚刚定义的 dst 信道
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("cancel() 被调用了，ctx.Done() 的信道中传出值了...")
				return
			case dst <- num:
				num++

			}
		}
	}()
	// 让上面的 Goroutine 单独运行去，gen 函数先将 dst 信道返回给主协程
	return dst
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	// TIPS: 这里写 defer cancel() 是看不到 "cancel() 被调用了，ctx.Done() 的信道中传出值了..." 被打印的，因为主协程率先结束了
	defer cancel()

	// 循环从信道中取出元素（只要信道不关闭就不结束 for 循环）
	for elem := range generate(ctx) {
		fmt.Println(elem)
		if elem == 5 {
			break
		}
	}
	// TIPS: 想看到  "cancel() 被调用了，ctx.Done() 的信道中传出值了..." 被打印，就先调用 cancel()，再让主协程 sleep 一秒
	cancel()
	time.Sleep(time.Second)

}
```

## Context.Deadline()

```go
func main() {
   // 定义一个截至时间点
   deadLineTime := time.Now().Add(time.Second * 2)
   ctx, cancel := context.WithDeadline(context.Background(), deadLineTime)
   // 尽管 ctx 将过期，但最佳实践是
   // 在任何情况下都调用 Cancel() 函数
   // 否则 ctx 及其父级的生存时间可能会超过必要的时间，这也是为什么 WithDeadLine 函数仍然返回一个 cancel 函数的原因
   defer cancel()
   select {
   case <-time.After(3 * time.Second): // time.After(time time.Duration) 会等待 3 秒后，在返回的信道上发送当前时间
      fmt.Println("overslept")
   case <-ctx.Done(): // 设置了 DeadLine 的 ctx 在 2 秒后，在返回的信道上发送一个 struct{} 类型空值，那么上面的 select 语句就不会被执行
      fmt.Println(ctx.Err())
   }
}
```

## Context.Timeout()

```go
package main

import (
   "context"
   "fmt"
   "sync"
   "time"
)

func worker4Timeout(ctx context.Context, wg *sync.WaitGroup) {
   defer wg.Done()
   for {
      select {
      case <-ctx.Done():
         fmt.Println("Timeout Signal Received by Main routine ....")
         return // return 语句执行之前会执行 defer 哦，但是如果有返回值，会先赋值返回值（可复习 defer 的正确使用姿势）
      default:
         time.Sleep(time.Second)
         fmt.Println("Connecting Mysql Databases ...")
      }
   }
}
func main() {
   var wg sync.WaitGroup
   ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
   // 和 deadline 例子一样，在任何情况下都调用 cancel，避免上下文以及其父级生命周期过长，超过必要时间
   // 实际上，即使不调用 cancel，也会因为 Timeout 超时时间而结束 worker 携程
   defer cancel()
   wg.Add(1)
   go worker4Timeout(ctx, &wg)
   wg.Wait()
}
```

## Context.WithValue()

> 我们在传递 Value 时，最佳实践是定义一个新的类型，如本例中的 `type MachineInfo string`，是为了在传递上下文过程中，避免这个键值对，因为键的类型相同，导致值可以被重新赋值而覆盖

```go
func worker4WithValue(ctx context.Context, wg *sync.WaitGroup) {
   defer wg.Done()
   res, ok := ctx.Value(MachineInfo("SN")).(string)
   if !ok {
      fmt.Println("Invalid SN ...")
   }
LOOP:
   for {
      select {
      case <-ctx.Done(): // timeout 时间到，就会退出循环
         break LOOP
      default:
         fmt.Println("worker4WithValue Worker Get MachineInfo SN :", res)
         time.Sleep(time.Second * 1)
      }
   }
   // 我们再向下传递试试，看看在子协程中，再次定义传递 SN 到孙子协程中，是否会被覆盖
   type PersonInfo string // 为了避免上下文在协程之间传递时，被覆盖的问题，我们用 type 定义新的类型，不同类型的 Key 值无法覆盖
   ctx = context.WithValue(ctx, MachineInfo("SN"), "0000000")
   ctx = context.WithValue(ctx, MachineInfo("factory"), "英特尔")
   ctx = context.WithValue(ctx, PersonInfo("name"), "乔丹")
   wg.Add(1)
   go func(ctx context.Context, wg *sync.WaitGroup) {
      fmt.Println("=========> Grandson routine :")
      fmt.Println(ctx.Value(MachineInfo("SN"))) // 打印 0000000 也就是说会被覆盖
      fmt.Println(ctx.Value(MachineInfo("factory")))
      fmt.Println(ctx.Value(PersonInfo("name")))
      wg.Done()
   }(ctx, wg)
   // 这里不用再加 wg.Wait() 了哈，只需要主协程那里同一等待即可
}

// worker4WithValueBro 和 worker4WithValue 同为主协程的子协程，看看拿 MachineInfo 的值是否有问题（没有）
func worker4WithValueBro(ctx context.Context, wg *sync.WaitGroup) {
   defer wg.Done()
   res, ok := ctx.Value(MachineInfo("SN")).(string)
   if !ok {
      fmt.Println("Invalid SN ...")
   }
LOOP:
   for {
      select {
      case <-ctx.Done(): // timeout 时间到，就会退出循环，然后执行 wg.Done() 退出子 goroutine
         break LOOP
      default:
         fmt.Println("worker4WithValueBro Worker Get MachineInfo SN :", res)
         time.Sleep(time.Second * 1)
      }
   }

}

type MachineInfo string

func main() {
   var wg sync.WaitGroup
   ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
   defer cancel()
   // 传递一个键值对
   ctx = context.WithValue(ctx, MachineInfo("SN"), "1688888")
   wg.Add(2)
   go worker4WithValue(ctx, &wg)
   go worker4WithValueBro(ctx, &wg)
   wg.Wait()
}
```





## Reflect 反射的作用

1. Reflect 动态获取值的类型和信息（程序运行期间）

2. Reflect 实现了（程序运行期间）对值的修改能力

Go 是编译型语言，和动态语言不太一样，反射比较复杂，编译后变量会变成内存地址

Go 语言中，一个变量分为两个部分，类型信息和值信息，值信息（在程序运行期间）是可变化的



## reflect.TypeOf() 获取信息

> 结构体 Cat 有如下字段和方法

```go
type Cat struct {
	Alias string `json:"alias_name"`
	Age   uint8  `json:"age"`
	Kind  string `json:"kind"`
	Info  `json:"info"`
}
type Info struct {
	Hobby []string `json:"hobby"`
	Food  string   `json:"food"`
}

func (self *Cat) run() {
}
func (self *Cat) eat() {
}
func (self *Cat) Miao() {
}
func (self *Cat) Speak() {
}

```

**测试 reflect.TypeOf() 的一些使用场景：**

- ① 通过索引获取 Struct 字段信息

- ② 通过字段名获取 Struct 字段信息

- ③ Struct 中嵌套了 Struct ，可以用 FieldByIndex() 向下一层 Struct 获取字段

- ④ 通过索引获取 Struct 方法信息 

  

**<font size="6" color="red">如下代码中有些细节需要注意，比如何时需要用 Elem() 解引用 !!!!! 个人建议写法（重要总结）如下 =====></font>**

- Ⅰ. 不论是 reflect.TypeOf() 还是 reflect.ValueOf() 统统传入实例的地址指针（如下使用方法，是基于在传入的是实例的地址指针的）
- Ⅱ. 不论是 reflect.TypeOf() 还是 reflect.ValueOf() 返回的对象，使用 NumMethod()/Method() 等方法<font color="blue">**（方法相关）**</font>时，都不要用 Elem() ，否则方法数量不准确，会缺失指针接收器方法数量
- Ⅲ. 不论是 reflect.TypeOf() 还是 reflect.ValueOf() 返回的对象，使用 NumField()/FieldByName()/FieldByIndex()/Name()/Kind() 等方法<font color="blue">**（字段相关）**</font>时，都要用 Elem() ，解开引用，不然会报错 ~

```go
package main

import (
	"EthanCode/day06/entity"
	"fmt"
	"reflect"
)

type Fish struct {
	Name string
	Age  int
}

func (fish *Fish) Pop() {

}
func (fish *Fish) Speak() { // 这里用指针接收器

}
func (fish Fish) Swim() { // 这里用值接收器

}

// 如上 Fish 结构体一共三个方法，使用反射时，传入 Ptr 才能得到正确的方法数量，也就是这样写 reflect.ValueOf(&fishIns).NumMethod()，不要用 Elem()

func main() {
	var catIns = entity.Cat{}
	res4Ty := reflect.TypeOf(&catIns)                                 // 这里注意，传入了 Pointer，那么后面和 Field 有关的方法，都要有 Elem() 解开引用！！！
	fmt.Printf("%v %T\n", res4Ty, res4Ty)                             // *entity.Cat *reflect.rtype
	fmt.Printf("%v %T\n", res4Ty.Elem().Name(), res4Ty.Elem().Name()) // Cat string
	fmt.Printf("%v %T\n", res4Ty.Elem().Kind(), res4Ty.Elem().Kind()) // struct reflect.Kind 用 Elem() 解开了引用才是 struct，否则是 Ptr

	fmt.Printf("%v %T\n", res4Ty.Elem().NumField(), res4Ty.Elem().NumField()) // 4 int

	// 1. 通过索引，指定从 struct 中取出字段（可以获取字段名字等...）
	fmt.Println("============ reflect.TypeOf().FieldByName(\"xxx\")获取方法信息 ============> ")
	for i := 0; i < res4Ty.Elem().NumField(); i++ {
		field := res4Ty.Elem().Field(i) // 返回 StructField 类型
		fmt.Println(field)
		fmt.Printf("field.Name: %v , field.Type: %v , field.Index: %v , field.Tag: %v , field.Tag.Get(\"json\"): %v \n", field.Name, field.Type, field.Index, field.Tag, field.Tag.Get("json"))
	}

	// 2. 通过字符串，指定从 struct 中取出字段（缺点: 如果不知道字段名，那么就没法循环取出全部字段信息）
	fmt.Println("============ reflect.TypeOf().Elem().FieldByName(\"xxx\")获取方法信息 ============> ")
	field, _ := res4Ty.Elem().FieldByName("Info") // 返回 StructField 类型
	fmt.Printf("field.Name: %v , field.Type: %v , field.Index: %v , field.Tag: %v , field.Tag.Get(\"json\"): %v \n", field.Name, field.Type, field.Index, field.Tag, field.Tag.Get("json"))

	// 3. 如果 Struct 中嵌套了 Struct ，利用索引可以获取子结构体中的字段信息
	fmt.Println("============ reflect.TypeOf().FieldByIndex([]int)获取嵌套结构体信息 ============> ")
	fmt.Println(res4Ty.Elem().FieldByIndex([]int{3, 1})) // 找到索引为 3 的字段，是 Info ,在从 Info Struct 中找到索引为 1 的字段，是 Food
	fmt.Println(res4Ty.Elem().FieldByIndex([]int{3, 1}).Name)

	// 4. 获取方法的名字，这里不要用 Elem()，直接用 ！！！如果用了 Elem()，获取到的 NumMethod() 有问题
	fmt.Println("============ reflect.TypeOf().Method(i)获取方法信息 ============> ")
	fishIns := &Fish{Name: "ethan", Age: 1}
	// 传入地址！！！如果结构体方法使用了指针接收器，这里没有传入地址，那么将得不到正确的效果！！！如果传入的是值副本，那么只能得到值接收器的方法
	// 传入地址，无论值接收器类型方法，还是指针接收器类型方法，NumMethod() 获取到的方法数量都不会有偏差
	res4FishTy := reflect.TypeOf(fishIns)
	fmt.Println(res4FishTy.NumMethod()) // 先获取了方法的数量
	for i := 0; i < res4FishTy.NumMethod(); i++ {
		fmt.Println("=============>", res4FishTy.Method(i))
		fmt.Println("=============>", res4FishTy.Method(i).Name)
	}
	// 再试试获取 entity.Cat 类型的方法，只能获取到可访问的方法
	fmt.Println(res4Ty.NumMethod()) // 这里 CatIns 的方法只能
	for i := 0; i < res4Ty.NumMethod(); i++ {
		method := res4Ty.Method(i)
		fmt.Println(method)
	}

	// 5. 关于 Elem() 使用细节的验证
	fmt.Println("========== 关于 Elem() 使用细节的验证 ===========")
	// 这里传入的都是实例的地址指针 Ptr
	res4FishTy = reflect.TypeOf(fishIns)
	res4FishVal := reflect.ValueOf(fishIns)
	fmt.Println(res4FishTy.Elem().NumField())   // 字段数量 correct print 2
	fmt.Println(res4FishTy.NumMethod())         // 方法数量 correct print 3
	fmt.Println(res4FishTy.Elem().NumMethod())  // 方法数量 wrong print 1	也就是只拿到了值类型接收器方法 Swim，把指针类型接收器的方法忽略了
	fmt.Println(res4FishVal.Elem().NumField())  // 字段数量 correct print 2
	fmt.Println(res4FishVal.NumMethod())        // 方法数量 correct print 3
	fmt.Println(res4FishVal.Elem().NumMethod()) // 方法数量 wrong print 1	也就是只拿到了值类型接收器方法 Swim，把指针类型接收器的方法忽略了

}

/*
	- 关于 NumMethod()
	1. reflect.Type 类型（reflect.TypeOf()返回值）去调用 ===> reflect.TypeOf(&xxx).NumMethod()
	2. reflect.Value 类型（reflect.ValueOf()返回值）去调用 ===> reflect.ValueOf(&xxx).NumMethod()
	记住都如上写法，不要用 Elem() !!! 输出是符合预期的

	- 关于 NumField()
	1. reflect.Type 类型（reflect.TypeOf()返回值）:
	如果是传入了结构体地址的指针（Ptr），那么要用 Elem() 解引用
	如果是传入了结构体副本，那么不需要用 Elem() 解引用
	2. reflect.Value 类型（reflect.ValueOf()返回值）:
	传入方法地址指针！！！而且不要用 Elem()！！！
	示例:
	var catIns = entity.Cat{}
	res4Val := reflect.ValueOf(&catIns)
	fmt.Println(res4Val.NumMethod())
*/

```

## reflect.ValueOf() 获取信息

> 上面的 reflect.TypeOf() 的代码段落中，第 5 个例子已经参杂了一些 reflect.ValueOf() 的用法，获取 Field/Method  的相关信息

**测试更多的 reflect.ValueOf() 的使用场景：**

- ① 用 reflect.ValueOf() 获取值信息
- ② 注意 reflect.ValueOf() 获取的值类型，如何转换到需要的类型
- ③ 获取运行时变量的信息

```GO
package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Info
}
type Info struct {
	Account int
	Addr    string
}

func main() {
	// 1. 使用 reflect.ValueOf 可以获取变量值，传递给 ValueOf 的参数确保是一个地址
	user := User{Id: 1, Name: "root", Info: Info{Account: 500000, Addr: "Earth"}}
	val := reflect.ValueOf(&user)
	fmt.Println(val.Kind()) // Ptr
	// 虽然能够获取变量值，但是这个变量值的类型是 reflect.Value 类型，并不是 User 类型
	fmt.Println("================== reflect.ValueOf() ========================")
	fmt.Printf("%T, %v\n", val, val) // reflect.Value  &{1 root {500000 Earth}}

	// 2. 那么怎么把 Value 的 Type 恢复成 User？ 用 reflect.ValueOf().Interface() 先将 value 转成接口类型，然后使用类型断言
	// 如果不是结构体，将会方便很多，例如字符串类型，reflect.ValueOf().String()
	correctTy4UserIns := val.Interface().(*User)
	fmt.Println("=========== reflect.ValueOf().Interface() ===================")
	fmt.Printf("%T %v\n", correctTy4UserIns, correctTy4UserIns) // *main.User &{1 root {500000 Earth}}

	// 3. 获取运行时变量的信息
	// Elem returns the value that the interface v contains or that the pointer v points to.
	// It panics if v's Kind is not Interface or Pointer. It returns the zero Value if v is nil.
	// 可以清楚，使用 Elem() 的前提 v.Kind() 要么是 Interface 要么是指针（ptr），否则会 Elem() 引起 panic
	// 至于 val.Kind() 可以看 22 行代码，如果不取 &user 地址，val 将会是 struct 类型，从而导致 panic
	elem := val.Elem()
	fmt.Printf("%v\n", elem) // {1 root {500000 Earth}}

	fmt.Println("=========== reflect.ValueOf().elem().Type() ===================")
	elemType := elem.Type()
	fmt.Printf("%v\n", elemType) // main.User

	fmt.Println("=========== reflect.ValueOf().elem().Kind() ===================")
	elemKind := elem.Kind()
	fmt.Printf("%v\n", elemKind) // struct

	fmt.Println("=========== reflect.ValueOf().elem().NumField() ===================")
	numField := elem.NumField()
	fmt.Printf("%v\n", numField) // 3
	fmt.Println("=========== reflect.ValueOf().elem().Field(i) ==== 当前 struct 的字段值（返回类型是 reflect.Value） ===============")
	fmt.Println("=========== reflect.ValueOf().elem().Field(i).Interface() ==== 当前 struct 的字段值（返回类型不是 reflect.Value） ===============")
	fmt.Println("=========== reflect.ValueOf().elem().Field(i).Type() ==== 当前 struct 字段类型（返回 reflect.Type 类型）===============")
	fmt.Println("=========== reflect.ValueOf().elem().Type().Field(i) ==== 当前 struct 字段信息（返回 StructField 类型） ===============")
	fmt.Println("=========== reflect.ValueOf().elem().Type().Field(i).Name ==== 当前 struct 字段名字（String 类型）===============")
	fmt.Println("=========== reflect.ValueOf().elem().Type().Field(i).Type ==== 当前 struct 字段类型（reflect.Type 类型）===============")
	for i := 0; i < numField; i++ {
		fmt.Println("第", i, "个字段获取 elem.Field()")
		field := elem.Field(i)
		fmt.Printf("%T,%v\n", field, field)
		fieldName := elem.Type().Field(i).Name // reflect.ValueOf().elem().Type().Field(i).Name 获取字段名字
		fieldType := field.Type()
		fieldValue := field.Interface() // reflect.ValueOf().elem().field(i).Interface() 将值的类型返回，并转成了具体类型
		fmt.Printf("%d: %s %s = %v (类型:%T)\n\n", i, fieldName, fieldType, fieldValue, fieldValue)
		fmt.Println(reflect.TypeOf(fieldValue)) // 可以观察 fieldValue 也并非 Interface 类型，而是具体类型
	}

	fmt.Println("=========== reflect.ValueOf().elem().FieldByIndex() ==== 结构体嵌套结构体，深度遍历 ===============")
	fmt.Println(elem.FieldByIndex([]int{2, 1})) // 打印 Earth，打印了第 2 个字段（Info）的第 1 个字段（Addr）的值，也就是 Earth
}
```





 ## reflect.ValueOf() 修改结构体实例字段值

```GO
package main

import (
   "EthanCode/day06/entity"
   "encoding/json"
   "fmt"
   "reflect"
)

func main() {
   // 使用 reflect.ValueOf().Set*** 系列方法可以修改值
   strIns := "12345"
   val := reflect.ValueOf(&strIns) // 一定要加取地址符号，修改变量本身
   elem := val.Elem()
   elem.SetString("abcde") // 还有 SetInt 等待方法，不一一列举
   fmt.Println(elem)

   // 如果是结构体怎么修改？
   var dogIns = entity.Dog{Name: "ethan", Age: "3"}
   val4Dog := reflect.ValueOf(&dogIns)
   // 可以用如下方式去改
   val4Dog.Elem().FieldByName("Name").SetString("Noah")
   fmt.Println(val4Dog.Elem().FieldByName("Name"))
   // 如果不想一个个写，那么可以用下面的方式遍历哇，Elem() 是解引用 * 的，传入 reflect.ValueOf 函数中的是指针
   fmt.Println("先试试用 Reflect 遍历结构体字段 ===>")
   for i := 0; i < val4Dog.Elem().NumField(); i++ {
      fmt.Println(val4Dog.Type().Elem().Field(i).Name)
   }
   fmt.Println("有方法可以获取到字段，那么遍历修改 ===>")
   for i := 0; i < val4Dog.Elem().Type().NumField(); i++ {
      fieldName := val4Dog.Elem().Type().Field(i).Name

      val4Dog.Elem().FieldByName(fieldName).SetString("Test...") // 这里图个方便都赋值 Test...
      // 当然字段不一定都是 String 类型，这里可以写个函数用 Kind() 做判断，如果是 struct 就要递归，可参考 reflect_tag 中的代码，62 行
   }
   info, _ := json.MarshalIndent(dogIns, "", "\t")
   fmt.Println(string(info))
}
```

## reflect.ValueOf() 反射结构体方法并执行

```go
type Order struct {
   orderId string
   price   uint8
}

func (order *Order) PriceChange(price uint8) {
   order.price = price

}
func (order *Order) PrintPrice() {
   fmt.Println("Order Price is", order.price)
}
func (order *Order) GetPrice() (price uint8) {
   return order.price
}
func main() {
   var orderIns = &Order{orderId: "1688", price: 68}
   res4OrderTy := reflect.TypeOf(orderIns)
   res4OrderVal := reflect.ValueOf(orderIns) 
   fmt.Println(res4OrderVal.NumMethod())
   for i := 0; i < res4OrderVal.NumMethod(); i++ {
      methodByTy := res4OrderTy.Method(i)
      methodByVal := res4OrderVal.Method(i)
      fmt.Println("通过 reflect.TypeOf().Method(i) 返回 Method 类型名字/接收器等信息：", methodByTy) // 我们一般用这个
      fmt.Println("通过 reflect.ValueOf().Method(i) 返回地址：", methodByVal)

   }

   //  通过方法名字，获取到方法 ==> 调用方法
   fn4PriceChange := res4OrderVal.MethodByName("PriceChange")
   var args4PriceChange = []reflect.Value{reflect.ValueOf(uint8(28))}
   fn4PriceChange.Call(args4PriceChange) // 调用方法，需要传入参数
   fn4PrintPrice := res4OrderVal.MethodByName("PrintPrice")
   var args4PrintPrice = []reflect.Value{}
   fn4PrintPrice.Call(args4PrintPrice) // 如果不需要传入参数，传入空参数即可
   fn4GetPrice := reflect.ValueOf(orderIns).MethodByName("GetPrice")
   var args4GetPrice = []reflect.Value{}
   res := fn4GetPrice.Call(args4GetPrice)              // 返回的是什么对象？  []reflect.Value{}
   fmt.Printf("%v %T %v %T", res[0], res[0], res, res) // 返回的也是个 Slice，即一堆返回值的话，可以按索引获取
}
```