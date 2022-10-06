package Move_Zeroes

// Input: nums = [0,1,0,3,12]
// Output: [1,3,12,0,0]
// Input: nums = [0]
// Output: [0]
func moveZeroes(nums []int) {
	slow := 0
	for current := range nums {
		if nums[current] != 0 {
			nums[slow], nums[current] = nums[current], nums[slow]
			slow++
		}
	}
}
