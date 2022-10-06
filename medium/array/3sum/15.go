package _sum

import (
	"sort"
)

// Input: nums = [-1,0,1,2,-1,-4]
// Output: [[-1,-1,2],[-1,0,1]]
// Explanation:
// nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
// nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
// nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
// The distinct triplets are [-1,0,1] and [-1,-1,2].
// Notice that the order of the output and the order of the triplets does not matter.
// 題目思考方向: 找出 a+b = -c 的答案
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}

	for t := 0; t < len(nums); t++ {
		// t > 0的話不執行
		if t > 0 && nums[t] == nums[t-1] {
			continue
		}
		i := t + 1
		j := len(nums) - 1
		for i < j {
			sum := nums[i] + nums[j] + nums[t]
			if sum == 0 {
				ans := []int{nums[t], nums[i], nums[j]}
				result = append(result, ans)
				i++
				for i < j && nums[i] == nums[i-1] {
					i++
				}
			} else if sum < 0 {
				// 總和<0 往右移動
				i++
			} else {
				// 總和>0 往左移動
				j--
			}
		}
	}
	return result
}
