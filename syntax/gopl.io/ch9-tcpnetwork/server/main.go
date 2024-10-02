package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	// 连接用完一定要进行关闭
	defer conn.Close()

	for {
		// 创建一个切片，准备：将读取的数据放入切片
		buf := make([]byte, 1024)

		// 读取客户端发送的数据
		n, err := conn.Read(buf)
		if err != nil {
			return
		}
		// 将读取到的数据输出到终端
		fmt.Printf("客户端发送的数据为: %v\n", string(buf[:n]))

	}
}

func main() {

	fmt.Println("服务器端启动")
	// 进行监听，指定协议和服务器IP和端口号
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}

	// 监听成功
	// 循环等待客户端的连接
	for {
		conn, err2 := listen.Accept()
		if err2 != nil { // 客户端的等待失败
			fmt.Println("listen.Accept err:", err2)
			return
		} else {
			fmt.Printf("客户端连接成功: con=%v, 接收到的客户端信息为: %v\n", conn, conn.RemoteAddr().String())
		}

		// 准备一个协程，协程处理客户端服务请求
		go process(conn) // 不同的客户端的请求，连接的conn不一样的

	}

}
