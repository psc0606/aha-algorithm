package array

// https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/
// intersect both of two array may duplicate elements.
func intersect(nums1 []int, nums2 []int) []int {
	var ans []int
	m := len(nums1)
	n := len(nums2)
	if m > n {
		intersect(nums2, nums1)
	}

	// m < n for saving memory
	ht := make(map[int]int)
	for i := 0; i < m; i++ {
		ht[nums1[i]]++
	}
	for i := 0; i < n; i++ {
		v, ok := ht[nums2[i]]
		if ok && v > 0 {
			ans = append(ans, nums2[i])
			ht[nums2[i]]--
		}
	}
	return ans
}

// sort + double pointer
