package Richest_Customer_Wealth

// Input: accounts = [[1,2,3],[3,2,1]]
// Output: 6
func maximumWealth(accounts [][]int) int {

	var res []int
	for _, v := range accounts {
		resInt := 0
		for _, d := range v {
			resInt += d
		}
		res = append(res, resInt)
	}
	max := res[0]
	for _, val := range res {
		if val > max {
			max = val
		}
	}
	return max
}
