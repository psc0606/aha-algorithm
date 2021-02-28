package optimalization

import (
	"math"
	"sort"
)

// https://leetcode-cn.com/problems/coin-change/
// dynamic programming
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)

	// init condition
	dp[0] = 0
	for money := 1; money < len(dp); money++ {
		dp[money] = -1
	}

	for money := 1; money < len(dp); money++ {
		dp[money] = math.MaxInt64
		for _, c := range coins {
			if money-c >= 0 && dp[money-c] != -1 {
				dp[money] = min(dp[money], dp[money-c]+1)
			}
		}
		if dp[money] == math.MaxInt64 {
			// restore to -1
			dp[money] = -1
		}
	}
	return dp[amount]
}

// If greedy policy is used, every time I pick a largest coin, it may not get the final result.
// For example:
// 	coins: [3, 8], amount: 9
// If you first pick 8, then 3, you cannot get the answer. Actual, it's 3 + 3 + 3.
// error solution
func coinChangeGreedyPolicy(coins []int, amount int) int {
	count := 0
	sort.Sort(IntSlice(coins))
	for i := len(coins) - 1; i >= 0; i-- {
		for amount >= coins[i] {
			amount -= coins[i]
			count++
		}
	}
	if amount != 0 {
		count = -1
	}
	return count
}

type IntSlice []int

func (s IntSlice) Len() int           { return len(s) }
func (s IntSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s IntSlice) Less(i, j int) bool { return s[i] < s[j] }
