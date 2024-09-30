package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

func demo1() {
	var a [3]int
	fmt.Println(a[0])

	// // 打印索引和对应的元素
	// for i, v := range a {
	// 	fmt.Printf("%d %d\n", i, v)
	// }

	// // 只打印元素
	// for _, v := range a {
	// 	fmt.Printf("%d\n", v)
	// }

	// 带有初始化
	// var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	fmt.Println(r[2]) // 0

	// 不指定长度
	q := [...]int{1, 2, 3}
	// q = [4]int{1,2,3,4}
	fmt.Println(q)
}

func demo2() {
	// 指定一个索引和对应值列表

	type Currency int

	const (
		USD Currency = iota // 美元
		EUR                 // 欧元
		GBP                 // 英镑
		RMB                 // 人民币
	)

	symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "￥"}
	fmt.Println(RMB, symbol[RMB]) // "3 ￥"
}

func demo3() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1) // false [32]uint8
}

// 可以显式传入一个数组指针，这样函数通过指针对数组的任何修改都可以直接反馈到调用者

func zero(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}

	// 另一种写法
	*ptr = [32]byte{}
}

// TODO: 编写一个函数，计算两个SHA256哈希码中不同bit的数目
func practiceDemo1(str1 string, str2 string) {
	c1 := sha256.Sum256([]byte(str1))
	c2 := sha256.Sum256([]byte(str2))

	// fmt.Printf("%x\n", c1)
	// fmt.Printf("%x\n", c2)
	fmt.Println(differentCount(c1, c2))
}

func differentCount(c1, c2 [32]byte) int {
	count := 0
	for i, _ := range c1 {
		count += popCount(c1[i], c2[i])
	}
	return count
}

// 可以参考popCount函数
func popCount(x, y byte) int {
	count := 0
	for i := 0; i < 8; i++ {
		// 取到二进制的最后一位
		xb := x % 2
		yb := y % 2

		if xb != yb {
			count++
		}
		// 向右移动一位
		x = x >> 1
		y = y >> 1
	}
	return count
}

/**
TODO:
	编写一个函数，默认情况下打印标准输入的SHA256编码，
	并支持通过命令行flag定制，输入SHA384或者SHA512哈希算法
*/

func Demo4() {
	var flag string
	if len(os.Args) > 1 {
		flag = os.Args[1]
	}
	// 标准输入的字符串
	var input string
	fmt.Println("请输入字符串：")
	fmt.Scanln(&input)
	if flag == "" {
		fmt.Println(sha256.Sum256([]byte(input)))
	} else if flag == "SHA384" {
		fmt.Println(sha512.Sum384([]byte(input)))
	} else if flag == "SHA512" {
		fmt.Println(sha512.Sum512([]byte(input)))
	} else {
		fmt.Println("请输入正确的加密方式")
	}
}
