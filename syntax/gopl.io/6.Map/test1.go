package main

import "fmt"

func main() {
	// test1()
	// dedupFunc()
	// charcount()
	// fmt.Println(hasEdge("1", "2"))
	wordfreq()
}

func test1() {
	// 写法一：内置函数
	// ages := make(map[string]int)

	// 写法二：字面值创建

	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
	}

	fmt.Println(ages)

	// 使用内置delete删除元素
	delete(ages, "alice")

	ages["bob"] += 1

	// 进行遍历
	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}

	if age, ok := ages["wjx"]; !ok { // wjx is not a key in this map
		fmt.Println(ok)
		fmt.Println(age, "wjx")
	}
}

// map之间不能进行相等比较；必须遍历

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}

	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv { // !ok 用来区分元素不存在
			return false
		}
	}
	return true
}

// Map的value类型也可以是一个聚合类型，比如一个map或者slice
var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}

	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}
