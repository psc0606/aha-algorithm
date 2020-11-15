package array

// https://leetcode-cn.com/problems/intersection-of-two-arrays/
func intersection(nums1 []int, nums2 []int) []int {
	hashmap := make(map[int]bool)
	for _, ele := range nums1 {
		// avoid nums1 duplicates
		hashmap[ele] = true
	}
	var intersection []int
	for _, ele := range nums2 {
		_, ok := hashmap[ele]
		if ok {
			intersection = append(intersection, ele)
			delete(hashmap, ele)
		}
	}
	return intersection
}
