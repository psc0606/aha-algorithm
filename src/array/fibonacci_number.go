package array

// Recursively
// Time: O(2^n)
// Space: O()
func fibonacciNumberRecursively(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonacciNumberRecursively(n-1) + fibonacciNumberRecursively(n-2)
}

// Dynamic programing
// Time: O(n)
// Space: O(n)
func fibonacciNumberDp(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i < n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// Dynamic programing
// Time: O(n)
// Space: O(1)
func fibonacciNumberDpOptimizeSpace(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	a := 0
	b := 1
	for i := 2; i < n+1; i++ {
		tmp := b
		b += a
		a = tmp
	}
	return b
}
