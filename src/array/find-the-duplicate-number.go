package array

import "sort"

// https://leetcode-cn.com/problems/find-the-duplicate-number/
// Required:
// 	1. CAN NOT MODIFY original `nums`
//  2. Time complexity O(n)
//  3. Space complexity O(1)

// Time: O(n)
// Space: O(1), but it will update original array
// If will sort the nums, if the element is unique, index of the element equals to the element self.
// step:
// input: [2, 3, 1, 0, 6, 5, 3]
// i=0, nums[0] != 0, swap -> [1, 3, 2, 0, 6, 5, 3]
// repeat
// i=0, nums[0] != 0, swap -> [3, 1, 2, 0, 6, 5, 3]
// repeat
// i=0, nums[0] != 0, swap -> [0, 1, 2, 3, 6, 5, 3]
// now nums[0] == 0,
// nums[1]=1
// nums[2]=2
// nums[3]=3
// nums[4] != 4, swap -> [0, 1, 2, 3, 3, 5, 6]
// repeat
// nums[4] != 4, find nums[4] == nums[nums[4]], find the duplicates.
func findDuplicate(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] == i {
			continue
		}

		for nums[i] != i {
			// find the duplicate element
			if nums[nums[i]] == nums[i] {
				return nums[i]
			}

			// not equal then swap both
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
	return -1
}

// Sort
// Time: O(nlogn)
// Space: O(1), update original array.
func findDuplicate2(nums []int) int {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return nums[i]
		}
	}
	return -1
}

// Hast map
// Time: O(n)
// Space: O(n)
func findDuplicate3(nums []int) int {
	ht := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		_, ok := ht[nums[i]]
		if ok {
			return nums[i]
		}
		ht[nums[i]] = true
	}
	return -1
}

// https://leetcode-cn.com/problems/find-the-duplicate-number/solution/xun-zhao-zhong-fu-shu-by-leetcode-solution/
// binary search
func findDuplicate4(nums []int) {

}
