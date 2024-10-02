package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

func main() {
	var address string
	// 通过命令行参数来获取端口
	// flag来通过 -port指定端口
	flag.StringVar(&address, "port", ":8000", "端口号，默认8000")
	flag.Parse()
	fmt.Println(address)
	// 开启服务
	ListenAndServer(string(":" + address))
}

func ListenAndServer(address string) error {
	// 端口监听
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	// 阻塞线程
	waitGroup := sync.WaitGroup{}
	// 循环监听
	for {
		conn, err := listener.Accept()
		// 新的连接等待组+1
		waitGroup.Add(1)
		if err != nil {
			break
		}

		// 协程处理对应连接
		go func(conn net.Conn) {
			// 协程结束 等待组-1
			defer waitGroup.Done()
			// 向客户端写入时间以及当前服务器端口
			io.WriteString(conn, time.Now().Format("15:04:05\n"))
		}(conn)

	}
	waitGroup.Wait()
	return nil
}
