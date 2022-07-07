package Decompress_Run_Length_Encoded_List

// Input: nums = [1,2,3,4]
// Output: [2,4,4,4]
// [2]+[4,4,4]
// result [2,4,4,4]
// 2 <= nums.length <= 100
// nums.length % 2 == 0
func decompressRLElist(nums []int) []int {
	leng := len(nums)
	resultArr := []int{}
	for i := 0; i < leng; i += 2 {
		num := nums[i]
		target := nums[i+1]
		for j := 1; j <= num; j++ {
			resultArr = append(resultArr, target)
		}
	}

	return resultArr
}
