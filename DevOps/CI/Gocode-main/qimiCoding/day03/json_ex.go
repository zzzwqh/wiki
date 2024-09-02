package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	// 打开Excel文件
	f, err := excelize.OpenFile("./example.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 修改单元格的值
	f.SetCellValue("Sheet1", "A2", "New Value")

	// 保存更改并关闭文件
	if err := f.Save(); err != nil {
		fmt.Println(err)
		return
	}
}
