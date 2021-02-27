package backtrace

// https://leetcode-cn.com/problems/permutations/
// backtrace
func permute(nums []int) [][]int {
	var ans [][]int
	var curr []int
	_permute(nums, 0, curr, &ans)
	return ans
}

// `nums` - input nums
// `index` - the current element should be append to curr.
// `curr` - the current dfs array to store permute
// `ans` - the final result.
func _permute(nums []int, index int, curr []int, ans *[][]int) {
	// reach the end
	if index == len(nums) {
		// we must copy the slice, because slice will is a pointer, it will be changed by upper caller.
		tmp := make([]int, len(curr))
		copy(tmp, curr)
		// copy
		*ans = append(*ans, tmp)
		return
	}
	ele := nums[index]
	curr = append(curr, ele)
	for i := 0; i < len(curr); i++ {
		j := len(curr) - 1
		swap(curr, i, j)
		_permute(nums, index+1, curr, ans)
		// swap back for next
		swap(curr, i, j)
	}
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
