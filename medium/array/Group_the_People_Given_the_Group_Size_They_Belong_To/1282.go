package Group_the_People_Given_the_Group_Size_They_Belong_To

// Input: groupSizes = [3,3,3,3,3,1,3]
// Output: [[5],[0,1,2],[3,4,6]]
// Explanation:
// The first group is [5]. The size is 1, and groupSizes[5] = 1.
// The second group is [0,1,2]. The size is 3, and groupSizes[0] = groupSizes[1] = groupSizes[2] = 3.
// The third group is [3,4,6]. The size is 3, and groupSizes[3] = groupSizes[4] = groupSizes[6] = 3.
// Other possible solutions are [[2,1,6],[5],[0,4,3]] and [[5],[0,6,2],[4,3,1]].
// example2
// Input: groupSizes = [2,1,3,3,3,2]
// Output: [[1],[0,5],[2,3,4]]
func groupThePeople(groupSizes []int) [][]int {

	var groups [][]int
	gSizes := make(map[int][]int)
	// 運用map的方式建立對應,存入的長度如果=實際長度則清空值,並且清除值留給下一個用
	for idx, val := range groupSizes {
		gSizes[val] = append(gSizes[val], idx)

		// 如果陣列內的數量跟期望一樣,清除map陣列資料,同時寫入group結果內
		if len(gSizes[val]) == val {
			groups = append(groups, gSizes[val])
			gSizes[val] = nil
		}
	}
	for _, v := range gSizes {
		if v != nil {
			groups = append(groups, v)
		}
	}
	return groups
}
