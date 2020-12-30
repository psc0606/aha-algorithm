package tree

// https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/
// dfs + divide and rule
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxLeftDepth := maxDepth(root.Left)
	maxRightDepth := maxDepth(root.Right)
	if maxLeftDepth > maxRightDepth {
		return maxLeftDepth + 1
	} else {
		return maxRightDepth + 1
	}
}
