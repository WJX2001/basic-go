package main

import (
	"fmt"
	"unicode/utf8"
)

// 前缀测试
func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func StringFix(s, prefix string) string {
	return s[:len(prefix)]
}

// 后缀测试
func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

// 包含子串测试
func Contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}

// 如果关心每个Unicode字符，可以使用其他处理方式
func UnicodeFunc() {
	s := "hello, 世界"
	// 在go里面"hello, 为ASCII字符，每个字符占一个字符，一共7个字符
	// “世界”是中文字符，为Unicode字符，在UTF-8编码下，每个字符占3个字节，共6个字节
	fmt.Println(len(s))                    // 13
	fmt.Println(utf8.RuneCountInString(s)) // 9
}

// 为了处理这些真实的字符，需要一个UTF8解码器
func UTF8Decoder(s string) {
	for i := 0; i < len(s); {
		// r对应字符本身，长度对应r采用UTF-8编码后的编码字节数目
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}
}

// range循环在处理字符串的时候，会自动隐式解码UTF8字符串
func RangeFunc(s string) {
	for i, r := range "Hello, 世界" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
}

// 讲一个整数转型为字符串
