package array

// https://leetcode-cn.com/problems/rotate-array/
// a[i] -> a[(i + k) % n]
// Time: O(n)
// Space: O(n)
func rotate(nums []int, k int) {
	n := len(nums)
	tmp := make([]int, n)
	for i := 0; i < n; i++ {
		tmp[(i+k)%n] = nums[i]
	}

	// restore
	for i := 0; i < n; i++ {
		nums[i] = tmp[i]
	}
}

// Reverse the whole array first:
// eg: input [1, 2, 3, 4, 5, 6, 7], k=3
// reverse -> [7, 6, 5, 4, 3, 2, 1]
// reverse [0~k-1], [k, length-1] separately
// Time: O(n)
// Space: O(1)
func rotate3(nums []int, k int) {
	ReverseArray(nums, 0, len(nums)-1)
	k %= len(nums)
	ReverseArray(nums, 0, k-1)
	ReverseArray(nums, k, len(nums)-1)
}

func ReverseArray(nums []int, start, end int) {
	i, j := start, end
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
