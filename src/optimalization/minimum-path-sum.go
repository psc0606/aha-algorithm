package optimalization

// https://leetcode-cn.com/problems/minimum-path-sum/
// dp[i][j] = min(dp[i-1][j], dp[i][j+1]) + grid[i][j], 0 <= i < m, 0 <= j < n
// init condition:
// dp[0][0] = grid[0][0]
// dp[i][0] = dp[i][0] + grid[i][0], i >= 1
// dp[0][j] = dp[0][j-1] + grid[0][j], j >= 1
func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	// init condition
	// init first column
	dp[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	// init first row
	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[m-1][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
