package main

import "fmt"

func main() {
	// var x, y int

	// // 提示用户输入
	// fmt.Println("请输入两个整数（以空格分隔）:")

	// // 读取用户输入的两个整数
	// fmt.Scan(&x, &y)

	// // 计算并打印最大公约数
	// fmt.Printf("最大公约数是: %d\n", gcd(x, y))

	fmt.Println(fib(10))
}

// 最小公约数
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

// 斐波那契数列
func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
