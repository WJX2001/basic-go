package main

import "fmt"

func main() {

	// 定义管道
	var intChan chan int
	// 通过make初始化, 可以存放3个int类型的数据
	intChan = make(chan int, 3)

	// 证明管道是引用类型：
	fmt.Printf("intChan的值: %v\n", intChan) // 地址：0x1400012c000

	// 向管道存放数据
	intChan <- 10
	num := 20
	intChan <- num
	// 注意：不能存放大于容量的数据
	intChan <- 30

	// 输出管道的长度
	fmt.Printf("intChan的长度: %v, 管道的容量是：%v\n", len(intChan), cap(intChan))

	// 从管道中取数据
	num1 := <-intChan
	fmt.Printf("num1的值: %v, intChan的长度: %v\n", num1, len(intChan))
}
