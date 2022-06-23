package concatenation_of_array

// Input: nums = [1,2,1]
// Output: [1,2,1,1,2,1]
func getConcatenation(nums []int) []int {
	var out []int

	for i := 0; i < 2; i++ {
		for _, v := range nums {
			out = append(out, v)
		}
	}
	return out
}
