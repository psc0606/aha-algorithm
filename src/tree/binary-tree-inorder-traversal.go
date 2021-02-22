package tree

// https://leetcode-cn.com/problems/binary-tree-inorder-traversal/
// An elegant implementation by stack and a pointer.
func inorderTraversal(root *TreeNode) []int {
	var stack []*TreeNode
	var ret []int
	ptr := root
	for len(stack) > 0 || ptr != nil {
		if ptr != nil {
			stack = append(stack, ptr)
			ptr = ptr.Left
		} else {
			// pop a element from stack top
			ptr = stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			ret = append(ret, ptr.Val)
			ptr = ptr.Right
		}
	}
	return ret
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
