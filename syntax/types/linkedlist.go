package main

import "fmt"

type LinkedList struct {
}

// 方法接受器(方法定义在LinkedList上)
func (l LinkedList) Add(idx int, val any) {

}

// 方法接受器（方法定义在 *LinkedList上）
func (l *LinkedList) Add1(idx int, val any) {

}

type User struct {
	Name      string
	FirstName string
	Age       int
}

func (u User) ChangeName(name string) {
	fmt.Printf("change name中u的地址 %p \n", &u)
	u.Name = name
}

// 声明为指针，可以修改
func (u *User) changeAge(age int) {
	fmt.Printf("change age中u的地址 %p \n", u)
	u.Age = age
}

func changeUser() {
	// 这里的u1地址和changeAge中的u地址是一样的
	u1 := User{Name: "Tom", Age: 20}
	fmt.Printf("u1的地址 %p \n", &u1)
	// 值传递的时候，如果传递的是指针，则会修改
	u1.changeAge(35)
	// 执行这一步的时候，相当于复制了一个u1，改的是复制体，所以name没有修改
	u1.ChangeName("Jerry")
	fmt.Printf("%+v", u1)

	// 声明为指针去调用结构体
	up1 := &User{}
	up1.changeAge(35)
	up1.ChangeName("Jerry")
	fmt.Printf("%+v", up1)
}

// 结构体自引用只能用指针
// type node struct {
// 	prev *node // 取指针
// 	next *node // 取指针
// }
