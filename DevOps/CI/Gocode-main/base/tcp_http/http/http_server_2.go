package main

import (
	"fmt"
	"net/http"
)

type indexHandler struct {
	content string
}

// 实现了 Handler 接口，Handler 接口中只含有一个方法，即 ServeHTTP，实际上和 HandelFunc 函数一样，是 func (w http.ResponseWriter, r *http.Request)
func (ih *indexHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// 这里很有意思，使用了 Fprintf 函数，把 ih.content 内容写入了 writer，而 writer 是 http.ResponseWriter 接口类型，这个接口类型实现了 io.Writer 接口类型
	fmt.Fprintf(writer, ih.content)
}
func main() {
	// 使用 http.Handle 方法，需要传入一个 Handler 类型的对象（或者实现了 Handler 接口的结构体，即本例中的 indexHandler）
	http.Handle("/", &indexHandler{content: "hello world!"})
	http.ListenAndServe(":8001", nil)
}

/*
	Handler 接口
 	type Handler interface {
  	  ServeHTTP(ResponseWriter, *Request)
	}
*/
