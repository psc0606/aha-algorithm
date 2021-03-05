package bit

// https://leetcode-cn.com/problems/counting-bits/
//
// Good first:
// https://leetcode-cn.com/problems/counting-bits/solution/bi-te-wei-ji-shu-by-leetcode-solution-0t1i/
// Time: O(n)
// Space: O(n)
func countBits(num int) []int {
	if num < 0 {
		return nil
	}
	ht := make(map[int]int)
	ret := make([]int, num+1)
	ret[0] = 0
	ht[0] = 0
	for i := 1; i <= num; i++ {
		v, _ := ht[i&(i-1)]
		ret[i] = v + 1
		ht[i] = ret[i]
	}
	return ret
}
