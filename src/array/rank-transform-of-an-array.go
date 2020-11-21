package array

import "sort"

// https://leetcode-cn.com/problems/rank-transform-of-an-array/
// Time: O(nlogn)
// Space: O(n)
func arrayRankTransform(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}
	sortArr := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		sortArr[i] = arr[i]
	}
	sort.Ints(sortArr)

	hashmap := make(map[int]int)
	rank := 0
	for i := 0; i < len(sortArr); i++ {
		_, ok := hashmap[sortArr[i]]
		if ok {
			hashmap[sortArr[i]] = rank
		} else {
			rank++
			hashmap[sortArr[i]] = rank
		}
	}

	// reuse the slice
	for i := 0; i < len(arr); i++ {
		arr[i] = hashmap[arr[i]]
	}
	return arr
}
