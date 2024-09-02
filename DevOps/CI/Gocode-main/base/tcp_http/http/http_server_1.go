package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/index", handleIndex)
	http.ListenAndServe(":9090", nil)
}

func handleRoot(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.RemoteAddr)
	fmt.Println(request.Body)
	writer.Write([]byte("Hello,This is handleRoot Page"))
}
func handleIndex(writer http.ResponseWriter, request *http.Request) {
	indexFileData, err := ioutil.ReadFile("D:\\Users\\ethan\\GolandProjects\\GoCode\\hexo-blog\\public\\index.html")
	if err != nil {
		fmt.Println("indexFileData Read Error:", err)
	}
	writer.Write(indexFileData)
}
