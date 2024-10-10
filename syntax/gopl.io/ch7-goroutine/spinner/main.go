package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
	const n = 45

	start := time.Now() // 记录开始时间
	fibN := fib(n)
	elapsed := time.Since(start) // 计算经过的时间

	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
	fmt.Printf("计算耗时: %s\n", elapsed)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			// \r 是一个转义字符，表示回车，将光标移动到当前行的开头，而不换行
			// 在输出新的字符时，会覆盖当前行的内容
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
