package Merge_Intervals

import (
	"math"
	"sort"
)

// Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
// Output: [[1,6],[8,10],[15,18]]
// Explanation: Since intervals [1,3] and [2,6] overlap, merge them into [1,6].
func merge(intervals [][]int) [][]int {
	// 先對二維陣列做排序arr[0]
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := [][]int{}
	for _, interval := range intervals {
		l := len(merged)
		if l == 0 || merged[l-1][1] < interval[0] {
			merged = append(merged, interval)
		} else {
			merged[l-1][1] = int(math.Max(float64(merged[l-1][1]), float64(interval[1])))
		}

	}

	return merged
}
