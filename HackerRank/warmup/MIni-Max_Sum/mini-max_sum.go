package MIni_Max_Sum

import "fmt"

// [1,2,3,4,5]
// 思考方向:先將所以值加總起來,在迴圈過程中找出最大值及最小值,扣掉後就分別是要的數列
func miniMaxSum(arr []int32) {
	// Write your code here
	min := int64(int64(arr[0]))
	max := int64(int64(arr[0]))
	all := int64(0)

	for i := 0; i < len(arr); i++ {
		all += int64(arr[i])
		if min < int64(arr[i]) {
			min = int64(arr[i])
		}
		if max > int64(arr[i]) {
			max = int64(arr[i])
		}
	}
	fmt.Printf("%d\n", all-min)
	fmt.Printf("%d", all-max)
}
