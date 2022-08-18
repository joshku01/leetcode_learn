package Fibonacci_Number

// Input: n = 2
// Output: 1
// 解題思路
// f(0) 0
// f(1) 1
// f(2) 1
// f(3) 2
// f(4) 3
// f(i) f(i-1)+f(i-2)
// 從最一開始值 f(2)開始 f(0)=0 f(1)=1
// 使用dynamic programming 解題(效能好)
func fib(n int) int {
	if n == 0 {
		return n
	}
	// n1=前兩個值(f(i)-2)   n2=前兩個值(f(i)-1) temp=暫時存取值
	n1, n2, tmp := 0, 1, 0
	for i := 2; i <= n; i++ {
		tmp = n2
		n2 = n1 + n2
		n1 = tmp
	}

	return n2
}
