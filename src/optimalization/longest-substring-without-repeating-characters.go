package optimalization

// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/

// First: dynamic programming
// Time: O(n)
// Space: O(n)
func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}
	ans := 1
	ht := make([]int, 256)
	for i := 0; i < 256; i++ {
		ht[i] = -1
	}
	dp := make([]int, n)
	ht[s[0]] = 0
	dp[0] = 0
	for i := 1; i < n; i++ {
		dp[i] = max(dp[i-1], ht[s[i]]+1)
		ht[s[i]] = i
		ans = max(ans, i-dp[i]+1)
	}
	return ans
}

// Slide window
func lengthOfLongestSubstring2(s string) int {
	if len(s) == 0 {
		return 0
	}
	ht := make(map[uint8]int)
	ht[s[0]] = 0
	i := 0
	ans := 1

	for j := 1; j < len(s); j++ {
		index, ok := ht[s[j]]
		if ok {
			// `index` may be out of window.
			i = max(i, index+1)
		}
		ans = max(j-i+1, ans)
		ht[s[j]] = j
	}
	return ans
}
