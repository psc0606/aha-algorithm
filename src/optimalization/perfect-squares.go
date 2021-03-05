package optimalization

// https://leetcode-cn.com/problems/perfect-squares/
// see https://leetcode-cn.com/problems/coin-change/
//
// The first solution is dynamic programming.
// It's a little bit more like `coins change`, it's a Knapsack problem.
// Time: O(N*sqrt(N))
func numSquares(n int) int {
	squares := getSquares(n)
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i < n+1; i++ {
		dp[i] = n + 1 // n+1 out of range, means inf, or use math.MaxInt
		for _, s := range squares {
			if i-s >= 0 {
				dp[i] = min(dp[i], dp[i-s]+1)
			}
		}
	}
	return dp[n]
}

// get all possible squares
func getSquares(n int) []int {
	var ret []int
	i := 0
	for i*i <= n {
		ret = append(ret, i*i)
		i++
	}
	return ret
}

// We can also use theorem on the sum of foursquares.
