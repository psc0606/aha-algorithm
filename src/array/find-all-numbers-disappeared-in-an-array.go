package array

// https://leetcode-cn.com/problems/find-all-numbers-disappeared-in-an-array/
// Time: O(n)
// Space: O(1)
// origin: [4,3,2,7,8,2,3,1]
// new: [-4,-3,-2,-7,8,2,-3,-1]
func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		pos := abs(nums[i]) - 1
		nums[pos] = -abs(nums[pos])
	}
	var missing []int
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			missing = append(missing, i+1)
		}
	}
	return missing
}
