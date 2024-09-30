package main

type IntSet struct {
	words []uint64
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// func (s *IntSet) Add(x int) {
// 	word, bit := x/64, uint(x%64)
// 	for word >= len(s.words) {
// 		s.words = append(s.words, 0)
// 	}
// 	s.words[word] =
// }

type Couter struct {
	n int
}

func (c *Couter) N() int {
	return c.n
}

func (c *Couter) Increment() {
	c.n++
}

func (c *Couter) Reset() {
	c.n = 0
}
