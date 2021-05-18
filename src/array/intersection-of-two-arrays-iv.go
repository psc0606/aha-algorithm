package array

// intersection of two sorted array.
// NOTE: both of two WITHOUT DUPLICATE elements.
// two arrays are too different in size, such as len(nums1) is 38,
// but len(nums2) is 3358, 3358 is much bigger than 38.
// Time:O(mlogn), n > m * m
func intersectIV(nums1 []int, nums2 []int) []int {
	var ans []int
	m := len(nums1)
	n := len(nums2)
	if m > n {
		intersectIV(nums2, nums1)
	}

	var binarySearch func(arr []int, start, end, target int) int
	// if not find target, return the last left position.
	binarySearch = func(arr []int, start, end, target int) int {
		if start > end {
			return -1
		}
		i, j := start, end
		for i <= j {
			mid := (i + j) / 2
			if arr[mid] == target {
				return mid
			} else if arr[mid] < target {
				i = mid + 1
			} else {
				j = mid - 1
			}
		}
		// not found but return the left.
		return i
	}

	index := 0
	for i := 0; i < m; i++ {
		index = binarySearch(nums2, index, n-1, nums1[i])
		// check equal again
		if nums1[i] == nums2[index] {
			ans = append(ans, nums1[i])
		}
	}
	return ans
}

// 1,3,5, 7
// 4, 5

// 1,3,5,7,8
