package Climbing_Stairs

// Input: n = 2
// Output: 2
// clasic Finonacci Sequence
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	result := 0
	pre := 1
	next := 2
	for i := 2; i < n; i++ {
		result = pre + next
		pre = next
		next = result
	}
	return result
}
