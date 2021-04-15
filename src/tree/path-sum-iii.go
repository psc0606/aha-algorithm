package tree

// https://leetcode-cn.com/problems/path-sum-iii/
// bfs + dfs
// Time: O(n^2)
func pathSumIII(root *TreeNode, sum int) (ans int) {
	if root == nil {
		return 0
	}

	var dfs func(*TreeNode, int, *int)
	dfs = func(root *TreeNode, sum int, count *int) {
		if root == nil {
			return
		}
		sum -= root.Val
		if sum == 0 {
			*count++
		}
		dfs(root.Left, sum, count)
		dfs(root.Right, sum, count)
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		n := len(queue)
		for n >= 0 {
			n--
			root, queue = queue[0], queue[1:]
			dfs(root, sum, &ans)
			if root.Left != nil {
				queue = append(queue, root.Left)
			}
			if root.Right != nil {
				queue = append(queue, root.Right)
			}
		}
	}
	return
}

// dfs + dfs
// Time: O(n^2)
func pathSumIII2(root *TreeNode, sum int) (ans int) {
	if root == nil {
		return 0
	}

	var dfs func(*TreeNode, int, *int)
	dfs = func(root *TreeNode, sum int, count *int) {
		if root == nil {
			return
		}
		sum -= root.Val
		if sum == 0 {
			*count++
		}
		dfs(root.Left, sum, count)
		dfs(root.Right, sum, count)
	}
	dfs(root, sum, &ans)
	return ans + pathSumIII2(root.Left, sum) + pathSumIII(root.Right, sum)
}

// Better
// https://leetcode-cn.com/problems/path-sum-iii/solution/dui-qian-zhui-he-jie-fa-de-yi-dian-jie-s-dey6/
// prefix sum + hashmap
func pathSumIII3(root *TreeNode, sum int) (ans int) {
	if root == nil {
		return 0
	}
	ht := make(map[int]int)
	ht[0] = 1
	var dfs func(*TreeNode, int, int) int
	dfs = func(root *TreeNode, sum int, currSum int) (ans int) {
		if root == nil {
			return 0
		}
		currSum += root.Val
		c := 0
		v, ok := ht[currSum-sum]
		if !ok {
			v = 0
		}
		c += v
		ht[currSum]++
		c += dfs(root.Left, sum, currSum) + dfs(root.Right, sum, currSum)
		ht[currSum]--
		return c
	}
	return dfs(root, sum, 0)
}

//          5
//      4      8
//  11  nil 13  4
//7  2     5 1

// 1

//   0
//  1 1
