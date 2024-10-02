package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	ch <- 1 // 错误 死锁

	fmt.Println(ch)

}
