package main

import "fmt"

func LoopBug() {
	users := []User{
		{
			name: "wjx",
		},
		{
			name: "wjx2",
		},
	}

	m := make(map[string]*User)
	for _, u := range users {
		m[u.name] = &u // u 永远是一个地址
	}

	fmt.Printf("%v", m)
}

type User struct {
	name string
}
