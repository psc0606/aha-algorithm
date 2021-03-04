package backtrace

// https://leetcode-cn.com/problems/combination-sum/
// backtrace is a good solution to search all answers.
func combinationSum(candidates []int, target int) [][]int {
	var ret [][]int
	// the search path
	var path []int
	dfsCombinationSum(candidates, 0, path, 0, target, &ret)
	return ret
}

// `index` - use to record position of candidates to avoid duplicate search.
func dfsCombinationSum(candidates []int, index int, path []int, sum int, target int, ret *[][]int) {
	if sum == target {
		// be careful! golang slice is a pointer, copy values
		// arr := make([]int, len(path))
		// copy(arr, path)
		arr := append([]int{}, path...)
		*ret = append(*ret, arr)
		return
	} else if sum > target {
		// no need deep into
		return
	}
	// i >= index, or if `i` starts from 0, will lead to duplicate searching.
	for i := index; i < len(candidates); i++ {
		path = append(path, candidates[i]) // append to tail
		dfsCombinationSum(candidates, i, path, sum+candidates[i], target, ret)
		path = path[0 : len(path)-1] // delete tail
	}
}
