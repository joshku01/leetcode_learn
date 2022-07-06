package length_of_last_word

import "strings"

// "hello world"
func lengthOfLastWord1(s string) int {
	tri := strings.Trim(s, " ")
	spl := strings.Split(tri, " ")
	lenth := len(spl)
	last := spl[lenth-1]
	lastStrLen := len([]rune(last))

	return lastStrLen

}
