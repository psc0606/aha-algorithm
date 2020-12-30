package tree

// https://leetcode-cn.com/problems/merge-two-binary-trees/
// 输入:
// 		  Tree 1                     Tree 2
//          1                         2
//         / \                       / \
//        3   2                     1   3
//       /                           \   \
//      5                             4   7
//输出:
//合并后的树:
//	     3
//	    / \
//	   4   5
//	  / \   \
//	 5   4   7
// One: recursive
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	newLeft := mergeTrees(t1.Left, t2.Left)
	newRight := mergeTrees(t1.Right, t2.Right)
	return &TreeNode{
		Val:   t1.Val + t2.Val,
		Left:  newLeft,
		Right: newRight,
	}
}
