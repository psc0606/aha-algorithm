package array

// https://leetcode-cn.com/problems/move-zeroes/
// The problem is more like the gc algorithm,

// implement by swap
func moveZeroes(nums []int) {
	i := 0
	j := 0
	for {
		// i find first zero
		for ; i < len(nums) && nums[i] != 0; i++ {
		}
		// reach the end
		if i == len(nums) {
			return
		}
		// if not reach the end, j find the first non-zero from i+1 or j+1
		// for j = i+1; j < len(nums) && nums[j] == 0; j++ {
		for j = max(i+1, j+1); j < len(nums) && nums[j] == 0; j++ {
		}
		if j == len(nums) {
			for i < len(nums) {
				nums[i] = 0
				i++
			}
		} else {
			nums[i] = nums[j]
			nums[j] = 0
		}
	}
}

// implement by copy algorithm, more like gc algorithm
// better and simple answer
func moveZeroes2(nums []int) {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
	}
	for j < len(nums) {
		nums[j] = 0
		j++
	}
}

func max(a int, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}
