package tree

// https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal/
// @see https://leetcode-cn.com/problems/binary-tree-level-order-traversal/
func nTreeLevelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}
	var ret [][]int
	queue := []*Node{root}
	for len(queue) > 0 {
		var level []int
		c := len(queue)
		for c > 0 {
			ele := queue[0]
			queue = queue[1:]
			level = append(level, ele.Val)
			for _, node := range ele.Children {
				if node != nil {
					queue = append(queue, node)
				}
			}
			c--
		}
		ret = append(ret, level)
	}
	return ret
}
