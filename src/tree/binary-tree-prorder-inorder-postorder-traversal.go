package tree

// https://www.nowcoder.com/practice/a9fec6c46a684ad5a3abd4e365a9d362
// Input:    1
//         2  3
// Output: [[1,2,3], [2,1,3],[2,3,1]]
// [preorder, inorder, postorder]

// dfs only once
func threeOrders(root *TreeNode) [][]int {
	var ans [][]int
	if root == nil {
		return ans
	}
	ans = make([][]int, 3)
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		ans[0] = append(ans[0], root.Val)
		dfs(root.Left)
		ans[1] = append(ans[1], root.Val)
		dfs(root.Right)
		ans[2] = append(ans[2], root.Val)
	}
	dfs(root)
	return ans
}
