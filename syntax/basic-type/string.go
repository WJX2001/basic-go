package main

import (
	"fmt"
	"unicode/utf8"
)

func String() {
	println("hello" + "go")
	println(fmt.Sprintf("hello %d", 123))

	// 输出中文
	println(len("你好")) // 6

	println(utf8.RuneCountInString("你好")) // 2 正确写法
}
