package Number_of_Arithmetic_Triplets

// Input: nums = [0,1,4,6,7,10], diff = 3
// Output: 2
// Explanation:
// (1, 2, 4) is an arithmetic triplet because both 7 - 4 == 3 and 4 - 1 == 3.
// (2, 4, 5) is an arithmetic triplet because both 10 - 7 == 3 and 7 - 4 == 3.
// 題意: 給定一個整數陣列 找出連續數字相差diff的數量
// step1 使用迴圈方式由數字大往下找相差的數字為diff 則繼續尋找,順利完成一個答案則++
// step2 印出總數ans 的數量,即為總數
// 時間複雜度 O(n^3)
func arithmeticTriplets(nums []int, diff int) int {

	//var output []int
	ans := 0
	for i := len(nums) - 1; i >= 2; i-- {
		for j := i - 1; j >= 1; j-- {
			if nums[i]-nums[j] > diff {
				break
			}
			if nums[i]-nums[j] != diff {
				continue
			}
			for k := j - 1; k >= 0; k-- {
				if nums[j]-nums[k] > diff {
					break
				}
				if nums[j]-nums[k] != diff {
					continue
				}
				ans++
			}
		}
	}
	return ans
}
