package optimalization

// https://leetcode-cn.com/problems/house-robber/
// Time: O(n)
// Space: O(n)
// [1,2,3,1]
func rob(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	op := make([]int, len(nums))
	op[0] = nums[0]
	op[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		if op[i-2]+nums[i] > op[i-1] {
			op[i] = op[i-2] + nums[i]
		} else {
			op[i] = op[i-1]
		}
	}
	return op[len(op)-1]
}

// two: can also update nums in place.
// three: can also use two variable.
