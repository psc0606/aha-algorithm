package tree

import "errors"

// binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// N-tree node
type Node struct {
	Val      int
	Children []*Node
}

// Build a tree from array, if the given array cannot be construct to a tree, it will give an error.
// Use level traversal to construct a tree.
// For testing purpose.
func BuildTree(arr []interface{}) (*TreeNode, error) {
	if len(arr) == 0 {
		return nil, nil
	}
	if arr[0] == nil {
		return nil, errors.New("the first element of array should not be nil")
	}
	// root TreeNode
	var queue []*TreeNode
	root := &TreeNode{Val: arr[0].(int)}
	queue = append(queue, root)

	index := 1
	for len(queue) != 0 && index < len(arr) {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			return nil, errors.New("nil tree node should not have child")
		}

		// create left child
		val := arr[index]
		var leftChildNode *TreeNode = nil
		if val != nil {
			leftChildNode = &TreeNode{Val: arr[index].(int)}
			node.Left = leftChildNode
		}
		queue = append(queue, leftChildNode)

		// create right child
		index++
		val = arr[index]
		var rightChildNode *TreeNode = nil
		if index < len(arr) && val != nil {
			rightChildNode = &TreeNode{Val: arr[index].(int)}
			node.Right = rightChildNode
		}
		queue = append(queue, rightChildNode)
		index++
	}
	return root, nil
}
