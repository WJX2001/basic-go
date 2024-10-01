package demo2io

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func IOBuffer() {
	// 获取文件路径
	pwd, _ := os.Getwd()
	f1 := filepath.Join(pwd, "demo1", "1.txt")
	// 打开文件
	file, err := os.Open(f1)
	// 防止内存泄露
	defer file.Close()

	if err != nil {
		fmt.Println("文件打开失败，原因: ", err)
	}

	// 创建一个流
	reader := bufio.NewReader(file)
	// 读取操作
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF { // 已经读取到文件结尾
			break
		}
		// 正常输出文件内容
		fmt.Print(line)
	}

	fmt.Println("文件读取结束")

}
