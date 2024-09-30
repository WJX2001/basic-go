package main

import "fmt"

func main() {
	// demo1()
	// fmt.Println(f1() == f1()) // false
	demo2()
}

func demo1() {
	// 声明一个x变量
	x := 1
	// 取x变量的内存地址
	p := &x
	// *p 表达式对应p指针指向的变量的值，
	// *p表达式读取指针指向的变量的值
	fmt.Println(*p) // 1
	// *p 对应一个变量，可以更新指针所指向的变量的值
	*p = 2
	fmt.Println(x) // 2
}

var p = f1()

func f1() *int {
	v := 1
	return &v
}

// TODO: 通过指针来更新变量的值
/*
	通过指针来更新变量的值，返回更新后的值，可以在一个表达式中
	指针特别有价值的地方在于我们可以不用名字而访问一个变量
	但是缺点：
		要找到一个变量的所有访问者并不容易，必须知道变量全部的别名
*/

func incr(p *int) int {
	// 只是增加p指向的变量的值，并不改变p指针
	*p++
	return *p
}

func demo2() {
	v := 1
	// incr(&v)
	fmt.Println(incr(&v))
}
