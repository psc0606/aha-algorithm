package bit

// https://leetcode-cn.com/problems/single-number/
// Time: O(n)
// Space: O(1)
func singleNumber(nums []int) int {
	// assert at least one elements
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		res ^= nums[i]
	}
	return res
}
