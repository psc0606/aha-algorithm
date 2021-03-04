package optimalization

// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/
// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/solution/zui-jia-mai-mai-gu-piao-shi-ji-han-leng-dong-qi-4/484682
// A better dynamic programming solution than office. Multi-state
// Time: O(n)
// Space: O(n)
// We can do some optimizations to reduce space cost.
func maxProfitWithCoolDown(prices []int) int {
	n := len(prices)
	if n < 1 {
		return 0
	}
	state := 5
	dp := make([][]int, state)
	for i := 0; i < state; i++ {
		dp[i] = make([]int, n)
	}
	// dp[i][0] means maximum profit, buy
	// dp[i][1] means maximum profit, sold
	// dp[i][2] means maximum profit, hold
	// dp[i][3] means maximum profit, in cool down
	// dp[i][4] means maximum profit, no stock(can buy but not buy)
	for i := 1; i < n; i++ {
		dp[0][i] = max(dp[3][i-1], dp[4][i-1])
		// actually, dp[1][i] equals dp[2][1], so we can merge both into one.
		dp[1][i] = max(prices[i]-prices[i-1]+dp[0][i-1], prices[i]-prices[i-1]+dp[2][i-1])
		dp[2][i] = max(prices[i]-prices[i-1]+dp[0][i-1], prices[i]-prices[i-1]+dp[2][i-1])
		dp[3][i] = dp[1][i-1]
		// dp[0][i] equals dp[4][i], so we can merge both into one.
		dp[4][i] = max(dp[3][i-1], dp[4][i-1])
	}
	return max(max(max(dp[0][n-1], dp[1][n-1]), max(dp[2][n-1], dp[3][n-1])), dp[4][n-1])
}

func maxProfitWithCoolDownBetter(prices []int) int {
	n := len(prices)
	if n < 1 {
		return 0
	}
	// buy - buy or idle
	// hold - hold or sell
	// cooldown - cool down time
	buy, hold, cooldown := 0, 0, 0
	for i := 1; i < n; i++ {
		buy, hold, cooldown = cooldown, max(hold, buy)+prices[i]-prices[i-1], hold
	}
	return max(max(buy, hold), cooldown)
}
