package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(sum(1, 2, 3))
	values := []int{1, 2, 3, 4}
	fmt.Println(sum(values...))
	max, min, err := maxMin()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("max:", max, "min:", min)
}

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func maxMin(nums ...int) (max, min int, err error) {
	// 如果参数个数为0，则返回错误
	err = fmt.Errorf("err: 参数个数为0")
	if len(nums) == 0 {
		return 0, 0, err
	}

	// 初始化最小值为最大整数
	min = math.MaxInt
	for _, num := range nums {
		if num > max {
			max = num
		}

		if num < min {
			min = num
		}
	}
	return max, min, nil
}
