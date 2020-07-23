package squares_of_a_sorted_array

import (
	"math"
	"sort"
)

func sortedSquares(A []int) []int {
	var b []int
	var square int
	if len(A) == 0 {
		return []int{}
	}
	for i := 0; i < len(A); i++ {
		if A[i] < 0 {
			positive := int(math.Abs(float64(A[i])))
			square = positive * positive
			b = append(b, square)
		} else {
			square = A[i] * A[i]
			b = append(b, square)
		}
	}
	sort.Ints(b)

	return b
}
