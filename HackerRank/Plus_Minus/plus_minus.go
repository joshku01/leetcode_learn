package Plus_Minus

import (
	"fmt"
	"strconv"
)

// arr=[1,1,,0,-1,-1]
// return 0.400000,0.400000,0.200000
func plusMinus(arr []int32) {
	// Write your code here
	output := make([]int, 3)
	totalLen := len(arr)
	for _, v := range arr {
		if v > 0 {
			output[0] = output[0] + 1
		}
		if v < 0 {
			output[1] = output[1] + 1
		}
		if v == 0 {
			output[2] = output[2] + 1
		}
	}

	for _, v := range output {
		num, _ := strconv.ParseFloat(fmt.Sprintf("%.6f", float64(v)/float64(totalLen)), 64)
		fmt.Println(num)
	}

}
