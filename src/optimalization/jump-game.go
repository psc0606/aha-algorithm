package optimalization

// https://leetcode-cn.com/problems/jump-game/
//
// This problem can be converted to a optimization, if we find the longest position,
// it will be easy to figure out whether it can reach the end or not.
//
//         |- max(dp[i-1], i+nums[i]),  when dp[i-1] >= i, means dp[i-1] can reach i.
// dp[i] = |
//         |- dp[i-1]
// `i` is the position
// Time: O(n)
// Space: O(n), we can also use only one variable to O(1)
func canJump(nums []int) bool {
	n := len(nums)
	if n < 1 {
		return false
	}
	dp := make([]int, n)
	dp[0] = nums[0]
	for i := 1; i < n; i++ {
		if dp[i-1] >= i {
			dp[i] = max(dp[i-1], i+nums[i])
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[n-1] >= n-1
}

func canJumpWithOptimize(nums []int) bool {
	n := len(nums)
	if n < 1 {
		return false
	}
	maxDis := nums[0]
	for i := 1; i < n; i++ {
		if maxDis >= i {
			maxDis = max(maxDis, i+nums[i])
		}
	}
	return maxDis >= n-1
}
