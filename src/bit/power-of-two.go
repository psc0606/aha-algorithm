package bit

// https://leetcode-cn.com/problems/power-of-two/
func isPowerOfTwo(n int) bool {
	if n == 0 {
		return true
	}
	return n&(n-1) == 0
}
