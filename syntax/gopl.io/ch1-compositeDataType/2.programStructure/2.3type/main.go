package main

import "fmt"

func main() {
	var temp Celsius = 25.5
	fmt.Println(temp)
}

type Celsius float64

// 表明声明的是Celsius类型的一个叫Stirng的方法，该方法返回该类型对象对应的string方法
func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}
