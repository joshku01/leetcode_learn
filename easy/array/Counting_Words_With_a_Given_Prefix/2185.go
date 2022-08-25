package Counting_Words_With_a_Given_Prefix

// Input: words = ["pay","attention","practice","attend"], pref = "at"
// Output: 2
// Explanation: The 2 strings that contain "at" as a prefix are: "attention" and "attend".
func prefixCount(words []string, pref string) int {
	ans, n := 0, len(pref)
	for _, word := range words {
		if len(word) >= n && word[:n] == pref {
			ans++
		}
	}
	return ans
}
