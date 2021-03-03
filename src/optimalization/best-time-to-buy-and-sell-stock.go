package optimalization

// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/
// Time: O(n)
// Space: O(1)
func maxProfit(prices []int) int {
	if len(prices) < 1 {
		return 0
	}
	minPrice := prices[0]
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		maxProfit = max(maxProfit, prices[i]-minPrice)
		minPrice = min(minPrice, prices[i])
	}
	return maxProfit
}
