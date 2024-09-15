package main

import (
	"flag"
	"fmt"
	"strings"
)

// 命令行标志参数的名字"n" 标志参数的默认值:false 标志参数对应的描述信息
var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	// 程序运行时，必须在使用标志参数对应的变量之前先调用flag.Parse函数，
	/*
		程序运行时，必须在使用标志参数对应的变量之前先调用flag.Parse函数，用于更新每个标志参数（之前是默认值）。
		对于非标志参数的普通命令行参数可以通过调用flag.Args() 函数来访问，返回值对应一个字符串类型的slice
		如果在flag.Parse函数解析命令行参数时遇到错误，默认将打印相关的提示信息，然后调用os.Exit(2)退出程序。
	*/
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))

	// 如果-n标志被设置，则不打印换行符
	if !*n {
		fmt.Println()
	}
}
