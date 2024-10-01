package polymorphic

import "fmt"

// 定义一个接口
type SayHello interface {
	sayHello()
}

// 接口的实现
type Chinese struct {
	name string
}

func (person Chinese) sayHello() {
	fmt.Println("你好")
}

type American struct {
	name string
}

func (person American) sayHello() {
	fmt.Println("Hello")
}

type English struct {
}

// 只有满足SayHello 接口类型的对象才可以调用great方法
// 这个person 就是一个多态参数
func great(person SayHello) {
	person.sayHello()
}

func Test() {
	chinese := Chinese{}
	american := American{}

	great(chinese)
	great(american)
}

// 多态数组
func polymorphicArray() {
	var array [3]SayHello
	array[0] = Chinese{"王吉祥"}
	array[1] = American{"rose"}
}
