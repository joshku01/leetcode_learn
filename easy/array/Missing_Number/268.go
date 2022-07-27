package Missing_Number

// Input: nums = [3,0,1]
// Output: 2
func missingNumber(nums []int) int {
	ans := 0
	ans ^= len(nums)

	for i, v := range nums {
		ans ^= i ^ v
	}

	return ans
}
