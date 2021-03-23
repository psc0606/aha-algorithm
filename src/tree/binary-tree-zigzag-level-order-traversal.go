package tree

import "aha-algorithm/src/util"

// https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/
//       1
//     /  \
//    2    3
//  /  \  /
// 4 nil  5

// zigzag: [ [1]
//          [3, 2]  // just reverse odd level
//          [4, 5]
//          ]
// Still use level order traversal by bfs, But only reverse the odd result.
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var queue []*TreeNode
	queue = append(queue, root)
	var ret [][]int
	zigzag := 0
	for len(queue) > 0 {
		var level []int
		count := len(queue)
		for count > 0 {
			var node *TreeNode
			node = queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			count--
			level = append(level, node.Val)
		}
		// just reverse the odd level nodes
		if zigzag%2 == 1 {
			util.SliceReverse(level)
		}
		zigzag++
		ret = append(ret, level)
	}
	return ret
}
