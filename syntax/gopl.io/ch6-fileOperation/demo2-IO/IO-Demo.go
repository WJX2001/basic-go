package demo2io

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func IoFunc() {
	// 获取当前目录
	pwd, _ := os.Getwd()
	f1 := filepath.Join(pwd, "demo1", "1.txt")
	// 读取文件
	content, err := ioutil.ReadFile(f1)

	if err != nil {
		fmt.Println("read file error")
	}

	// 如果读取成功，将内容显示在终端
	fmt.Printf("%v\n", string(content))
}
