package string

// https://leetcode-cn.com/problems/word-break/
// Time: O(n^2)
// Space: O(n)
// dp[i] means the substring from 0 to i (inclusive) which can be word broke.
func wordBreak(s string, wordDict []string) bool {
	if len(s) == 0 {
		return false
	}

	hashmap := make(map[string]bool)
	for _, v := range wordDict {
		hashmap[v] = true
	}

	// for processing convenience, the length of dp array is len(s)+1.
	// and init dp[0]=true
	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 0; i < len(s); i++ {
		for j := i; j >= 0; j-- {
			_, ok := hashmap[s[j:i+1]]
			if dp[j] && ok {
				dp[i+1] = true
				break
			}
		}
	}
	return dp[len(s)]
}
