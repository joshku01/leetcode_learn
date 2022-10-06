package Maximum_Product_of_Three_Numbers

import (
	"sort"
)

// Input: nums = [1,2,3]
// Output: 6
//----------------------------
// Input: nums = [1,2,3,4]
// Output: 24
// 題意:找出三個數字的乘積最大數
// 運用排序方式,有兩個結果 1.前面兩個數字+最後面一個 2.後面三個數字
func maximumProduct(nums []int) int {
	sort.Ints(nums)

	// 可能是負數狀況
	Max1 := nums[0] * nums[1] * nums[len(nums)-1]
	// 可能都是正數的狀況
	Max2 := nums[len(nums)-1] * nums[len(nums)-2] * nums[len(nums)-3]

	if Max1 > Max2 {
		return Max1
	}

	return Max2

}
