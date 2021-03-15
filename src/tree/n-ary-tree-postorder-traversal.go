package tree

import "aha-algorithm/src/util"

// https://leetcode-cn.com/problems/n-ary-tree-postorder-traversal
func postorder(root *Node) []int {
	if root == nil {
		return nil
	}
	var ret []int
	for _, node := range root.Children {
		if node != nil {
			child := postorder(node)
			ret = append(ret, child...)
		}
	}
	ret = append(ret, root.Val)
	return ret
}

// non-recursively dfs by stack
func postorderDfsByStack(root *Node) []int {
	if root == nil {
		return nil
	}
	var ret []int
	stack := []*Node{root}
	for len(stack) > 0 {
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret = append(ret, root.Val)
		for _, node := range root.Children {
			if node != nil {
				stack = append(stack, node)
			}
		}
	}
	return util.SliceReverse(ret)
}
