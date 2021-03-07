package array

// https://leetcode-cn.com/problems/plus-one/
// Time: O(n)
// But we can still optimize it.
func plusOne(digits []int) []int {
	n := len(digits)
	digits[n-1] += 1
	c := 0
	for i := len(digits) - 1; i >= 0; i-- {
		v := digits[i] + c
		digits[i] = v % 10
		c = v / 10
	}
	if c != 0 {
		return append([]int{1}, digits...)
	}
	return digits
}

// https://leetcode-cn.com/problems/plus-one/solution/java-shu-xue-jie-ti-by-yhhzw/
// This can avoid useless loop.
func plusOne2(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		digits[i]++
		digits[i] %= 10
		if digits[i] != 0 { // (9 + 0) % 10 == 0
			return digits
		}
	}
	digits = make([]int, n+1)
	digits[0] = 1
	return digits
}
