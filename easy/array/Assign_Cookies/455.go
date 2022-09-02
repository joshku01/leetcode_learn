package Assign_Cookies

import "sort"

// Input: g = [1,2], s = [1,2,3]
// Output: 2
// Explanation: You have 2 children and 3 cookies. The greed factors of 2 children are 1, 2.
// You have 3 cookies and their sizes are big enough to gratify all of the children,
// You need to output 2.
// 題意:兩個整數陣列,滿足陣列g數量的大小符合陣列s的大小數量
func findContentChildren(g []int, s []int) int {
	// 先對陣列做排序
	sort.Ints(g)
	sort.Ints(s)
	output := 0
	i, j := 0, 0
	// 依序比對所有陣列的值,滿足s[i]>g[i]的話,加入總數
	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			i++
			j++
			output++
		} else {
			j++
		}
	}
	return output

}
