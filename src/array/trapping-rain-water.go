package array

// https://leetcode-cn.com/problems/trapping-rain-water/
// https://leetcode-cn.com/problems/trapping-rain-water/solution/jie-yu-shui-by-leetcode/

// Solution1:
// Time: O(n)
// Space: O(n)
func trap(height []int) int {
	n := len(height)
	if n <= 0 {
		return 0
	}
	left := make([]int, n)
	right := make([]int, n)
	left[0], right[n-1] = height[0], height[n-1]

	// find the left biggest value of every one
	for i := 1; i < len(height); i++ {
		left[i] = max(left[i-1], height[i])
	}

	// find the right biggest value of every one
	for i := n - 2; i >= 0; i-- {
		right[i] = max(right[i+1], height[i])
	}

	ret := 0
	for i := 1; i < n-1; i++ {
		// the water above height[i] is max(min(left[i], right[i])-height[i], 0)
		ret += max(min(left[i], right[i])-height[i], 0)
	}
	return ret
}

// We can optimize space by two pointers.
// Time: O(n)
// Space: O(1)
func trap2(height []int) int {
	n := len(height)
	if n <= 0 {
		return 0
	}
	ret := 0
	i, j := 1, n-2
	leftMax, rightMax := height[0], height[n-1]
	for i <= j {
		if leftMax < rightMax {
			ret += max(leftMax-height[i], 0)
			leftMax = max(leftMax, height[i])
			i++
		} else { // if leftMax > rightMax {
			ret += max(rightMax-height[j], 0)
			rightMax = max(rightMax, height[j])
			j--
		}
	}
	return ret
}
