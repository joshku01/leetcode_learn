package Valid_Palindrome

import (
	"regexp"
	"strings"
)

// Input: s = "A man, a plan, a canal: Panama"
// Output: true
// Explanation: "amanaplanacanalpanama" is a palindrome.
func isPalindrome(s string) bool {
	// 作法:先做正規表達式->移除非英文字母的字元, 再做轉小寫動作
	newS := regexp.MustCompile(`[^a-z0-9]+`).ReplaceAllString(strings.ToLower(s), "")

	// 從頭跟尾做比較是否為回文
	i := 0
	j := len(newS) - 1
	for i < j {
		if newS[i] == newS[j] {
			i++
			j--
		} else {
			return false
		}
	}
	return true

}
