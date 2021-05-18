package array

// intersection of two sorted array.
// NOTE: both of two may contains DUPLICATE elements.
// Time: O(m+n)
func intersectIII(nums1 []int, nums2 []int) []int {
	var ans []int
	m := len(nums1)
	n := len(nums2)
	i, j := 0, 0
	for i < m && j < n {
		if nums1[i] == nums2[j] {
			ans = append(ans, nums1[i])
			i++
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}
	return ans
}

// If two arrays are too different in size, how can we optimize ??
// see intersect IV
