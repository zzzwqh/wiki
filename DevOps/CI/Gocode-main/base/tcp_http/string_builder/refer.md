

与许多支持类型的语言一样，golang中的string类型也是只读且不可变的，如果使用‘+’操作符或者fmt.Sprinf进行拼接，会导致字符串创建、销毁、内存分配、数据拷贝等操作，在高并发系统中不得不考虑更优的解决方案。bytes.Buffer提供了较好的解决方法。

bytes.Buffer内部使用[]byte来存储写入的数据，在一定程度避免了每次数据写入都重新分配内存和数据拷贝操作。但要注意buf.String()方法会进行[]byte到string的类型转换，最终还是会导致一次内存申请和数据拷贝。为了避免这一次的性能消耗，golang标准库中提供了strings.Builder，这个包在功能上和bytes.Buffer基本上是一致的。

strings.Builder 和 bytes.Buffer 的使用方法几乎一模一样，
两者都是通过WriteString()来做字符串的拼接，都是通过String()来获得拼接后的字符串。不过根据 Go 官方的说法 strings.Builder 在内存使用上的性能要略优于 bytes.Buffer，
这里推荐按照官方的建议，在做字符串的拼接时使用 strings.Builder，
做字节的拼接时使用bytes.Buffer。 
>https://www.bilibili.com/read/cv14762497/

