package tree

// https://leetcode-cn.com/problems/unique-binary-search-trees/
// Dynamic programing
// G(n) = G(0)*G(n-1)+G(1)*(n-2)+...+G(n-1)*G(0), G(0)=0, G(1)=1, n >=2
// G(n) is the number of unique bst: when `1` is the root, the rest must be the right child of the tree,
// when `2` is the root of tree, `1` must be the left child of the tree, ... when `n` must be the tree, the rest n-1 must
// be the left child of the tree.

// Time: O(n^2)
// Space: O(n)
func numTrees(n int) int {
	if n < 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i < n+1; i++ {
		for j := 0; j < i; j++ {
			dp[i] += dp[j] * dp[i-j-1]
		}
	}
	return dp[n]
}
