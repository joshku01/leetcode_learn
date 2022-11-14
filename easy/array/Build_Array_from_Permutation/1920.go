package Build_Array_from_Permutation

import "fmt"

// Input: nums = [0,2,1,5,3,4]
// Output: [0,1,2,4,5,3]
func buildArray(nums []int) []int {
	var outArr []int
	for k := range nums {
		outArr = append(outArr, nums[nums[k]])
	}
	a := 7
	b := 2
	o := a / b
	fmt.Println("--->ans", o)

	return outArr
}
