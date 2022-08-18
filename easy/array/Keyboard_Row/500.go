package Keyboard_Row

import (
	"strings"
)

// Input: words = ["Hello","Alaska","Dad","Peace"]
// Output: ["Alaska","Dad"]
func findWords(words []string) []string {
	top := []string{"q", "w", "e", "r", "t", "y", "u", "i", "o", "p"}
	mid := []string{"a", "s", "d", "f", "g", "h", "j", "k", "l"}
	last := []string{"z", "x", "c", "v", "b", "n", "m"}

	var output []string
	for _, v := range words {
		var countS1 int
		var countS2 int
		var countS3 int
		toLow := strings.ToLower(v)
		splitStr := strings.Split(toLow, "")
		lengS := len(splitStr)
		for _, s := range splitStr {
			for _, s1 := range top {
				if s1 != s {
					continue
				}
				if s1 == s {
					countS1++
				}
			}
			for _, s2 := range mid {
				if s2 != s {
					continue
				}
				if s2 == s {
					countS2++
				}
			}
			for _, s3 := range last {
				if s3 != s {
					continue
				}
				if s3 == s {
					countS3++
				}
			}
		}
		if lengS == countS1 {
			output = append(output, v)
		}
		if lengS == countS2 {
			output = append(output, v)
		}
		if lengS == countS3 {
			output = append(output, v)
		}
	}
	return output
}
