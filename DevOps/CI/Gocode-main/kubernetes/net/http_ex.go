package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
)

type ImageInfo struct {
	RepoNamespace string `json:"reponamespace"`
	RepoName      string `json:"reponame"`
}

func main() {

	http.HandleFunc("/api/get", ApiDealGetHandler)
	http.HandleFunc("/api/post", ApiDealPostHandler)
	http.HandleFunc("/api/get/repo", ApiDealGetRepoHandler)

	http.ListenAndServe(":9090", nil)
}
func ApiDealGetRepoHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	repoNamespace := query.Get("repoNamespace")
	repoName := query.Get("repoName")
	log.Println("从 URL 中获取到的 repoNamespace 值为 => ", repoNamespace)
	log.Println("从 URL 中获取到的 repoName 值为 => ", repoName)
	cmd := exec.Command("/bin/bash", "-c", fmt.Sprintf("/usr/bin/aliyun cr GetRepoTags --RepoNamespace %s --RepoName %s", repoNamespace, repoName))
	jsonRes, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("ApiDealGetRepoHandler() Error:", err)
		//	return	// 如果打开 return ，那么执行命令获得到的 [正确/错误] 内容也不会打印了，只会打印 ApiDealGetRepoHandler() Error: exit status 1
	}
	fmt.Println(string(jsonRes))
	w.Write([]byte(jsonRes))
	/*
		这个函数会打印如下内容
		Exec4AllRes() Error: exit status 1
		CentOS Linux release 7.6.1810 (Core)
		cat: /etc/pam.d/: Is a directory
	*/
}

func ApiDealGetHandler(w http.ResponseWriter, r *http.Request) {
	// r.URL.Query() 返回一个 url.Values 类型的对象，其中包含了请求 URL 中的所有参数和对应的值。通过访问这个对象，可以获取特定参数的值。
	query := r.URL.Query()
	repoNamespace := query.Get("reponamespace") // 使用 Query().Get 方法，默认会返回 []string 切片的第一个值（为什么是切片，因为可能有 name=ethan&&age=18&&name=noah 的情况）
	repoName := query.Get("reponame")
	imageInfoIns := &ImageInfo{RepoNamespace: repoNamespace, RepoName: repoName}
	log.Println("收到客户端发起的 Get 请求，Body 体内容为: ", imageInfoIns)
	json.NewEncoder(w).Encode(imageInfoIns) // 这行代码，其实和下面两行代码实现的功能一致，Encode() 将结构体序列化，NewEncoder(w) 创建了一个新的编码器，用于将 JSON 编码后的数据写入到指定的输出流 w 中
	//tagDataInsStr, _ := json.Marshal(tagDataIns)
	//w.Write([]byte(tagDataInsStr))

}
func ApiDealPostHandler(w http.ResponseWriter, r *http.Request) {
	bodyBytesSlice, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	log.Println("收到客户端发起的 Post 请求，Body 体内容为: ", string(bodyBytesSlice))

}
