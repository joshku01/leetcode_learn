package reverse_interger

import (
	"math"
)

func reverse(x int) int {
	// example x=123
	sign := 1

	// 處理負數
	if x < 0 {
		sign = -1
		x = -1 * x
	}
	res := 0
	for x > 0 {
		// 取x的尾數
		temp := x % 10
		// put in the res
		res = res*10 + temp
		// x 去除末尾
		x = x / 10
	}
	res = sign * res
	if res > math.MaxInt32 || res < math.MinInt32 {
		res = 0
	}
	return res

}

