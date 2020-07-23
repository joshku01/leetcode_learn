package sort_array_by_parity

import (
	"fmt"
	"math"
)

/*
input: 一個沒有負數的數字陣列
output: 回傳前面是偶數後面是奇數的陣列
example: [1,2,5,6]
limit exaple  [1,2,1]
*/
func sortArrayByParity(A []int) []int {
	var b []int
	c := -1
	fmt.Println(int64(math.Abs(float64(c))))
	if len(A) == 0 {
		return []int{}
	}

	for i := 0; i < len(A); i++ {
		if A[i]%2 == 0 {
			b = append([]int{A[i]}, b...)
		} else {
			b = append(b, A[i])
		}
	}
	return b
}
