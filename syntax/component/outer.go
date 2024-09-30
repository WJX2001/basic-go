package main

type Inner struct{}

func (i Inner) DoSomething() {

}

type Outer struct {
	Inner // 组合类（不能循环嵌套）
}

func UserInnner() {
	var o Outer
	o.DoSomething()

}
