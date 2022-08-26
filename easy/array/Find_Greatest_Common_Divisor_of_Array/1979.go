package Find_Greatest_Common_Divisor_of_Array

// Input: nums = [2,5,6,9,10]
// Output: 2
// Explanation:
// The smallest number in nums is 2.
// The largest number in nums is 10.
// The greatest common divisor of 2 and 10 is 2.
// 題意: 給定一個陣列,找出陣列中最大數,最小數之間的最大公因數
func findGCD(nums []int) int {

	smallest := nums[0]
	largest := nums[0]
	for _, v := range nums {
		if v < smallest {
			smallest = v
		}
		if v > largest {
			largest = v
		}
	}
	// find the largest common divisor
	for i := largest; i > 0; i-- {
		if largest%i == 0 && smallest%i == 0 {
			return i
		}
	}
	return 1

}
