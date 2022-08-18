package plun_one

import (
	"fmt"
	"strconv"
	"strings"
)

//example digits=[1,2,3]
//return 123+1=124
// answer is [1,2,4]
func plusOne(digits []int) []int {
	arrSum := ""
	for i := 0; i < len(digits); i++ {
		str := strconv.Itoa(digits[i])
		arrSum = arrSum + str
	}
	fmt.Println("---step1", arrSum)
	strToInt, _ := strconv.Atoi(arrSum)
	fmt.Println("--->to str", strToInt)
	fin := strToInt + 1

	toStr := strconv.Itoa(fin)
	sp := strings.Split(toStr, "")

	var finArr []int
	for k := range sp {
		strToInt, _ := strconv.Atoi(sp[k])
		finArr = append(finArr, strToInt)
	}
	fmt.Println("----->ans", finArr)

	return finArr
}
