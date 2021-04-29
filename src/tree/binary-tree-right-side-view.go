package tree

// https://leetcode-cn.com/problems/binary-tree-right-side-view/
// bfs
func rightSideView(root *TreeNode) []int {
	var ans []int
	if root == nil {
		return ans
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		n := len(queue)
		for n > 0 {
			n--
			root, queue = queue[0], queue[1:]
			if root.Left != nil {
				queue = append(queue, root.Left)
			}
			if root.Right != nil {
				queue = append(queue, root.Right)
			}
		}
		// when exist for loop, root is the right-most node
		ans = append(ans, root.Val)
	}
	return ans
}

// dfs
func rightSideViewII(root *TreeNode) []int {
	var ans []int
	if root == nil {
		return ans
	}
	var dfs func(root *TreeNode, depth int)
	dfs = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		if depth > len(ans) {
			ans = append(ans, root.Val)
		}
		dfs(root.Right, depth+1)
		dfs(root.Left, depth+1)
	}
	dfs(root, 1)
	return ans
}
