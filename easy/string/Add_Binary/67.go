package Add_Binary

import (
	"math"
)

// Input: a = "11", b = "1"
// Output: "100"
func addBinary(a string, b string) string {
	ia, ib := len(a)-1, len(b)-1
	carry := 0
	result := ""

	for (ia >= 0) || (ib >= 0) || (carry > 0) {
		if ia >= 0 {
			carry += int(a[ia] - '0')
			ia--
		}

		if ib >= 0 {
			carry += int(b[ib] - '0')
			ib--
		}

		if carry%2 == 0 {
			result = "0" + result
		} else {
			result = "1" + result
		}

		carry /= 2
	}

	return result

}

// convertToTen 轉換十進位
func convertToTen(str string) float64 {
	count := float64(0)
	total := float64(0)
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] != byte('0') {
			total += math.Pow(2, count)
		}
		count++
	}
	return total
}
