package demo1

import (
	"fmt"
	"os"
	"path/filepath"
)

func OpenFile() {
	// 获取当前目录
	pwd, _ := os.Getwd()
	fmt.Println("当前目录为：", pwd)
	// 文件路径拼接(只拿到了ch6这一级)
	f1 := filepath.Join(pwd, "demo1", "1.txt")
	fmt.Println("当前拼接后：", f1)
	file, err := os.Open(f1)
	if err != nil {
		fmt.Println("文件打开出错,对应错误为：", err)
	}

	// 没有出错
	fmt.Printf("文件=%v\n", file)
	err1 := file.Close()
	if err1 != nil {
		fmt.Println("文件关闭出错,对应错误为：", err1)
	}
}
