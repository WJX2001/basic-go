package main

import (
	"unicode"
	"unicode/utf8"
)

const length = 10

func reverseFunc1(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// TODO: 重写reverse函数，使用数组指针代替slice
func reverseByArrPointer(arr *[length]int) *[length]int {
	for i := 0; i < length/2; i++ {
		// 换位
		(*arr)[i], (*arr)[length-1-i] = (*arr)[length-1-i], (*arr)[i]
	}
	return arr
}

// TODO: 编写一个rotate函数，通过一次循环完成旋转

// 循环，需要额外空间
func rangeRotate(slice []int, k int) []int {
	// 创建一个结果变量
	res := make([]int, len(slice))
	for i, v := range slice {
		res[(i+k)&len(slice)] = v
	}
	return res
}

// 通过反转完成，不需要额外空间

// 方法一：旋转切片()
func rotateSliceAll(slice []int, k int) {
	// 全部反转
	// reverse(slice)

	// 2.局部反转
	reverse(slice[k:])
	reverse(slice[:k])
}

// TODO: 写一个函数在(原地完成)消除[]string中相邻重复的字符串的操作

func distinctFunc(s []string) []string {
	// 额外空间
	// slice := make([]string, 0)

	// for i := 0; i < len(s); i++ {
	// 	if i == 0 || s[i] != s[i-1] {
	// 		slice = append(slice, s[i])
	// 	}
	// }
	// return slice

	// 原地完成
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			// 跳过第i个元素
			copy(s[i:], s[i+1:])
			s = s[:len(s)-1]

		}
	}
	return s
}

// TODO: 编写一个函数，原地将一个UTF-8编码的[]byte类型的slice中相邻的空格替换成一个空格返回
// 参考unicode.IsSpace
// 判断是否是空格

func replaceAdjacentSpaces(slice []byte) []byte {
	length := len(slice)
	// length 表示切片的容量
	res := make([]byte, 0, length)

	for i := 0; i < length; {
		// 解码出第一个Unicode字符 例如a的话就是97
		r, size := utf8.DecodeRune(slice[i:])
		if unicode.IsSpace(r) {
			// 如果当前字符是空格
			res = append(res, ' ') // 将第一个空格添加到结果切片
			for i+size < length {
				nextRune, nextSize := utf8.DecodeRune(slice[i+size:])
				if !unicode.IsSpace(nextRune) {
					break // 直到找到下一个非空字符
				}
				size += nextSize
			}
		} else {
			res = append(res, slice[i:i+size]...)
		}
	}
	return res
}
