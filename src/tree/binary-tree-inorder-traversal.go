package tree

// https://leetcode-cn.com/problems/binary-tree-inorder-traversal/
// An elegant implementation by stack and a pointer.
// dfs + stack
//     3
//   /  \
//  9   20
//     /  \
//    15   7
func inorderTraversal(root *TreeNode) (res []int) {
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return
}

// Time: O(n)
// Space: O(n)
func inorderTraversalRecursively(root *TreeNode) []int {
	var ret []int
	doInorderTraversalRecursively(root, &ret)
	return ret
}

// for `collect` as a slice parameter will be modified by function,
// a pointer to slice will be used
func doInorderTraversalRecursively(root *TreeNode, collect *[]int) {
	if root == nil {
		return
	}
	doInorderTraversalRecursively(root.Left, collect)
	*collect = append(*collect, root.Val)
	doInorderTraversalRecursively(root.Right, collect)
}
