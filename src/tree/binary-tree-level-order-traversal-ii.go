package tree

// https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/
// In java we can use linked list.
func levelOrderBottom(root *TreeNode) (ans [][]int) {
	if root == nil {
		return ans
	}
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		n := len(queue)
		var level []int
		for n > 0 {
			n--
			root = queue[0]
			queue = queue[1:]
			level = append(level, root.Val)
			if root.Left != nil {
				queue = append(queue, root.Left)
			}
			if root.Right != nil {
				queue = append(queue, root.Right)
			}
		}
		ans = append(ans, level)
	}
	i, j := 0, len(ans)-1
	for i < j {
		ans[i], ans[j] = ans[j], ans[i]
		i++
		j--
	}
	return ans
}
