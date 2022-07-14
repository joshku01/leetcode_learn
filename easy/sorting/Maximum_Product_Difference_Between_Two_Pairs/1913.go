package Maximum_Product_Difference_Between_Two_Pairs

import (
	"sort"
)

//Input: nums = [5,6,2,7,4]
//Output: 34
func maxProductDifference(nums []int) int {

	sort.Ints(nums)
	last := nums[len(nums)-1:]
	last2 := nums[len(nums)-2 : len(nums)-1]

	combine1 := last[0] * last2[0]
	combine2 := nums[0] * nums[1]

	result := combine1 - combine2

	return result
}
