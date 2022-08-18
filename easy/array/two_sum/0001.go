package two_sum

func twoSum(nums []int, target int) []int {
	// give init index map length
	index := make(map[int]int, len(nums))

	// 透過for迴圈 取得b的index
	for i, b := range nums {
		// 透過查詢map 取得a=target-b的index
		if j, ok := index[target-b]; ok {
			// 說明在i之前 ,存在 num[j]==a
			return []int{j, i}
		}
		// 把b和i的值 存入map
		index[b] = i
	}
	return nil
}
