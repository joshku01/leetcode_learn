package Two_Out_of_Three

// Input: nums1 = [1,1,3,2], nums2 = [2,3], nums3 = [3]
// Output: [3,2]
// 題意:給定三個陣列,找出元素同時包含至少在任意兩個陣列當中,並返回元素陣列
func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	result := []int{}
	output := []int{}

	for _, n1 := range nums1 {
		for _, n2 := range nums2 {
			if n1 == n2 {
				result = append(result, n1)
			}
		}
		for _, n3 := range nums3 {
			if n1 == n3 {
				result = append(result, n1)
			}
		}
	}
	for _, n2 := range nums2 {
		for _, n3 := range nums3 {
			if n2 == n3 {
				result = append(result, n2)
			}
		}
	}
	hmap := map[int]bool{}

	for _, v := range result {
		if _, ok := hmap[v]; !ok {
			hmap[v] = true
		}
	}
	for k := range hmap {
		output = append(output, k)
	}

	return output
}
