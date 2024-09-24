package main

import (
	"bufio"
	"fmt"
	"os"
)

func dedupFunc() {
	seen := make(map[string]bool)
	fmt.Println("请输入")
	input := bufio.NewScanner(os.Stdin)

	// 输入单行
	// if input.Scan() {
	// 	line := input.Text() // 获取输入的文本行
	// 	fmt.Println("您输入的内容是:", line)
	// }
	for input.Scan() {
		line := input.Text()
		fmt.Printf("您输入的内容是: %s", line)
		if !seen[line] {
			seen[line] = true
		}
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
}

// TODO: 辅助函数
// 将slice转为map对应的string类型的key
func k(list []string) string {
	// fmt.Sprintf 用于格式化字符串并将结果返回
	// %q 表示双引号括起来的字符串
	return fmt.Sprintf("%q", list)
}

var m = make(map[string]int)

// 每次对map操作时先用k辅助函数将slice转为string类型
// 这样就实现了slice作为key
func Add(list []string) {
	m[k(list)]++
}

func Count(list []string) int {
	return m[k(list)]
}
