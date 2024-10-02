package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("客户端启动")
	// 调用Dial函数: 指定tcp协议，指定服务器端的IP+Port
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil { // 连接失败
		fmt.Println("客户端连接失败，err:", err)
		return
	}
	fmt.Println("客户端连接成功", conn)

	// 客户端发送单行数据
	reader := bufio.NewReader(os.Stdin)

	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("终端输入失败，err:", err)
		return
	}

	// 将str发送过去
	n, err := conn.Write([]byte(str)) // 使用短变量声明
	if err != nil {
		fmt.Println("发送数据失败，err:", err)
		return
	}
	fmt.Println("发送数据成功，n:", n)
}
