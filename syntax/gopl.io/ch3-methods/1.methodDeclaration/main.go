package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func main() {

}

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// 基于指针对象的方法
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func Test() {
	p := Point{1, 2}
	pptr := &p
	pptr.ScaleBy(2)
	fmt.Println(p)
}
