package optimalization

import "math"

// https://leetcode-cn.com/problems/maximum-subarray/
// https://leetcode-cn.com/problems/maximum-subarray/solution/zui-da-zi-xu-he-by-leetcode-solution/

// Solution one:
// dynamic programming
// Time: O(n)
// Space: O(1)
// But it update the original array.
// [-2,1,-3,4,-1,2,1,-5,4]
// [-2,1,-2,4,
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return math.MinInt32
	}
	maxVal := nums[0]
	for i := 1; i < len(nums); i++ {
		nums[i] = max(nums[i]+nums[i-1], nums[i])
		if nums[i] > maxVal {
			maxVal = nums[i]
		}
	}
	return maxVal
}

// Solution two:
// Time: O(n)
// Space: O(n)
// greedy
func maxSubArray2(nums []int) int {
	if len(nums) == 0 {
		return math.MinInt32
	}
	maxVal := nums[0]
	for i := 1; i < len(nums); i++ {
		// if num[i-1] > 0, greedy add it, or, discard num[i-1]
		nums[i] = max(nums[i-1], 0) + nums[i]
		if nums[i] > maxVal {
			maxVal = nums[i]
		}
	}
	return maxVal
}
