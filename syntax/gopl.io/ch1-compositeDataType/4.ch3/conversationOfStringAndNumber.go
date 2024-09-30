package main

import (
	"fmt"
	"strconv"
)

func conversationFunc() {
	// 整数到ASCII
	var i = strconv.Itoa(123)
	fmt.Println(i)

	// 用不同的进制格式化数字
	var i2 = strconv.FormatInt(int64(123), 2) // 转换为2进制
	fmt.Println(i2)

	// 将字符串解析为整数
	y, err := strconv.ParseInt("123", 10, 64) // 10进制，返回 64位int
	fmt.Println(y, err)
}
