package Sum_of_Unique_Elements

// Input: nums = [1,2,3,2]
// Output: 4
// Explanation: The unique elements are [1,3], and the sum is 4.
func sumOfUnique(nums []int) int {
	s := map[int]int{}
	r := 0
	for _, v := range nums {
		if _, ok := s[v]; ok {
			s[v] = 0
		} else {
			s[v]++
		}
	}
	for _, v := range nums {
		if s[v] == 1 {
			r += v
		}
	}
	return r
}
