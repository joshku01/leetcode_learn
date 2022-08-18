package Climbing_Stairs

/* 題目解釋
一次可爬一階或二階，假如給訂個正整數n階,爬到n階有幾種可能？
*/
/**
step1 先找出題目規律性
n=1  => 1
n=2  => 2
n=3  => 3
n=4  => 5  (n-1)+(n-2)
從這裡發現 規律為前兩個數字相加,跟fib係數一樣~
*/
// Input: n = 2
// Output: 2
// classic Fibonacci Sequence 類似fib係數問題
// 可以使用遞迴解法,但是速度很慢 不建議(Leetcode runtime直接Time Limit Exceeded暴掉...)
// 改用DP方式運行
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
