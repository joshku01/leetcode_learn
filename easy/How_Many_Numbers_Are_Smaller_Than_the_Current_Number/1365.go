package How_Many_Numbers_Are_Smaller_Than_the_Current_Number

// Input: nums = [8,1,2,2,3]
// Output: [4,0,1,1,3]
func smallerNumbersThanCurrent(nums []int) []int {

	arrLen := len(nums)
	var outArr []int
	for i := 0; i < arrLen; i++ {
		var comp int
		for _, v := range nums {
			if nums[i] > v {
				comp++
			}
		}
		outArr = append(outArr, comp)
	}
	return outArr
}
