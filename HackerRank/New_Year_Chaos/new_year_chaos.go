package New_Year_Chaos

import "fmt"

// q= [1,2,3,5,4,6,7,8]
// output print 1
// person 5 should bribes person 4 , only 1 bribes
func minimumBribes(q []int32) {
	// Write your code here
	//example [2,1,5,3,4]
	//output 3
	count := int32(0)
	for i := int32(len(q)) - 1; i >= 0; i-- {
		if q[i]-(i+1) > 2 {
			fmt.Println("Too chaotic")
			return
		}
		for j := max(0, q[i]-2); j < i; j++ {
			if q[j] > q[i] {
				count++
			}
		}
	}
	fmt.Println(count)

}

func max(a, b int32) int32 {
	if a >= b {
		return a
	}
	return b
}
