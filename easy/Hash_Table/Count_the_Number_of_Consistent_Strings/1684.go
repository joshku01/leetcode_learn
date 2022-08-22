package Count_the_Number_of_Consistent_Strings

import "strings"

// Input: allowed = "ab", words = ["ad","bd","aaab","baa","badab"]
// Output: 2
// Explanation: Strings "aaab" and "baa" are consistent since they only contain characters 'a' and 'b'.
func countConsistentStrings(allowed string, words []string) int {
	count := 0
	for _, v := range words {
		toStr := strings.Split(v, "")
		allow := strings.Split(allowed, "")
		wordLen := len(toStr)
		countNum := 0
		for _, v2 := range toStr {
			status := stringContain(allow, v2)
			if status {
				countNum++
			}
		}
		if countNum == wordLen {
			count++
		}
	}
	return count

}

func stringContain(arr []string, ele string) bool {
	for _, v := range arr {
		if v == ele {
			return true
		}
	}
	return false
}
