package main

import "fmt"

func main() {
	// fmt.Println(BasenameVersion2("a/b/c.go")) // c
	// fmt.Println(BasenameVersion2("c.d.go"))   // c.d
	// fmt.Println(BasenameVersion2("abc"))      // abc

	// fmt.Println(initsToString([]int{1, 2, 3}))
	fmt.Println(comma2("+1234567890"))
	fmt.Println(comma2("-1234567890"))
	fmt.Println(comma3("1444.2345"))
	fmt.Println(isAnagram("abcd", "dcba"))
	conversationFunc()
}
