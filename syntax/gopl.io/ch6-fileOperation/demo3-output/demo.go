package demo3output

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func Test() {
	// 获取打开文件的路径
	pwd, _ := os.Getwd()
	f1 := filepath.Join(pwd, "demo1", "1.txt")

	// 打开文件
	file, err := os.OpenFile(f1, os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("open file err:")
		return
	}
	defer file.Close()

	// 写入文件 利用IO流  -》 缓冲输出流
	writer := bufio.NewWriter(file)
	writer.WriteString("hello world\n")

	// 流带缓冲区，需要刷新数据 --> 写入文件
	writer.Flush()
}
