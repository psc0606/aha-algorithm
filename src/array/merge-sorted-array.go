package array

// https://leetcode-cn.com/problems/merge-sorted-array/
// Merge from end to start
// Time: O(m+n)
func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j := m-1, n-1
	c := len(nums1) - 1
	for i >= 0 && j >= 0 {
		if nums2[j] >= nums1[i] {
			nums1[c] = nums2[j]
			j--
		} else {
			nums1[c] = nums1[i]
			i--
		}
		c--
	}
	for ; i >= 0; i-- {
		nums1[c] = nums1[i]
		c--
	}
	for ; j >= 0; j-- {
		nums1[c] = nums2[j]
		c--
	}
}
