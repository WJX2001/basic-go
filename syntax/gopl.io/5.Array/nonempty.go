package main

import "fmt"

// 返回不包含空字符串的切片
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

// 使用append函数实现
func nonempty2(strings []string) []string {
	// 复制一组
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

// TODO: slice 模拟栈，使用append 实现

func simulatedStack() {
	stack := make([]string, 0)
	stack = append(stack, "1", "2", "3")

	fmt.Println(stack)
	// stack 顶部位置对应slice的最后一个元素(先进后出) 所以从后往前找
	top := stack[len(stack)-1]
	fmt.Println(top)

	// 通过收缩stack 可以弹出栈顶的元素
	stack = stack[:len(stack)-1]
	fmt.Println(stack)

}

// 删除元素后保留原来的顺序
func deleteStackElement(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	var res = slice[:len(slice)-1]
	return res
}

// 删除元素后不用保持原来顺序的话，用最后一个元素覆盖被删除的元素
func deleteStackElement2(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}
