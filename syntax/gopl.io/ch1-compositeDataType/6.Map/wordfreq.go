package main

import (
	"bufio"
	"fmt"
	"os"
)

// 写一个程序 报告输入文本中每个单词出现的频率
// 在第一次调用Scan前先调用input.Split(bufio.ScanWords),这样可以按单词而不是按行输入
func wordfreq() {
	fmt.Println("请输入")
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	tmpMap := make(map[string]int)
	for input.Scan() {
		line := input.Text()
		tmpMap[line]++
		fmt.Printf("单词 '%s' 出现了 %d 次\n", line, tmpMap[line])

	}

}
