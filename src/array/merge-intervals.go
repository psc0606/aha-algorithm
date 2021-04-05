package array

import "sort"

// https://leetcode-cn.com/problems/merge-intervals/
// eg: intervals = [[1,3],[2,6],[8,10],[15,18]]
//     out = [[1,6],[8,10],[15,18]]
// Time: O(nlogn)
func mergeIntervals(intervals [][]int) [][]int {
	n := len(intervals)
	if n <= 0 {
		return [][]int{}
	}
	// sort by left value
	sort.Sort(IntSlice(intervals))
	var ret [][]int
	for i := 0; i < n; i++ {
		k := len(ret)
		L, R := intervals[i][0], intervals[i][1]
		if k == 0 || ret[k-1][1] < L {
			ret = append(ret, []int{L, R})
		} else {
			ret[k-1][1] = max(ret[k-1][1], R)
		}
	}
	return ret
}

type IntSlice [][]int

func (s IntSlice) Len() int           { return len(s) }
func (s IntSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s IntSlice) Less(i, j int) bool { return s[i][0] < s[j][0] }
