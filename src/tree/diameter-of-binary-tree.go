package tree

// https://leetcode-cn.com/problems/diameter-of-binary-tree/
//      1
//     / \
//    2   3
//   / \
//  4   5
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	diameter := 0
	depth(root, &diameter)
	return diameter
}

func depth(root *TreeNode, diameter *int) int {
	if root == nil {
		return 0
	}
	// left child depth
	leftChildDepth := depth(root.Left, diameter)
	// right child depth
	rightChildDepth := depth(root.Right, diameter)

	// update diameter
	newDiameter := leftChildDepth + rightChildDepth
	if *diameter < newDiameter {
		*diameter = newDiameter
	}

	if leftChildDepth > rightChildDepth {
		return leftChildDepth + 1
	} else {
		return rightChildDepth + 1
	}
}
