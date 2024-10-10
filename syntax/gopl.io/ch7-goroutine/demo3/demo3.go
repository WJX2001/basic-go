package demo3

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func Demo3() {

	// 启动五个协程
	for i := 1; i <= 5; i++ {
		go func(n int) {
			fmt.Println(n)
			wg.Done() // 协程执行完成减1
		}(i)
	}
	wg.Wait() // 等待所有协程执行完成
}
