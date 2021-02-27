package backtrace

// https://leetcode-cn.com/problems/generate-parentheses/
// Time: O(4^n / n^0.5)
func generateParenthesis(n int) []string {
	if n <= 0 {
		return nil
	}
	var ans []string
	var curr string
	backtrack(0, 0, n, curr, &ans)
	return ans
}

// `open` - the number of "(" in the `curr`
// `close` - the number of ")" in the `curr`
// `curr` - the current dfs string
// `ans` - store the final result
func backtrack(open int, close int, n int, curr string, ans *[]string) {
	// the number "(" and ")" reach max.
	if open == n && close == n {
		*ans = append(*ans, curr)
		return
	}
	if open < n {
		backtrack(open+1, close, n, curr+"(", ans)
	}
	// the number of "(" >= the number of ")", or it will lead invalid parenthesis
	if close < n && open > close {
		backtrack(open, close+1, n, curr+")", ans)
	}
}
