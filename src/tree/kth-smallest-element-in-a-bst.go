package tree

// https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/
// dfs + array
// in-order bst is a sorted array.
// Time: O(n)
// Space: O(n)
// `k` - small then number of three nodes.
func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return -1
	}
	var nodes []int
	inorder(root, &nodes)
	return nodes[k-1]
}

func inorder(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, arr)
	*arr = append(*arr, root.Val)
	inorder(root.Right, arr)
}

// dfs with stack
// Time: O(logn + k)
// Space: O(1)
func kthSmallestBfs(root *TreeNode, k int) int {
	if root == nil {
		return -1
	}
	var stack []*TreeNode

	for {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// root is nil now, then pop the parent
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		k--
		if k == 0 {
			return root.Val
		}
		root = root.Right
	}
}
