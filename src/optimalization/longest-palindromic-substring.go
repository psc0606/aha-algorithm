package optimalization

// https://leetcode-cn.com/problems/longest-palindromic-substring/
// First: dynamic programming
// Time: O(n^2)
// Space: O(n^2)
func longestPalindrome(s string) string {
	n := len(s)
	if n < 1 {
		return ""
	}
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		// init diag
		dp[i][i] = true
	}

	max, start, end := 1, 0, 0
	for k := 0; k < n; k++ {
		i := 0
		for j := k + 1; j < n; j++ {
			if s[i] == s[j] {
				// length < 3
				// neighbor i, i+1
				if j-i+1 < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			} else {
				dp[i][j] = false
			}

			// update answer
			if j-i+1 > max && dp[i][j] {
				// update max palindrome length
				max = j - i + 1
				start, end = i, j
			}
			i++
		}
	}
	return s[start : end+1]
}
