package Find_the_Town_Judge

import "fmt"

// Input: n = 2, trust = [[1,2]]
// Output: 2
func findJudge(n int, trust [][]int) int {
	if n == 1 {
		return 1
	}
	trusteeH := map[int]int{}
	trustSrcH := map[int]int{}

	for _, t := range trust {
		trustSrcH[t[0]]++
		trusteeH[t[1]]++
	}
	fmt.Println("_--->tree", trusteeH)
	fmt.Println("---->trs", trustSrcH)

	nM1 := n - 1
	for k := range trusteeH {
		if trusteeH[k] == nM1 && trustSrcH[k] == 0 {
			return k
		}
	}

	return -1
}
