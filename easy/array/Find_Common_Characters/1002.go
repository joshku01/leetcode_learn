package Find_Common_Characters

import (
	"strings"
)

// Input: words = ["cool", "lock", "cook""]
// Output: ["c","l"]
func commonChars(words []string) []string {
	totalLen := len(words)
	var pass []string
	first := strings.Split(words[0], "")
	for _, v := range first {
		var markNum int
		for i := 1; i < totalLen; i++ {
			other := strings.Split(words[i], "")
			for k, c := range other {
				if v == c {
					markNum++
					RemoveIndex(other, k)
				}
			}
			expectNum := totalLen - 1
			if markNum >= expectNum {
				pass = append(pass, v)
			}
		}
	}
	inResult := make(map[string]bool)
	var finalResult []string
	for _, str := range pass {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			finalResult = append(finalResult, str)
		}
	}
	//res := unique(pass)
	return finalResult
}

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
