package tree

// https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/
// dfs
// care:
//     1
//   /  \
// nil   2
//     /  \
//   nil   3
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := minDepth(root.Left)
	right := minDepth(root.Right)
	if left == 0 {
		// means no left child, but it can not as minimum left child length
		return right + 1
	}
	if right == 0 {
		return left + 1
	}
	return min(left, right) + 1
}

// bfs level traversal
func minDepthBfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	minDepth := 0
	for len(queue) > 0 {
		c := len(queue)
		minDepth += 1
		for c > 0 {
			ele := queue[0]
			queue = queue[1:]
			if ele.Left == nil && ele.Right == nil {
				return minDepth
			}
			if ele.Left != nil {
				queue = append(queue, ele.Left)
			}
			if ele.Right != nil {
				queue = append(queue, ele.Right)
			}
			c--
		}
	}
	return 0
}
