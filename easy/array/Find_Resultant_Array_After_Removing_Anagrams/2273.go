package Find_Resultant_Array_After_Removing_Anagrams

import (
	"fmt"
	"sort"
	"strings"
)

// Input: words = ["abba","baba","bbaa","cd","cd"]
// Output: ["abba","cd"]
func removeAnagrams(words []string) []string {

	var resliceArr []string
	for _, v := range words {
		element := strings.Split(v, "")
		sort.Strings(element)
		var restAtr string
		for _, v1 := range element {
			restAtr = restAtr + v1
		}
		resliceArr = append(resliceArr, restAtr)
	}
	fmt.Println("---->after sort", resliceArr)

	fmt.Println("---->", resliceArr)

	return []string{"abba", "cd"}

}

func removeElement(slice []string, i int) []string {
	copy(slice[i:], slice[i+1:])
	slice[len(slice)-1] = ""
	slice = slice[:len(slice)-1]
	return slice
}
