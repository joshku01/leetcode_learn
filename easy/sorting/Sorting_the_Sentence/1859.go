package Sorting_the_Sentence

import (
	"sort"
	"strings"
)

// Input: s = "is2 sentence4 This1 a3"
// Output: "This is a sentence"
func sortSentence(s string) string {

	var sortArr []string
	strArr := strings.Split(s, " ")
	for _, v := range strArr {
		lastElement := v[len(v)-1:]
		deleteLast := v[:len(v)-1]
		tempStr := lastElement + deleteLast
		sortArr = append(sortArr, tempStr)
		sort.Strings(sortArr)
	}
	var result []string
	for _, v := range sortArr {
		temp := v[1:]
		result = append(result, temp)
	}
	var restStr string
	for _, v := range result {
		restStr = restStr + " " + v
	}
	trimString := strings.Trim(restStr, " ")

	return trimString

}
