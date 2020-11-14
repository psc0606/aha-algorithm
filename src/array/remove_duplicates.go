package array

// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/
// remove duplicate elements from sorted array with once traversal.
// [0,0,1,1,1,2,2,3,3,4]
func removeDuplicates(nums []int) int {
	lens := len(nums)
	cursor := 0
	for i := 1; i < lens; i++ {
		if nums[i] != nums[cursor] {
			cursor++
			nums[cursor] = nums[i]
		}
	}
	return cursor + 1
}
