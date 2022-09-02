package Top_K_Frequent_Elements

import (
	"sort"
)

// Input: nums = [1,1,1,2,2,3], k = 2
// Output: [1,2]
func topKFrequent(nums []int, k int) []int {

	hMap := make(map[int]int)
	res := []int{}
	for _, v := range nums {
		hMap[v]++
	}
	n := map[int][]int{}
	var a []int
	for kk, v := range hMap {
		n[v] = append(n[v], kk)
	}
	for k := range n {
		a = append(a, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))

	for _, kkk := range a {
		for _, s := range n[kkk] {
			res = append(res, s)
		}
	}
	// 排序完成整理後的陣列,最後在依照需要的數量做切片
	res = res[:k]

	return res

}
