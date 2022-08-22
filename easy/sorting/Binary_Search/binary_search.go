package Binary_Search

// input [1,4,5,7,2,23,7,45,23,33]
// param target=5
func search(nums []int, target int) int {
	r := -1
	head := 0
	tail := len(nums) - 1
	middle := (head + tail) / 2

	for head <= tail {
		middle = (head + tail) / 2
		if nums[middle] == target {
			return middle
		} else if nums[middle] > target {
			// 因為middle不是要的答案,自然就必須排除在外
			tail = middle - 1
		} else {
			head = middle + 1
		}
	}
	return r

}
