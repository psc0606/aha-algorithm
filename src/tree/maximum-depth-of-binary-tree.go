package tree

// https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/
// dfs + divide and rule
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxLeftDepth := maxDepth(root.Left)
	maxRightDepth := maxDepth(root.Right)
	return max(maxLeftDepth, maxRightDepth) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// bfs
func maxDepthBfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	maxDepth := 0
	for len(queue) > 0 {
		c := len(queue)
		maxDepth += 1
		for c > 0 {
			ele := queue[0]
			queue = queue[1:]
			if ele.Left != nil {
				queue = append(queue, ele.Left)
			}
			if ele.Right != nil {
				queue = append(queue, ele.Right)
			}
			c--
		}
	}
	return maxDepth
}
