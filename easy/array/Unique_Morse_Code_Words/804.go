package Unique_Morse_Code_Words

import (
	"strings"
)

// Input: words = ["gin","zen","gig","msg"]
// Output: 2
func uniqueMorseRepresentations(words []string) int {
	englishMap := map[string]string{"a": ".-", "b": "-...", "c": "-.-.", "d": "-..", "e": ".", "f": "..-.", "g": "--.", "h": "....", "i": "..", "j": ".---", "k": "-.-", "l": ".-..", "m": "--", "n": "-.", "o": "---", "p": ".--.", "q": "--.-", "r": ".-.", "s": "...", "t": "-", "u": "..-", "v": "...-", "w": ".--", "x": "-..-", "y": "-.--", "z": "--.."}

	var output []string
	for _, v := range words {
		seperate := strings.Split(v, "")
		var temStr string
		for _, v1 := range seperate {
			str := englishMap[v1]
			temStr += str
		}
		output = append(output, temStr)
	}
	inResult := make(map[string]bool)
	var result []string
	for _, str := range output {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}
	total := len(result)

	return total
}
