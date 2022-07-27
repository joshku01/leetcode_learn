package Maximum_Number_of_Words_Found_in_Sentences

import "strings"

// Input: sentences = ["alice and bob love leetcode", "i think so too", "this is great thanks very much"]
// Output: 6
func mostWordsFound(sentences []string) int {
	var initLen = 0
	for _, v := range sentences {
		spl := strings.Split(v, " ")
		strLen := len(spl)
		if strLen > initLen {
			initLen = strLen
		}
	}
	return initLen
}
