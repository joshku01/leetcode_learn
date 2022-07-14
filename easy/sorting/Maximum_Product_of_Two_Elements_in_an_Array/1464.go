package Maximum_Product_of_Two_Elements_in_an_Array

import "sort"

//Input: nums = [3,4,5,2]
//Output: 12
func maxProduct(nums []int) int {
	sort.Ints(nums)
	last := nums[len(nums)-1:]
	last2 := nums[len(nums)-2 : len(nums)-1]
	ele1 := last[0] - 1
	ele2 := last2[0] - 1
	output := ele1 * ele2

	return output
}
