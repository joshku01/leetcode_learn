package Longest_Substring_Without_Repeating_Characters

import "math"

// Input: s = "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.
func lengthOfLongestSubstring(s string) int {

	// 使用hashmap來解題
	hmap := map[byte]int{}
	ans := 0
	left := 0
	for right := 0; right < len(s); right++ {
		rIndex, rExist := hmap[s[right]]

		if rExist && rIndex >= left {
			left = rIndex + 1
		} else {
			ans = int(math.Max(float64(ans), float64(right-left+1)))
		}
		hmap[s[right]] = right
	}

	return ans
}
