package palindrome_number

import (
	"strconv"
	"strings"
)

// example x=121
// output true
// example x=-121
// output false
// cons -231 <= x <= 231 - 1
// 1234321
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	str := strconv.Itoa(x)
	split := strings.Split(str, "")
	// 計算總數
	arrLen := len(split)
	isnum := arrLen % 2
	// 偶數
	// 1221
	var tag []string
	if isnum == 0 {
		for k := range split {
			if split[k] != split[arrLen-k-1] {
				tag = append(tag, "1")
			}
		}
		// 1234321
	} else {
		mid := (arrLen - 1) / 2
		for k := range split {
			if k == mid {
				continue
			}
			if split[k] != split[arrLen-k-1] {
				tag = append(tag, "1")
			}
		}
	}
	if len(tag) < 1 {
		return true
	}
	return false
}
