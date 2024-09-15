package main

import "fmt"

// TODO:
/*
	创建变量的方法是调用内建的new函数，表达式new(T)，创建一个T类型的匿名变量，初始化T类型的零值，然后返回变量地址
*/

func main() {
	p := new(int)
	fmt.Println(p)

	*p = 2
	fmt.Println(*p)
}
