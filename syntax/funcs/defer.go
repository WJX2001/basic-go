package main

import "fmt"

// 是一个后进先出的特性
func Defer() {
	// 允许在返回的前一刻 执行一段逻辑，起到延迟调用的逻辑

	defer func() {
		println("第一个 defer")
	}()

	defer func() {
		println("第二个 defer")
	}()
}

func DeferClosure() {
	i := 0
	defer func() {
		println(i) // 1
	}()
	i = 1
}

func DeferClosureV1() {
	i := 0
	defer func(val int) {
		println(val) // 0
	}(i)
	i = 1
}

func DeferReturn() int {
	a := 0

	defer func() {
		a = 1
	}()
	return a // 0
}

func DeferReturnV1() (a int) {
	a = 0
	defer func() {
		a = 1
	}()
	return a // 1
}

// 自测题目

func DeferClosureLoopV1() {
	for i := 0; i < 10; i++ {
		defer func() {
			fmt.Printf("i 的地址 %p, 值是%d \n", &i, i) // %p 是十六进制的数
			println(i)                             // 10 10 10 10 10 10 10 10 10 10
		}()
	}
}

func DeferClosureLoopV2() {
	for i := 0; i < 10; i++ {
		defer func(val int) {
			fmt.Printf("i 的地址 %p, 值是%d \n", &val, val) // %p 是十六进制的数
			println(val)                               // 0-9
		}(i) // 9 8 7 6 5 4 3 2 1 0  // 因为defer是后进先出
	}
}

func DeferClosureLoopV3() {
	for i := 0; i < 10; i++ {
		j := i
		defer func() {
			fmt.Printf("i 的地址 %p, 值是%d \n", &j, j)
			println(j) // 9 8 7 6 5 4 3 2 1 0
		}()
	}
}
