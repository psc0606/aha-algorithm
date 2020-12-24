package optimalization

// https://leetcode-cn.com/problems/climbing-stairs/
func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	a := 1
	b := 2
	for i := 3; i <= n; i++ {
		tmp := a
		a = b
		// tmp is last step is two.
		// a means last step is one.
		b = tmp + b
	}
	return b
}
