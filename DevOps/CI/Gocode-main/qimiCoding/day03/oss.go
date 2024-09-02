package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	// 替换为自己的 OSS Bucket 地址
	url := "https://ethan-myhexo.oss-cn-zhangjiakou.aliyuncs.com"

	// 获取当前时间
	now := time.Now().UTC()

	// 计算一小时前的时间
	oneHourAgo := now.Add(-1 * time.Hour)

	// 将时间转换为RFC1123格式，用于发送HTTP请求
	dateStr := oneHourAgo.Format(time.RFC1123)

	// 构造HTTP请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 添加HTTP头
	req.Header.Set("Date", dateStr)
	req.Header.Set("Authorization", "OSS AccessKeyId:Signature")

	// 发送HTTP请求
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 打印 IP 地址源
	fmt.Println(resp.Header.Get("X-Client-IP"))

	// 将 IP 地址源输出到文件中
	err = ioutil.WriteFile("ip.txt", []byte(resp.Header.Get("X-Client-IP")), 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Done")
}
