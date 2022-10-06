package Min_Cost_Climbing_Stairs

// Input: cost = [1,100,1,1,1,100,1,1,100,1]
// Output: 6
// Explanation: You will start at index 0.
//- Pay 1 and climb two steps to reach index 2.
//- Pay 1 and climb two steps to reach index 4.
//- Pay 1 and climb two steps to reach index 6.
//- Pay 1 and climb one step to reach index 7.
//- Pay 1 and climb two steps to reach index 9.
//- Pay 1 and climb one step to reach the top.
//The total cost is 6.
func minCostClimbingStairs(cost []int) int {
	// 初始化 並且前面先加2個0
	cost = append([]int{0, 0}, cost...)
	cost = append(cost, 0)

	dp := make([]int, len(cost))
	for i := 2; i < len(cost); i++ {
		dp[i] = min(dp[i-2], dp[i-1]) + cost[i]
	}
	return dp[len(dp)-1]

}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
