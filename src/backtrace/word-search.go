package backtrace

// https://leetcode-cn.com/problems/word-search/
// Backtrace is a great solution to do this.
// Time: O(M*N*3^L), L is the length of word.
// Space: O(min(L)), do not ignore stack length.
func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == word[0] {
				ret := dfs(board, &word, i, j, 0)
				if ret {
					return true
				}
			}
		}
	}
	return false
}

// A trick is how to reuse original board to avoid duplicate access.
func dfs(board [][]byte, word *string, r int, c int, index int) bool {
	// found
	if len(*word) == index {
		return true
	}
	if r < 0 || r >= len(board) {
		return false
	}
	if c < 0 || c >= len(board[0]) {
		return false
	}
	if board[r][c] != (*word)[index] {
		return false
	}
	tmp := board[r][c]
	// marked as a special value which means already accessed.
	board[r][c] = 0
	ret := dfs(board, word, r+1, c, index+1) ||
		dfs(board, word, r-1, c, index+1) ||
		dfs(board, word, r, c+1, index+1) ||
		dfs(board, word, r, c-1, index+1)
	// recover the original value
	board[r][c] = tmp
	return ret
}
