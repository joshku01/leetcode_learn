package Binary_Gap

import (
	"strconv"
)

// Input: n = 8
// Output: 0
// Explanation: 8 in binary is "1000".
// There are not any adjacent pairs of 1's in the binary representation of 8, so we return 0.
// step1 先轉成2進位
// step2 如果沒有相鄰的兩個1 return 0
// step3 相鄰兩個一最遠距離
func binaryGap(n int) int {
	result := ""
	// 轉換2進位
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}

	arr := []int{}
	for i := 0; i < len(result); i++ {
		if result[i] == '1' {
			arr = append(arr, i)
		}
	}
	res := 0
	for i := len(arr) - 1; i > 0; i-- {
		if arr[i]-arr[i-1] > res {
			res = arr[i] - arr[i-1]
		}
	}
	return res

}
