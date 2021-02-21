package tree

// https://leetcode-cn.com/problems/binary-tree-level-order-traversal/
// Tree level traversal, but it's output must require format as following:
//     3
//   / \
//  9  20
//    /  \
//   15   7
// output:
// [
//  [3],
//  [9,20],
//  [15,7]
//]
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var queue []*TreeNode
	queue = append(queue, root)
	currLevelNum, nextLevelNum := 1, 0
	var ret [][]int
	var level []int
	for len(queue) > 0 {
		node := queue[0]
		level = append(level, node.Val)
		queue = queue[1:]
		currLevelNum--
		if node.Left != nil {
			queue = append(queue, node.Left)
			nextLevelNum++
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
			nextLevelNum++
		}
		// when the current level number counting down to zero,
		// it means it will travel nex level.
		if currLevelNum == 0 {
			currLevelNum, nextLevelNum = nextLevelNum, 0
			ret = append(ret, level)
			level = []int{}
		}
	}
	return ret
}

// for a better understanding
func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var queue []*TreeNode
	queue = append(queue, root)
	var ret [][]int
	for len(queue) > 0 {
		var level []int
		// get current level node numbers
		count := len(queue)
		for count > 0 {
			node := queue[0]
			queue = queue[1:]
			count--

			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ret = append(ret, level)
	}
	return ret
}
