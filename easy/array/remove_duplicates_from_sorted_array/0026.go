package remove_duplicates_from_sorted_array

import "fmt"

func removeDuplicates(a []int) int {
	// example nums = [0,0,1,1,2,2,3,3,4] return length=5
	left, right, size := 0, 1, len(a)
	for ; right < size; right++ {
		if a[left] == a[right] {
			continue
		}
		left++
		a[left], a[right] = a[right], a[left]
	}
	fmt.Println("-------------")
	return left + 1
}

func removeDuplicatesTwo(nums []int) int {
	//example nums = [0,0,1,1,2]
	if len(nums) <= 1 {
		return len(nums)
	}
	tail := 1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] != nums[i] {
			nums[tail] = nums[i]
			tail++
		}
	}
	return tail
}
