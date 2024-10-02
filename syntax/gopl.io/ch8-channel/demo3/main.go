package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 写：协程
func writeData(intChan chan int) {
	defer wg.Done()
	for i := 1; i <= 50; i++ {
		intChan <- i
		fmt.Println("写入数据：", i)
		time.Sleep(time.Second)
	}

	// 管道关闭
	close(intChan)
}

// 读：协程
func readData(intChan chan int) {
	defer wg.Done()
	for v := range intChan {
		fmt.Println("读取数据：", v)
		time.Sleep(time.Second)
	}

}

func main() {
	intChan := make(chan int, 50)
	wg.Add(2)
	// 开启 读和写协程
	go writeData(intChan)
	go readData(intChan)
	// 等待协程结束
	wg.Wait()
}
