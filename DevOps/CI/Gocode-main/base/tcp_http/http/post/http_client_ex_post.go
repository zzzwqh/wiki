package main

import (
	"bufio"
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	params := url.Values{}
	params.Add("name", "ethan")
	params.Add("age", "25")
	resp, _ := http.PostForm("http://httpbin.org/post", params)
	var buf = make([]byte, 1024)
	reader := bufio.NewReader(resp.Body)
	reader.Read(buf)
	fmt.Println(string(buf))
}
