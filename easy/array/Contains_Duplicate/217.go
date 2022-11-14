package Contains_Duplicate

// Input: nums = [1,2,3,1]
// Output: true
func containsDuplicate(nums []int) bool {

	mapResult := make(map[int]bool)

	for _, str := range nums {
		if _, ok := mapResult[str]; !ok {
			mapResult[str] = true
		} else {
			return true
		}
	}
	return false

}
