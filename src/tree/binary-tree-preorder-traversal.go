package tree

// https://leetcode-cn.com/problems/binary-tree-preorder-traversal/description/
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	ret := []int{root.Val}
	left := preorderTraversal(root.Left)
	if left != nil {
		ret = append(ret, left...)
	}
	right := preorderTraversal(root.Right)
	if right != nil {
		ret = append(ret, right...)
	}
	return ret
}

// dfs + stack
//     3
//   / \
//  9  20
//    /  \
//   15   7
func preorderTraversalBfs(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var ret []int
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		ele := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret = append(ret, ele.Val)
		if ele.Right != nil {
			stack = append(stack, ele.Right)
		}
		if ele.Left != nil {
			stack = append(stack, ele.Left)
		}
	}
	return ret
}
