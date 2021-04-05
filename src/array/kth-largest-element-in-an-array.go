package array

// https://leetcode-cn.com/problems/kth-largest-element-in-an-array/
// input: [3,2,1,5,6,4], k = 2
// output: 2
// We can sort, then find the kth largest element. But its time complexity is O(nlogn).
// How can we do it by O(n)?
func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, k, 0, len(nums)-1)
}

func quickSelect(nums []int, k int, start, end int) int {
	q := partitionKth(nums, start, end)
	if q == k-1 {
		return nums[q]
	} else {
		if q > k-1 {
			q = quickSelect(nums, k, start, q-1)
		} else {
			q = quickSelect(nums, k, q+1, end)
		}
		return q
	}
}

func partitionKth(nums []int, start, end int) int {
	pivot := nums[end]
	i := start
	for j := start; j < end; j++ { // not include end, end is the pivot
		if nums[j] > pivot {
			swap(nums, i, j)
			i++
		}
	}
	swap(nums, i, end)
	return i
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

// Use minimum heap, O(nlogk), only keep k elements in heap.
// Use maximum heap, O(nlogn), delete k-1 elements, then get the top of element in heap.
