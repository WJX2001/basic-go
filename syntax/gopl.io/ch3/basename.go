package main

import (
	"strings"
)

// basename函数功能: 看起来像系统路径的前缀删除，同时将看似文件类型的后缀名部分删除

// 纯手写
func BasenameVersion1(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func BasenameVersion2(s string) string {
	slash := strings.LastIndex(s, "/") // -1 如果 "/" 没有被找到
	s = s[slash+1:]

	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}
