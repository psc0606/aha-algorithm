package optimalization

// https://leetcode-cn.com/problems/container-with-most-water/
// dynamic programming, but it has a better solution.
// Time: O(n^2)
// Space: O(n)
// input required: 2 <= n <= 3 * 104
func maxArea(height []int) int {
	n := len(height)
	dp := make([]int, n)
	for i := 1; i < n; i++ {
		maximum := dp[i-1]
		for j := i - 1; j >= 0; j-- {
			area := (i - j) * min(height[i], height[j])
			if maximum < area {
				maximum = area
			}
		}
		dp[i] = maximum
	}
	return dp[n-1]
}

// Solution2: double pointer, then dynamic programming
// Time: O(n)
func maxArea2(height []int) int {
	n := len(height)
	i, j := 0, n-1
	maximum := 0
	for i < j {
		maximum = max(maximum, (j-i)*min(height[i], height[j]))
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return maximum
}
