package main

import (
	"bytes"
	"strings"
)

// 函数功能：将一个表示整数的字符串，每隔三个字符插入一个逗号分隔符
// 12345 => 12,345
// 只支持整数类型

/*
	详细描述下递归过程：
		1. 第一层调用 执行递归 comma("1234567")+","+"890"
		2. 第二层调用 执行递归 comma("1234")+","+ "567"
		3. 第三层调用 执行递归 comma("1")+","+ "234"
		4. 满足基准情形，返回 "1"
		5. 进行回溯：
			 第三层 comma("1")返回1 因此返回值 "1,234"
			 第二层 comma("1234") 返回1234 因此返回值 "1,234,567"
			 第一层 comma("1234567") 返回1234567 因此返回值 "1,234,567,890"

*/

func comma(s string) string {
	n := len(s)

	if n <= 3 {
		return s
	}

	return comma(s[:n-3]) + "," + s[n-3:] //1,234,567,890
}

// 编写一个非递归版本的 comma 函数，使用bytes.Buffer代替字符串链接操作
func comma2(s string) string {
	var buf bytes.Buffer

	n := len(s)
	if n <= 3 {
		return s
	}

	firstPart := n % 3
	// 首先写入第一部分
	buf.WriteString(s[:firstPart])

	for i := firstPart; i < n; i += 3 {
		buf.WriteString(",")
		buf.WriteString(s[i : i+3])
	}

	return buf.String()
}

// 完善comma函数，以支持浮点数处理和一个可选的正负号处理

func comma3(s string) string {

	var buf bytes.Buffer

	// 查找小数点位置
	pointIndex := strings.Index(s, ".")
	if pointIndex == -1 {
		pointIndex = len(s)
	}

	// 处理整数部分
	// 整数部分字符串
	integerPart := s[:pointIndex]
	n := len(integerPart)

	// 第一部分长度
	firstPart := n % 3
	if firstPart == 0 {
		firstPart = 3
	}

	buf.WriteString(integerPart[:firstPart])

	// 对剩余的部分 每3位加一个逗号

	for i := firstPart; i < n; i += 3 {
		buf.WriteString(",")
		buf.WriteString(integerPart[i : i+3])
	}

	// 如果有小数部分 直接拼接
	if pointIndex < len(s) {
		buf.WriteString(s[pointIndex:])
	}

	return buf.String()
}

// 编写一个函数，判断两个字符串是否是互相打乱的，也就是说有这相同的字符串

func isAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	counts := make(map[rune]int)

	for _, c := range s1 {
		counts[c]++
	}

	for _, c := range s2 {
		counts[c]--

		if counts[c] < 0 {
			return false
		}
	}

	return true
}
