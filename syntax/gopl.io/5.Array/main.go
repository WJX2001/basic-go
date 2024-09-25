package main

import (
	"fmt"
)

func main() {
	// demo1()
	// demo2()
	// demo3()
	// practiceDemo1("x", "X")
	// practiceDemo1("x", "X")
	// Demo4()
	// sliceDemo1()
	// testReverse()
	a := [6]int{0, 1, 2, 3, 4, 5}
	rotate(a[:], 2)
	fmt.Println(a)

	appendDemo()
	simulatedStack()

	// 删除元素测试
	s := []int{5, 6, 7, 8, 9}
	fmt.Println(deleteStackElement(s, 2)) // [5 6 8 9]

	s1 := []int{5, 6, 7, 8, 9}
	// 乱序删除
	fmt.Println(deleteStackElement2(s1, 2)) // [5 6 9 8]

	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(*(reverseByArrPointer(&arr)))
	// reverse(arr[3:])
	// fmt.Println(arr)

	arr1 := []string{"11", "11", "22", "23", "24", "24", "37"}
	fmt.Println(distinctFunc(arr1[:]))

	// str := []byte("afdigneigu  \n\t   fjidfiefh  jsdifj frewf      ")
	// // fmt.Println(utf8.DecodeRune(str))
	// for i := 0; i < len(str); i++ {
	// 	r, size := utf8.DecodeRune(str[i:])
	// 	fmt.Printf("%c %d\n", r, size)
	// }
}
