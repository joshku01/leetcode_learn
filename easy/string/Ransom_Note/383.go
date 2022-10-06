package Ransom_Note

import "fmt"

// Input: ransomNote = "a", magazine = "b"
// Output: false
// Input: ransomNote = "aa", magazine = "aab"
// Output: true
func canConstruct(ransomNote string, magazine string) bool {
	chars := make([]int, 26)
	for _, c := range magazine {
		i := int(c - 'a')
		chars[i]++
	}

	for _, c := range ransomNote {
		i := int(c - 'a')
		fmt.Printf("%c\n", i)
		chars[i]--
		if chars[i] < 0 {
			return false
		}
	}

	return true
}
