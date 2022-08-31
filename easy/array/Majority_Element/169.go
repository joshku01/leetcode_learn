package Majority_Element

// Input: nums = [3,2,3]
// Output: 3
func majorityElement(nums []int) int {
	hMap := map[int]int{}
	for _, v := range nums {
		if _, ok := hMap[v]; !ok {
			hMap[v] = 1
		} else {
			hMap[v]++
		}
	}
	biggest := 0
	for _, v := range hMap {
		if v > biggest {
			biggest = v
		}
	}
	for k, v := range hMap {
		if v == biggest {
			return k
		}
	}
	return 0
}
