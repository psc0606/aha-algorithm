package tree

// https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	found := false
	// root -> ... -> p
	// root -> ... -> q
	pathP := findThePathRecursively(root, p, &found)
	pathQ := findThePathRecursively(root, q, &found)
	i := 0
	for i < min(len(pathP), len(pathQ)) && pathP[i] == pathQ[i] {
		i++
	}
	return pathP[i-1]
}

// find the path of node recursively, too slow.
func findThePathRecursively(root *TreeNode, node *TreeNode, found *bool) []*TreeNode {
	var path []*TreeNode
	if root == nil {
		*found = false
		return path
	}
	path = append(path, root)
	if root == node {
		*found = true
		return path
	}
	// find the target, then add to path
	subpath := findThePathRecursively(root.Left, node, found)
	if *found {
		path = append(path, subpath...)
		return path
	}
	subpath = findThePathRecursively(root.Right, node, found)
	if *found {
		path = append(path, subpath...)
		return path
	}
	return path
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Another better implementation
// Time: O(n)
// Space: O(n)
func lowestCommonAncestorRecursively(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}
	left := lowestCommonAncestorRecursively(root.Left, p, q)
	right := lowestCommonAncestorRecursively(root.Right, p, q)
	// that means p and q will be separately left sub-tree and right sub-three.
	if left != nil && right != nil {
		return root
	}

	// that means p and q must be in right sub-tree
	if left == nil {
		return right
	}

	// that means p and q must be in left sub-tree
	if right == nil {
		return left
	}
	// should not occur, because you should make sure p, q is the node of the tree.
	return nil
}
