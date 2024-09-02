package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	// 预定义变量
	fmt.Println(string(filepath.Separator), string(filepath.ListSeparator)) //  / :

	// filepath.Abs() 返回一个绝对路径，如果传入参数是相对路径，基于当前目录，给出该相对路径的绝对路径
	fmt.Println("======== filepath.Abs() ========")
	relativePathStr := "../temp/go/"
	realDir, _ := filepath.Abs(relativePathStr)
	fmt.Println(realDir) // 如果当前路径是 C:\Users\ethan\go\src\myProject\  那么将打印 C:\Users\ethan\go\temp\go

	// filepath.Clean() 返回 path 的最短路径，去掉多余的 // 以及冗余的文件上下级目录
	fmt.Println("======== filepath.Clean() ========")
	complexDir := "/usr/../etc/../tmp"
	cleanDir := filepath.Clean(complexDir)
	fmt.Println(cleanDir) // 打印 /tmp

	// filepath.Rel() 返回 targetPath 相对 basePath 路径（两者都不能是相对路径，否则会报错）
	fmt.Println("======== filepath.Rel() ========")
	basePath, targetPath := "/usr/local", "/usr/src/go/bin"
	realDir, _ = filepath.Rel(basePath, targetPath)
	fmt.Println(realDir) // ../src/go/bin

	// filepath.EvalSymlinks() 将软连接文件的指向，以 string 类型返回（即下面的 realDir）
	fmt.Println("======== filepath.EvalSymlinks() ========")
	symlink := "/bin"
	realDir, _ = filepath.EvalSymlinks(symlink)
	fmt.Println(realDir) // linux 文件系统中，指向 /usr/bin

	// filepath.Ext() 返回文件路径的扩展名
	fmt.Println("======== filepath.Ext() ========")
	fpStr := "/usr/local/mysql/my.cnf"
	ext4fpStr := filepath.Ext(fpStr)
	fmt.Println(ext4fpStr)

	// filepath.Base() 返回 path 的尾部元素
	fmt.Println("======== filepath.Base() ========")
	httpsFileStr := "https://www.itsky.tech/2022/11/16/Golang-SQL-%E6%93%8D%E4%BD%9C/image-20221120214244373.png"
	fmt.Println(filepath.Base(httpsFileStr)) // 打印 image-20221120214244373.png
	fmt.Println(filepath.Base(fpStr))        // 打印 my.cnf
	fmt.Println(filepath.Base(targetPath))   // 打印 bin

	// filepath.Dir() 返回 path 的路径部分（和上面相反）
	fmt.Println("======== filepath.Dir() ========")
	fmt.Println(filepath.Dir(httpsFileStr)) // 打印 https:\www.itsky.tech\2022\11\16\Golang-SQL-%E6%93%8D%E4%BD%9C
	fmt.Println(filepath.Dir(fpStr))        // 打印 /usr/local/mysql/
	fmt.Println(filepath.Dir(targetPath))   // 打印 /usr/src/go/ ，如果 targetPath = /usr/src/go/bin/ 那么结果会是 /usr/src/go/bin（略有不同）

	// filepath.Match() 检查 path 是否符合正则匹配，感觉不是很好用，Linux 中结果参考每行注释（Windows 下结果不同）
	fmt.Println("======== filepath.Match() ========")
	fmt.Println(filepath.Match("/home/catch/*", "/home/catch/foo"))     // true
	fmt.Println(filepath.Match("/home/catch/*", "/home/catch/foo/bar")) // false
	fmt.Println(filepath.Match("/home/?opher", "/home/gopher"))         // true
	fmt.Println(filepath.Match("/home/\\*", "/home/*"))                 // true
	fmt.Println(filepath.Match("*.log", "/home/access.log"))            // false

	// 路径分隔符替换为 `/`
	fmt.Println("======== filepath.ToSlash() ========")
	root := `\usr\local\go`
	needToSlashDir := filepath.ToSlash(root)
	fmt.Println(needToSlashDir) // /usr/local/go

	// 路径分隔符替换成 `\`
	fmt.Println("======== filepath.FromSlash() ========")
	root = `/usr/local/go`
	needFromSlashDir := filepath.FromSlash(root)
	fmt.Println(needFromSlashDir) // \usr\local\go

	// filepath.SplitList() 分隔多个路径，windows 中分隔符是 ; Linux/Unix 中分隔符是 :
	fmt.Println("======== filepath.SplitList() ========")
	pathList := filepath.SplitList("/usr/bin;/usr/local/;/opt/")
	for _, path := range pathList {
		fmt.Println(path)
	}

	// filepath.Walk() 可以做到遍历一个路径下所有的文件，如下代码块中，我用 os.GetWd() 获取了当前路径，然后用 filepath.Walk() 遍历当前路径下所有文件
	// 然后将遍历到的所有文件路径，做判断，如果是 .sample 文件，那么就放到 sampleFileList 这个切片中
	// 就是说，这行代码实现了当前文件夹，找到所有拓展名为 .sample 的文件
	fmt.Println("======== filepath.Walk() ========")
	sampleFileList := make([]string, 1)
	rootDir, err := os.Getwd()
	if err != nil {
		fmt.Println(rootDir)
	}
	filepath.Walk(rootDir, func(path string, info fs.FileInfo, err error) error {
		fmt.Println(path)
		if filepath.Ext(path) == ".sample" {
			sampleFileList = append(sampleFileList, path)
		}
		return err
	})
	fmt.Println(".sample 文件的 Slice ==>")
	fmt.Println(sampleFileList)
}
