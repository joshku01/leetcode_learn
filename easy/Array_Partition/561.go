package Array_Partition

import "sort"

//Input: nums = [1,4,3,2]
//Output: 4
func arrayPairSum(nums []int) int {

	sort.Ints(nums)
	arrLen := len(nums)
	var output int
	for i := 0; i < arrLen; i += 2 {
		output = output + nums[i]
	}

	return output

}
