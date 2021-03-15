package tree

// https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/
// recursively dfs
func preorder(root *Node) []int {
	if root == nil {
		return nil
	}
	ret := []int{root.Val}
	for _, node := range root.Children {
		if node != nil {
			child := preorder(node)
			ret = append(ret, child...)
		}
	}
	return ret
}

// non-recursively dfs by stack
func preorderDfsByStack(root *Node) []int {
	if root == nil {
		return nil
	}
	var ret []int
	stack := []*Node{root}
	for len(stack) > 0 {
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret = append(ret, root.Val)
		for i := len(root.Children) - 1; i >= 0; i-- { //reversed add to stack
			if root.Children[i] != nil {
				stack = append(stack, root.Children[i])
			}
		}
	}
	return ret
}
