package main

// 专门用于处理 []int 类型的slice

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1

	if zlen <= cap(x) {
		// 还有足够的空间 延长切片
		z = x[:zlen]
	} else {
		// 没有足够的空间，分配一个新的数组
		// 加倍增长
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}
