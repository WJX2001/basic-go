package main

import (
	"io"
	"log"
	"net"
	"os"
)

// 只关闭网络连接中写的部分，
// 这样的话后台goroutine可以在标准输入被关闭后继续打印从reverb1服务器传回的数据。

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	// 空结构体尺寸为0，通常被当作信号来传播
	ch := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		// 发送信号，让等待阻塞的main函数继续执行
		ch <- struct{}{}
	}()

	// 标准输入流数据 发送给服务器
	mustCopy(conn, os.Stdin)
	// 标准输入流关闭后，关闭conn的输入流
	// 告诉服务器不再发送数据，但仍保持读取状态
	conn.(*net.TCPConn).CloseWrite()
	// 阻塞等待服务端消息结束
	<-ch

}

func mustCopy(dst io.Writer, src io.Reader) {
	_, err := io.Copy(dst, src)
	if err != nil {
		println(err)
	}
}
