package tree

import "aha-algorithm/src/util"

// https://leetcode-cn.com/problems/path-sum-ii/
// dfs
func pathSum(root *TreeNode, targetSum int) [][]int {
	var ret [][]int
	preOrderDfs(root, targetSum, []int{}, &ret)
	return ret
}

func preOrderDfs(root *TreeNode, targetSum int, path []int, ans *[][]int) {
	if root == nil {
		return
	}
	path = append(path, root.Val)
	if root.Left == nil && root.Right == nil && targetSum-root.Val == 0 {
		*ans = append(*ans, append([]int{}, path...))
		return
	}
	preOrderDfs(root.Left, targetSum-root.Val, path, ans)
	preOrderDfs(root.Right, targetSum-root.Val, path, ans)
	// path slice will not be modified by function
	// path = path[:len(path)-1]
}

// bfs + hashmap
func pathSumBfs(root *TreeNode, targetSum int) (ans [][]int) {
	if root == nil {
		return
	}
	ht := make(map[*TreeNode]*TreeNode)
	queue, queueSum := make([]*TreeNode, 0), make([]int, 0)
	queue = append(queue, root)
	queueSum = append(queueSum, 0)
	for len(queue) != 0 {
		root, queue = queue[0], queue[1:]
		rec := queueSum[0] + root.Val
		queueSum = queueSum[1:]

		// find target
		if root.Left == nil && root.Right == nil && rec == targetSum {
			path := []int{root.Val}
			p := root
			for p != nil {
				v, ok := ht[p]
				if ok {
					path = append(path, v.Val)
				}
				p = v
			}
			util.SliceReverse(path)
			ans = append(ans, path)
			continue
		}
		if root.Left != nil {
			ht[root.Left] = root
			queue = append(queue, root.Left)
			queueSum = append(queueSum, rec)
		}
		if root.Right != nil {
			ht[root.Right] = root
			queue = append(queue, root.Right)
			queueSum = append(queueSum, rec)
		}
	}
	return
}
