package Minimum_Sum_of_Four_Digit_Number_After_Splitting_Digits

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//Input: num = 2932
//Output: 52
func minimumSum(num int) int {
	numStr := strconv.Itoa(num)

	split := strings.Split(numStr, "")
	fmt.Println("---->sp", len(split))
	var intnum []int
	for i := 0; i < len(split); i++ {
		data, _ := strconv.Atoi(split[i])
		intnum[i] = data
	}
	sort.Ints(intnum)
	fmt.Println("---->", intnum)

	d1 := intnum[0]
	d2 := intnum[1]
	d3 := intnum[2]
	d4 := intnum[3]
	res := (d1+d2)*10 + d3 + d4
	return res
}
