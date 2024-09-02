package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	// 获取当前目录
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("获取当前目录失败:", err)
		return
	}

	// 遍历当前目录下的所有Excel文件
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("遍历目录时发生错误:", err)
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".xlsx" {
			// 打开Excel文件
			f, err := excelize.OpenFile(path)
			if err != nil {
				fmt.Printf("打开Excel文件(%s)失败: %v\n", path, err)
				return nil
			}

			// 获取第一个sheet的名称
			sheetName := f.GetSheetName(0)
			// 插入随机百分数到L列22行
			percent := rand.Float64()*(6.5-5.8) + 5.8
			f.SetCellValue("总览", "L22", fmt.Sprintf("%.2f%%", percent))
			randInt := rand.Intn(76500-75500+1) + 75500
			f.SetCellValue(sheetName, "J22", strconv.Itoa(randInt))
			randInt = rand.Intn(1225694-1215694+1) + 1215694
			f.SetCellValue(sheetName, "K22", strconv.Itoa(randInt))
			// 插入随机百分数到L列27行
			percent = rand.Float64()*(0.80-0.69) + 0.69
			f.SetCellValue(sheetName, "L27", fmt.Sprintf("%.2f%%", percent))

			randInt = rand.Intn(265-240+1) + 240
			f.SetCellValue(sheetName, "J27", strconv.Itoa(randInt))
			randInt = rand.Intn(32683-31683+1) + 31683
			f.SetCellValue(sheetName, "K27", strconv.Itoa(randInt))

			// 插入随机百分数到L列29行
			percent = rand.Float64()*(8.30-8.02) + 8.02
			f.SetCellValue(sheetName, "L29", fmt.Sprintf("%.2f%%", percent))
			randInt = rand.Intn(43669-42669+1) + 42669
			f.SetCellValue(sheetName, "J29", strconv.Itoa(randInt))
			randInt = rand.Intn(499638-489638+1) + 489638
			f.SetCellValue(sheetName, "K29", strconv.Itoa(randInt))
			// 插入随机百分数到L列32行
			percent = rand.Float64()*(12.80-11.2) + 11.2
			f.SetCellValue(sheetName, "L32", fmt.Sprintf("%.2f%%", percent))
			randInt = rand.Intn(39254-37254+1) + 37254
			f.SetCellValue(sheetName, "J32", strconv.Itoa(randInt))
			randInt = rand.Intn(305437-295437+1) + 295437
			f.SetCellValue(sheetName, "K32", strconv.Itoa(randInt))
			// 插入随机百分数到L列36行
			percent = rand.Float64()*(1.28-1.09) + 1.09
			f.SetCellValue(sheetName, "L36", fmt.Sprintf("%.2f%%", percent))
			randInt = rand.Intn(6350-6050+1) + 6050
			f.SetCellValue(sheetName, "J36", strconv.Itoa(randInt))
			randInt = rand.Intn(540416-530416+1) + 530416
			f.SetCellValue(sheetName, "K36", strconv.Itoa(randInt))
			// 插入随机百分数到L列37行
			percent = rand.Float64()*(0.20-0.17) + 0.17
			f.SetCellValue(sheetName, "L37", fmt.Sprintf("%.2f%%", percent))
			randInt = rand.Intn(1603-1503+1) + 1503
			f.SetCellValue(sheetName, "J37", strconv.Itoa(randInt))
			randInt = rand.Intn(915394-905394+1) + 905394
			f.SetCellValue(sheetName, "K40", strconv.Itoa(randInt))
			// 插入随机百分数到L列40行
			percent = rand.Float64()*(0.25-0.21) + 0.21
			f.SetCellValue(sheetName, "L40", fmt.Sprintf("%.2f%%", percent))
			randInt = rand.Intn(888-842+1) + 842
			f.SetCellValue(sheetName, "J40", strconv.Itoa(randInt))
			randInt = rand.Intn(371917-361917+1) + 361917
			f.SetCellValue(sheetName, "K40", strconv.Itoa(randInt))

			// 增长率部分....
			percent = rand.Float64() * (0.3 - 0.0)
			f.SetCellValue(sheetName, "M29", fmt.Sprintf("%.1f%%", percent))
			percent = rand.Float64() * (0.4 - 0.0)
			f.SetCellValue(sheetName, "M32", fmt.Sprintf("%.1f%%", percent))
			percent = rand.Float64() * (0.5 - 0.0)
			f.SetCellValue(sheetName, "M36", fmt.Sprintf("%.1f%%", percent))
			percent = rand.Float64() * (0.6 - 0.0)
			f.SetCellValue(sheetName, "M37", fmt.Sprintf("%.1f%%", percent))

			// 保存文件
			if err := f.Save(); err != nil {
				fmt.Printf("保存Excel文件(%s)失败: %v\n", path, err)
			} else {
				fmt.Printf("已向%s的%s中插入%.2f%%\n", path, sheetName, percent)
			}
		}
		return nil
	})
}
