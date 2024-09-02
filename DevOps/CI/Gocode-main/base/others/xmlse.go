package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Device struct {
	// 此处只可以赋值 XMLName 获取，很奇怪，用别的比如 Name/XmlName 都不行
	XMLName xml.Name `xml:"devices"`
	Host    []Hosts  `xml:"host"`
	Version string   `xml:"version,attr"`
	//Content string   `xml:",innerxml"`	// Content 即 Xml 所有内容，如果解开注释，会重复打印
	Comment string `xml:",comment"` // 注释 <!-- Comment -->
}

type Hosts struct {
	Id       int    `xml:"id,attr"`
	HostName string `xml:"hostName"`
	HostCode string `xml:"hostCode"`
	HostDate string `xml:"hostDate"`
	Comment  string `xml:",comment"`
}

func main() {
	// 反序列化
	var deviceIns Device
	res, err := ioutil.ReadFile("./device.xml")
	if err != nil {
		fmt.Println(err)
	}
	err = xml.Unmarshal(res, &deviceIns)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("=================")
	fmt.Println(deviceIns.Comment)
	fmt.Println("=================")

	fmt.Println("Server List :")
	for _, v := range deviceIns.Host {
		fmt.Println(v.Id, v.HostName, v.HostCode, v.HostDate)
	}

	// ============================================== //
	// 加个 Host 进去
	var hostIns = Hosts{HostName: "镜像仓库服务器", HostCode: "123456", HostDate: "2022-01-01", Id: 4, Comment: "新加入的 Host 机器"}
	deviceIns.Host = append(deviceIns.Host, hostIns)
	fmt.Println("After Change, Server List :")
	for _, v := range deviceIns.Host {
		fmt.Println(v.Id, v.HostName, v.HostCode, v.HostDate)
	}

	// 序列化的方法
	res4Marshal, err4Marshal := xml.Marshal(deviceIns)
	// 加一个头部 Head
	headByte := []byte(xml.Header)
	res4All := append(headByte, res4Marshal...)

	if err4Marshal != nil {
		fmt.Println(err4Marshal)
	}
	fmt.Println(string(res4All))

	// 将写入文件
	file, err4file := os.OpenFile("./seres.xml", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	if err4file != nil {
		fmt.Println("File Open err,", err4file)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.Write(res4All)
	writer.Flush()

}
