package main

import (
	"encoding/base64"
	"fmt"
	"github.com/xuri/excelize/v2"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	minSize = 3       // 最小字符串大小
	maxSize = 1299000 // 最大字符串大小
)

func main() {
	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())

	// 遍历当前目录下的所有文件
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 如果是 Excel 文件，则插入随机字符串
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".xlsx") {
			if err := insertRandomString(path); err != nil {
				fmt.Printf("插入随机字符串失败：%s，错误：%v\n", path, err)
			}
		}
		return nil
	})
	if err != nil {
		fmt.Printf("遍历文件夹出错：%v\n", err)
	}
}

func insertRandomString(filename string) error {
	// 打开 Excel 文件
	f, err := excelize.OpenFile(filename)
	if err != nil {
		return err
	}
	var size int
	// 插入字符串到 AE 列的第 186-300 行
	for i := 186; i <= rand.Intn(70)+186; i++ {
		// 生成随机字符串
		size = rand.Intn(maxSize-minSize+1) + minSize
		n := 10
		b := make([]byte, n)
		if _, err := rand.Read(b); err != nil {
			panic(err)
		}
		fmt.Println(base64.URLEncoding.EncodeToString(b))
		str := strings.Repeat(base64.URLEncoding.EncodeToString(b), size)
		cell := fmt.Sprintf("BQ%d", i)
		if err := f.SetCellValue("RDS 诊断报告", cell, str); err != nil {
			return err
		}
	}

	// 保存文件
	if err := f.Save(); err != nil {
		return err
	}

	fmt.Printf("已向 %s 的 AE 列的第 186-300 行插入了 %d 个 NUL,\n", filename, size)
	return nil
}
