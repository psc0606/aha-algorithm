package tree

// https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/
// Space: O(1), not O(n)
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	tree2List(root)
}

func tree2List(root *TreeNode) *TreeNode {
	if root.Right == nil && root.Left == nil {
		return root
	}
	var leftTail, rightTail *TreeNode
	if root.Left != nil {
		leftTail = tree2List(root.Left)
		leftTail.Right = root.Right
	}

	if root.Right != nil {
		rightTail = tree2List(root.Right)
	}
	if root.Left != nil {
		root.Right = root.Left
		root.Left = nil
	}
	if rightTail != nil {
		return rightTail
	} else {
		return leftTail
	}
}
