package main

import (
	"fmt"

	"github.com/WJX2001/basic-go/syntax/gopl.io/ch5-objectOriented/1.encapsulation"
	"github.com/WJX2001/basic-go/syntax/gopl.io/ch5-objectOriented/2.polymorphic"
	"github.com/WJX2001/basic-go/syntax/gopl.io/ch5-objectOriented/3.assertion"
)

func main() {
	p := encapsulation.NewPerson("wjx")
	p.SetAge(20)
	fmt.Println(p.GetAge())
	polymorphic.Test()
	assertion.Test()
}
