package array

// https://leetcode-cn.com/problems/reverse-pairs/

// Solution1:
// Time: O(n^2)
// Space: O(1)
func reversePairs(nums []int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > 2*nums[j] {
				count++
			}
		}
	}
	return count
}

// Solution2:
// Solved by merge sort, merge first, count the reversed pairs by special binary search.
// Time: O(nlogn)
// Space: O(n)
// example: [2,4,3,5,1]
func reversePairs2(nums []int) int {
	step := 1
	count := 0
	for step < len(nums) {
		step *= 2
		for i := 0; i < len(nums); i += step {
			left := i
			right := i + step - 1
			mid := (left + right) / 2

			var validateRight int
			if right < len(nums) {
				validateRight = right
			} else {
				validateRight = len(nums) - 1
			}
			for j := left; j <= mid && j < len(nums); j++ {
				foundIndex := binarySearchLastIndexSmallThanTarget(nums, mid+1, validateRight, nums[j])
				if foundIndex != -1 {
					count += foundIndex - mid
				}
			}
			// the mid may be invalid because the rest of data may not applicationable for merge
			// [3, 2, 5, 8, 1]
			// when step is 4, the second step data is only a elements.
			if mid < validateRight {
				count += doMergeSort(nums, left, mid, validateRight)
			}
		}
	}
	return count
}

// Do the reversed pairs counting in merging two array segment.
// Space: O(n)
func doMergeSort(nums []int, left int, mid int, right int) int {
	if left >= right {
		return 0
	}
	count := 0
	p := left
	q := mid + 1
	help := make([]int, right-left+1)
	index := 0
	for p <= mid && q <= right {
		if nums[p] < nums[q] {
			help[index] = nums[p]
			index++
			p++
		} else {
			help[index] = nums[q]
			index++
			q++
		}
	}
	// need count
	for p <= mid {
		help[index] = nums[p]
		index++
		p++
	}
	// no need count any more, it rest is much bigger
	for q <= right {
		help[index] = nums[q]
		index++
		q++
	}
	index = left
	for _, v := range help {
		nums[index] = v
		index++
	}
	return count
}

// binary search the last value small than the target value, then return the index.
// if not found, -1 return.
func binarySearchLastIndexSmallThanTarget(nums []int, left int, right int, target int) int {
	p := left
	q := right
	for p <= q {
		mid := (p + q) / 2
		if 2*nums[mid] >= target {
			q = mid - 1
		} else {
			p = mid + 1
		}
	}
	if q < left {
		return -1
	}
	if p > right {
		return right
	}
	return q
}
