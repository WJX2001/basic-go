package main

import "fmt"

func main() {

	// 定义管道、声明管道
	var intChan chan int
	intChan = make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan <- i
	}
	// 在遍历前，需要关闭管道
	close(intChan)
	for v := range intChan {
		fmt.Println("vakue = ", v)
	}
}
