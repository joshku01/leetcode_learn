package Fibonacci_Number

// Input: n = 2
// Output: 1
func fib(n int) int {
	if n == 0 {
		return n
	}
	n1, n2, tmp := 0, 1, 0
	for i := 2; i <= n; i++ {
		tmp = n2
		n2 = n1 + n2
		n1 = tmp
	}
	return n2
}
