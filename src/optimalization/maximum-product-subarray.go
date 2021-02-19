package optimalization

// https://leetcode-cn.com/problems/maximum-product-subarray/
// Dynamic programing
// maxDF[i] means the maximum product subarray end with num[i]. But why we need a minDF, minDF[i] means the minimum
// product subarray end with num[i].
// Time: O(n)
// Space: O(n)
func maxProduct(nums []int) int {
	n := len(nums)
	maxDF, minDF := make([]int, n), make([]int, n)
	maxDF[0], minDF[0] = nums[0], nums[0]
	for i := 1; i < n; i++ {
		// When we take the negative numbers into account, we must use a minDF to record minimum.
		maxDF[i] = max(maxDF[i-1]*nums[i], max(minDF[i-1]*nums[i], nums[i]))
		minDF[i] = min(minDF[i-1]*nums[i], min(maxDF[i-1]*nums[i], nums[i]))
	}
	ret := maxDF[0]
	for i := 1; i < n; i++ {
		if maxDF[i] > ret {
			ret = maxDF[i]
		}
	}
	return ret
}
