package main

import (
	"fmt"

	"github.com/WJX2001/basic-go/syntax/gopl.io/2.programStructure/2.4packageAndFile/ch2/popcount"
)

func main() {

	// 将摄氏温度转换为华氏温度
	// fmt.Println(tempconv.CtoF(tempconv.BoilingC)) //212°F

	// // 将开尔文转为华氏
	// kelvin := tempconv.Kelvin(300)
	// celsius := tempconv.KToC(kelvin)
	// fmt.Println(celsius) // 26.850000000000023°C
	fmt.Println(popcount.PopCount(49))

}
