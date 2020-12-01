package array

// https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/
// Time: O(logn)
// Space: O(1)
// very easy understand
// [5,7,7,8,8,10,11]
func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	left := 0
	right := len(nums) - 1

	findOne := binarySearch(nums, left, right, target)
	if findOne == -1 {
		return res
	}

	leftMost := findOne
	for leftMost != -1 {
		res[0] = leftMost
		leftMost = binarySearch(nums, 0, leftMost-1, target)
	}
	rightMost := findOne
	for rightMost != -1 {
		res[1] = rightMost
		rightMost = binarySearch(nums, rightMost+1, right, target)
	}
	return res
}

func binarySearch(nums []int, left int, right int, target int) int {
	if left < 0 || left >= len(nums) {
		return -1
	}
	if right < 0 || right >= len(nums) {
		return -1
	}
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
