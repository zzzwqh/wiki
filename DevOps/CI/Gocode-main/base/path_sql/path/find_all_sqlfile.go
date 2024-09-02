package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	// 找到当前目录下，所有 md 结尾的文件
	extName := ".md"
	mdFileSlice := make([]string, 1)
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("获取当前路径出错", err)
	}
	filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if filepath.Ext(path) == extName {
			fmt.Println(path)
			mdFileSlice = append(mdFileSlice, path)
		}
		return err
	})
	fmt.Println(mdFileSlice)
}
