> 逃逸分析必读案例：https://www.cnblogs.com/brady-wang/p/15830088.html

> 堆栈内存相关概念：https://blog.csdn.net/fuhanghang/article/details/114527738


## Go 语言逃逸分析是什么
Go 语言，它的堆栈分配是通过 Compiler 进行分析，GC 去管理

逃逸分析就是在编译阶段（而非运行阶段）确定一个变量要放堆上还是栈上

逃逸分析规则如下：

- 是否有在其他地方（非局部）被引用。只要有可能被引用了，那么它一定分配到堆上。否则分配到栈上
- 即使没有被外部引用，但对象过大，无法存放在栈区上。依然有可能分配到堆上
对此你可以理解为，逃逸分析是编译器用于决定变量分配到堆上还是栈上的一种行为


## 怎么确定是否逃逸

第一，通过编译器命令，就可以看到详细的逃逸分析过程。而指令集 -gcflags 用于将标识参数传递给 Go 编译器，涉及如下：

- -m 会打印出逃逸分析的优化策略，实际上最多总共可以用 4 个 -m，但是信息量较大，一般用 1 个就可以了

- -l 会禁用函数内联，在这里禁用掉 inline 能更好的观察逃逸情况，减少干扰

> $ go build -gcflags '-m -l' main.go

第二，通过反编译命令查看

> $ go tool compile -S main.go

注：可以通过 go tool compile -help 查看所有允许传递给编译器的标识参数


## 为什么需要逃逸分析

为什么需要逃逸
这个问题我们可以反过来想，如果变量都分配到堆上了会出现什么事情？例如：

- 垃圾回收（GC）的压力不断增大 
- 申请、分配、回收内存的系统开销增大（相对于栈） 
- 动态分配产生一定量的内存碎片

其实总的来说，就是频繁申请、分配堆内存是有一定 “代价” 的。
会影响应用程序运行的效率，间接影响到整体系统。因此 “按需分配” 最大限度的灵活利用资源，才是正确的治理之道。
这就是为什么需要逃逸分析的原因，你觉得呢？

## 逃逸规律总结
要结合上述参考链接中的示例，需要做到的是掌握方法，遇到再看就好了

除此之外你还需要注意：

- 静态分配到栈上，性能一定比动态分配到堆上好 
- 底层分配到堆，还是栈。实际上对你来说是透明的，不需要过度关心 
- 每个 Go 版本的逃逸分析都会有所不同（会改变，会优化） 
- 直接通过 go build -gcflags '-m -l' 就可以看到逃逸分析的过程和结果 
- 到处都用指针传递并不一定是最好的，要用对


对于这块的知识点。建议适当了解，但没必要硬记。
靠基础知识点加命令调试观察就好了。
像是曹大之前讲的 “你琢磨半天逃逸分析，一压测，瓶颈在锁上”，完全没必要过度在意…