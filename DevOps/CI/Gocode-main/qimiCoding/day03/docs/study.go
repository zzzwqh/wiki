package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	// 打开 excel 文件

	filename := "青浦专有云学习日报__2022-04-29.xlsx"
	f, err := excelize.OpenFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(filename)

	// 在 sheet0 中 A11 到 A17 单元格写入数据
	data := []string{
		"Kubernetes Pod 是 Kubernetes 中最小的部署单元，是一组容器的集合。",
		"Pod 中的容器共享相同的命名空间和网络空间，可以通过 localhost 相互通信。",
		"Pod 可以存储一些共享数据卷，使得多个容器可以访问同一份数据。",
		"Pod 中的容器可以使用相同的 IP 地址和端口号，从而实现负载均衡和高可用性。",
		"Kubernetes 使用 Pod 来进行应用程序的部署和管理，可以根据需要创建、删除和扩展 Pod。",
		"Kubernetes 还提供了多种控制器和调度器，如 ReplicaSet、Deployment 等，帮助用户更方便地管理 Pod。",
		"Pod 具备高可用、高性能、高灵活性的特点，适用于各种规模的应用程序部署。",
	}

	for i, v := range data {
		cell := fmt.Sprintf("A%d", i+11)
		f.SetCellValue("学习日报", cell, v)
	}
	cell := fmt.Sprintf("G%d", 7)
	f.SetCellValue("学习日报", cell, "补充学习专有云运维相关理论和知识体系")

	// 保存并关闭 excel 文件
	if err := f.Save(); err != nil {
		fmt.Println(err)
		return
	}
}
