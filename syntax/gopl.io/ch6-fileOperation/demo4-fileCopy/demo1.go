package demo4filecopy

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func Test() {
	path, _ := os.Getwd()
	fmt.Println(path)
	// 定义源文件地址
	srcFile := filepath.Join(path, "demo1", "1.txt")
	// 定义目标文件地址
	destFile := filepath.Join(path, "demo4-fileCopy", "copy.txt")

	// 对文件进行读取
	content, err := ioutil.ReadFile(srcFile)
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}

	// 对文件进行写入
	err = ioutil.WriteFile(destFile, content, 0666)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}

}
