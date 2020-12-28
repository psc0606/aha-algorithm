package tree

import (
	"aha-algorithm/src/util"
	"container/list"
)

// https://leetcode-cn.com/problems/symmetric-tree/
//     1
//    / \
//   2   2
//  / \ / \
// 3  4 4  3

// One: recursive
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isEqualTreeNode(root.Left, root.Right)
}

func isEqualTreeNode(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	if root1.Val != root2.Val {
		return false
	}
	isLeftEq := isEqualTreeNode(root1.Left, root2.Right)
	isRightEq := isEqualTreeNode(root1.Right, root2.Left)
	return isLeftEq && isRightEq
}

// Two: iterate
func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	l := list.New()
	l.PushBack(root.Left)
	l.PushBack(root.Right)
	for l.Len() != 0 {
		e1 := l.Front()
		l.Remove(e1)
		e2 := l.Front()
		l.Remove(e2)

		node1 := e1.Value
		node2 := e2.Value
		if util.IsNil(node1) && util.IsNil(node2) {
			continue
		}
		if util.IsNil(node1) || util.IsNil(node2) {
			return false
		}
		if node1.(*TreeNode).Val != node2.(*TreeNode).Val {
			return false
		}

		l.PushBack(node1.(*TreeNode).Left)
		l.PushBack(node2.(*TreeNode).Right)
		l.PushBack(node1.(*TreeNode).Right)
		l.PushBack(node2.(*TreeNode).Left)
	}
	return true
}
