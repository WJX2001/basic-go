package assertion

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

func (person Chinese) duan() {
	fmt.Println("断言喽")
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
func great(s SayHello) {
	s.sayHello()
	// 断言
	if ch, flag := s.(Chinese); flag {
		ch.duan()
	} else {
		fmt.Println("不是中国人")
	}

}

func Test() {
	chinese := Chinese{}
	american := American{}

	great(chinese)
	great(american)
}
