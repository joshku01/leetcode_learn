package Minimum_Absolute_Difference

import (
	"fmt"
	"sort"
)

//Input: arr = [4,2,1,3]
//Output: [[1,2],[2,3],[3,4]]
func minimumAbsDifference(arr []int) [][]int {

	s := []int{1, 2, 3}
	arLen := len(s)

	for i := 0; i < arLen; i += 2 {
		fmt.Println("----->", s[i])
	}

	//step1 sort to [1,2,3,4]
	sort.Ints(arr)
	arrLen := len(arr)

	initNum := arr[1] - arr[0]
	markNum := []int{0}
	var outputArr [][]int
	for i := 1; i < arrLen-1; i++ {
		tempMin := arr[i+1] - arr[i]
		if tempMin < initNum {
			initNum = tempMin
			markNum = append(markNum, i)
			markNum = markNum[len(markNum)-1:]
		} else if tempMin == initNum {
			initNum = tempMin
			markNum = append(markNum, i)
		}
	}
	for _, v := range markNum {
		tempArrOut := []int{arr[v], arr[v+1]}
		outputArr = append(outputArr, tempArrOut)
	}
	return outputArr
}
