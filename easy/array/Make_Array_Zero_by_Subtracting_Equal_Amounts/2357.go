package Make_Array_Zero_by_Subtracting_Equal_Amounts

// Input: nums = [1,5,0,3,5]
// Output: 3
// Explanation:
// In the first operation, choose x = 1. Now, nums = [0,4,0,2,4].
// In the second operation, choose x = 2. Now, nums = [0,2,0,0,2].
// In the third operation, choose x = 2. Now, nums = [0,0,0,0,0].
// 題意:給一個陣列,找出最小數後減去每一次中陣列的最小數直到所以數為0為止
func minimumOperations(nums []int) int {
	hMap := map[int]bool{}
	for _, v := range nums {
		if v > 0 {
			if _, ok := hMap[v]; !ok {
				hMap[v] = true
			}
		}
	}
	return len(hMap)
}
