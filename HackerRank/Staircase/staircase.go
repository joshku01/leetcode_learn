package Staircase

import "fmt"

// n:=4
// output
//   #
//  ##
// ###
//####
func staircase(n int32) {
	// Write your code here

	for i := 1; i <= int(n); i++ {
		space := int(n) - i
		simbol := i
		outputSpace := ""
		outputSimbol := ""
		for i := 0; i < space; i++ {
			outputSpace += " "
		}
		for i := 0; i < simbol; i++ {
			outputSimbol += "#"
		}
		output := fmt.Sprintf("%s%s", outputSpace, outputSimbol)
		fmt.Println(output)

	}

}
