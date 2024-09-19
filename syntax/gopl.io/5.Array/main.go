package main

import "fmt"

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
}
