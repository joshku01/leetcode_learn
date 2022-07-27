package Shuffle_String

import (
	"fmt"
	"strings"
)

// Input: s = "codeleet", indices = [4,5,6,7,0,2,1,3]
// Output: "leetcode"
// c o d e l e e t
// 4 5 6 7 0 2 1 3
func restoreString(s string, indices []int) string {

	mapData := map[string]int{}
	maxLen := len(indices)
	strSeperate := strings.Split(s, "")
	for i := 0; i < maxLen; i++ {
		mapData[strSeperate[i]] = indices[i]
		fmt.Println("---->data is ", mapData[strSeperate[i]])
	}
	fmt.Println("----.all", mapData)
	//fmt.Println("--->str daa", strSeperate)
	//fmt.Println("--->len", maxLen)
	//fmt.Println("---->init data", mapData)
	return "123"
}
