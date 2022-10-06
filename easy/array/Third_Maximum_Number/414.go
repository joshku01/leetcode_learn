package Third_Maximum_Number

import (
	"sort"
)

// Input: nums = [1,1,-2147483648]
// Output: 1
// Explanation:
// The first distinct maximum is 3.
// The second distinct maximum is 2.
// The third distinct maximum is 1.
func thirdMax(nums []int) int {
	sort.Slice(nums, func(a, b int) bool {
		return nums[a] > nums[b]
	})

	count := 0
	if len(nums) < 3 {
		return nums[0]
	}

	// 找到不同數第三個大的為止 return num[i]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			count++
			if count == 2 {
				return nums[i]
			}
		}
	}

	// 找不到return最大值
	return nums[0]
}
