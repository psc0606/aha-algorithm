package tree

import "testing"

func TestIsSymmetric2(t *testing.T) {
	//
	//     1
	//    / \
	//   2   2
	//  / \ / \
	// 3  4 4  3
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 3},
		},
	}
	isSymmetric2(root)
}
