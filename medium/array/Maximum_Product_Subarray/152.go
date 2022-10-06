package Maximum_Product_Subarray

import "math"

// Input: nums = [2,3,-2,4]
// Output: 6
// Explanation: [2,3] has the largest product 6.
// 題意:找出一個最大的連續非負數
func maxProduct(nums []int) int {

	if len(nums) == 0 {
		return 0
	}
	max, min, result := 1, 1, math.MinInt32
	for i := 0; i < len(nums); i++ {
		max, min = getMax(nums[i], max*nums[i], min*nums[i]), getMin(nums[i], max*nums[i], min*nums[i])
		if max > result {
			result = max
		}
	}
	return result
}

func getMax(m, n, k int) int {
	if m >= n && m >= k {
		return m
	}
	if n >= m && n >= k {
		return n
	}
	return k
}

func getMin(m, n, k int) int {
	if m <= n && m <= k {
		return m
	}
	if n <= m && n <= k {
		return n
	}
	return k
}
