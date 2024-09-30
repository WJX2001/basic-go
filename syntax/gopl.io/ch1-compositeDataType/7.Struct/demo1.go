package main

func demo1() {

	type address struct {
		hostname string
		port     int
	}

	hits := make(map[address]int)
	hits[address{"golang.org", 443}]++
}
