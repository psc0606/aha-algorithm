package greedy

import "sort"

// https://leetcode-cn.com/problems/assign-cookies/
func findContentChildren(g []int, s []int) int {
	ans := 0
	sort.Ints(g)
	sort.Ints(s)
	i, j := 0, 0
	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			ans++
			i++
		}
		j++
	}
	return ans
}
