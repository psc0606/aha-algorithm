package math

// https://leetcode-cn.com/problems/powx-n/

// Time: O(n)
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	positive := true
	if n < 0 {
		n = -n
		positive = false
	}
	var ret = 1.0
	for i := 0; i < n; i++ {
		ret *= x
	}
	if !positive {
		ret = 1 / ret
	}
	return ret
}

// Time: O(logn)
// 2^11 = 2^8 * 2^2 * 2^1
func myPow2(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	positive := true
	if n < 0 {
		n = -n
		positive = false
	}
	var ret, cal = 1.0, x
	k := 1
	for n != 0 {
		k *= 2
		if k <= n {
			cal *= cal
		} else {
			n -= k / 2
			k = 1
			ret *= cal
			cal = x
		}
	}
	if !positive {
		ret = 1.0 / ret
	}
	return ret
}

// A beautiful solution
// @see https://leetcode-cn.com/problems/powx-n/comments/10689
// Time: O(logn)
func myPow3(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	ret := 1.0
	for i := n; i != 0; i /= 2 {
		if i%2 != 0 {
			ret *= x
		}
		x *= x
	}
	if n < 0 {
		ret = 1 / ret
	}
	return ret
}
