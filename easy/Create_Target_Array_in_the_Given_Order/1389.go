package Create_Target_Array_in_the_Given_Order

import "fmt"

func isset(arr []int, index int) bool {
	return (len(arr) > index)
}

// Input: nums = [0,1,2,3,4], index = [0,1,2,2,1]
// Output: [0,4,1,3,2]
//
//nums       index     target
//0            0        [0]
//1            1        [0,1]
//2            2        [0,1,2]
//3            2        [0,1,3,2]
//4            1        [0,4,1,3,2]
func createTargetArray(nums []int, index []int) []int {
	var res []int

	big := 0
	for _, v := range index {
		if v > big {
			big = v
		}
	}

	for k, v := range index {
		// exist
		if isset(res, v) {
			fmt.Println("---->is not null", nums[k])
			res = append(res[:v], nums[v])
		}
		res = append(res, nums[k])
		//res[v] = nums[k]
		//fmt.Println("---->res[v]", res[v])
		//res[v] = nums[v]
	}

	fmt.Println("--->out", res)

	return res
}
