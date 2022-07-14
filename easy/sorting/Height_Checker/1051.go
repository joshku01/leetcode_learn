package Height_Checker

import (
	"sort"
)

//Input: heights = [1,1,4,2,1,3]
//Output: 3
func heightChecker(heights []int) int {

	var origin []int
	origin = heights

	for _, v := range heights {
		origin = append(origin, v)
	}

	sort.Ints(heights)
	lenArr := len(heights)
	var output int

	for i := 0; i < lenArr; i++ {
		if origin[i] != heights[i] {
			output = output + 1
		}
	}
	return output
}
