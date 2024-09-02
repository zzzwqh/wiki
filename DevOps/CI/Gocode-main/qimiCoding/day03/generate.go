package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"io"
	"os"
	"path/filepath"
	"time"
)

func main() {
	// 原始文件名和路径
	srcFile := "example.xlsx"
	srcPath, err := filepath.Abs(srcFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 目标文件名前缀和路径
	destPrefix := "青浦专有云学习日报_"
	destPath, err := filepath.Abs(".")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 日期格式
	const layout = "2006-01-02"

	// 打开 Excel 文件

	// 循环拷贝文件和替换日期
	for i := 1; i <= 30; i++ {
		// 计算日期
		newDate := time.Date(2022, time.April, i, 0, 0, 0, 0, time.Local)
		newDateStr := newDate.Format(layout)
		newDateStr1 := newDate.Format("2006-01-02")
		// 构造目标文件名
		destName := fmt.Sprintf("%s_%s.xlsx", destPrefix, newDateStr)
		destFile := filepath.Join(destPath, destName)

		// 拷贝文件
		err := copyFile(srcPath, destFile)
		if err != nil {
			fmt.Printf("拷贝文件失败: %s\n", err)
			continue
		}

		// 打开目标 Excel 文件
		xlsxDest, err := excelize.OpenFile(destFile)
		if err != nil {
			fmt.Printf("打开目标文件失败: %s\n", err)
			continue
		}

		// 替换日期
		xlsxDest.SetCellValue("学习日报", "A2", newDateStr1)

		// 保存目标 Excel 文件
		err = xlsxDest.SaveAs(destFile)
		if err != nil {
			fmt.Printf("保存目标文件失败: %s\n", err)
			continue
		}

		fmt.Printf("已拷贝文件 %s 至 %s\n", srcPath, destFile)
	}

	fmt.Println("处理完成")
}

// 拷贝文件函数
func copyFile(src, dest string) error {
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	return err
}
