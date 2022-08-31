package Monotonic_Array

// Input: nums = [1,3,2,4]
// Output: true
// Input: nums = [6,6,5,4]
// Output: true
func isMonotonic(nums []int) bool {
	if len(nums) <= 2 {
		return true
	}
	inc := true
	ch := false
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			continue
		}
		// increase 遞增判斷
		if nums[i-1] < nums[i] {
			if !ch {
				ch = true
				inc = true
			} else {
				if !inc {
					return false
				}
			}
		} else {
			// decrease遞減判斷
			if !ch {
				ch = true
				inc = false
			} else {
				if inc {
					return false
				}
			}
		}
	}
	return true

}
