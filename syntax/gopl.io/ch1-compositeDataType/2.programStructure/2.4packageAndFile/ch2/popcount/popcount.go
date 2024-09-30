package popcount

var pc [256]byte

func Init() {
	for i := range pc {
		// i/2 用于获取i的右移一位的值，i&1用于获取i的最低位
		// pc[0] = 0  00000000
		// pc[1] = 1  00000001
		// pc[2] = 1  00000010
		// pc[3] = 2  00000011
		// pc[4] = 1  00000100 -> pc[2]+byte(4&1) 因为4的二进制表示是00000100 最低有效位位0，意味着没有额外的1要添加
		pc[i] = pc[i/2] + byte(i&1)
	}
	// fmt.Println(pc[4])
}

func PopCount(x uint64) int {

	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// 使用循环方式
func PopCountLoop(x uint64) int {
	Init()
	sum := 0
	for i := 0; i < 8; i++ {
		sum += int(pc[byte(x>>(i*8))])
	}
	return sum
}

// 使用移位算法重写PopCount函数，每次测试最右边的1bit，然后统计总数
func PopCountShift(x uint64) int {
	count := 0
	for i := 0; i < 64; i++ {
		if x&1 == 1 {
			count++
		}

		x >>= 1
	}
	return count
}
func PopCountEach8(x uint64) int {
	var res int
	for step := 7; step >= 0; step-- {
		res += int(pc[byte(x>>(8*step))])
	}
	return res
}

// 表达式 x&(x-1) 会将x的最低的一个非零的bit位清零，使用这个PopCount函数
func PopCountClearLowest(x uint64) int {
	count := 0
	for x != 0 {
		x = x & (x - 1)
		count++
	}
	return count
}
