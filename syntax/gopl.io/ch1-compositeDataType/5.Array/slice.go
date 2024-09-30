package main

import "fmt"

// TODO:
/**
1. slice三部分构成：指针、长度、容量
2. 指针指向第一个slice元素对应的底层数组元素的地址，slice的第一个元素并不一定是数组的第一个元素
3. 长度不能超过容量，容量一般是从slice的开始位置到底层数据的结尾位置
4. 内置的len和cap函数分别返回slice的长度和容量
*/

func sliceDemo1() {
	months := [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June", 7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}
	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2)     // [April May June]
	fmt.Println(summer) // [June July August]
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func testReverse() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a)
}

// 向左旋转n个元素
func rotate(s []int, n int) {
	// 三次调用reverse反转函数，第一次反转开头n个元素，然后反转剩下元素，最后反转整个slice元素
	reverse(s[:n])
	reverse(s[n:])
	reverse(s)
}

// TODO: slice之间不能比较，我们不能使用==操作符来判断两个slice是否含有全部相等元素
// 对于byte类型可以使用库， bytes.Equal函数来判断两个字节型是否相等,其他类型的必须展开比较

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}

	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

// apend 函数

func appendDemo() {
	var runes []rune
	for _, r := range "hello, 世界" {
		runes = append(runes, r)
	}

	fmt.Printf("%q\n", runes) // ['h' 'e' 'l' 'l' 'o' ',' ' ' '世' '界']
	// 类似写法
	var a = []rune("hello, 世界")
	fmt.Printf("%q\n", a)
}
