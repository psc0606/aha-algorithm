package tree

// https://leetcode-cn.com/problems/validate-binary-search-tree/
// validate a binary search tree.
//        5
//       / \
//      4   6
//  	   /  \
//        3    7
func isValidBST(root *TreeNode) bool {
	return _isValidBST(root)
}

func _isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	valid := true

	// find the right most node of the left child.
	if root.Left != nil {
		if findRightMostNode(root.Left).Val >= root.Val {
			valid = false
		}
	}

	// find the left most node of the right child.
	if root.Right != nil {
		if findLeftMostNode(root.Right).Val <= root.Val {
			valid = false
		}
	}
	return valid && _isValidBST(root.Left) && _isValidBST(root.Right)
}

func findLeftMostNode(root *TreeNode) *TreeNode {
	if root.Left == nil {
		return root
	}
	return findLeftMostNode(root.Left)
}

func findRightMostNode(root *TreeNode) *TreeNode {
	if root.Right == nil {
		return root
	}
	return findRightMostNode(root.Right)
}
