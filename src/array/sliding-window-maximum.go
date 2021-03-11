package array

// https://leetcode-cn.com/problems/sliding-window-maximum/

// Use two stack, but time complexity is O(n^2)
func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	var stack, help []int
	var ret []int
	for i := 0; i < n; i++ {
		// stack store max value, the max value is always on the top.
		// the element in stack sort by: min(bottom) -> max(top)
		for len(stack) > 0 && stack[len(stack)-1] > nums[i] {
			ele := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			help = append(help, ele)
		}
		stack = append(stack, nums[i])
		for len(help) > 0 { // restore the popped elements
			stack = append(stack, help[len(help)-1])
			help = help[:len(help)-1]
		}

		if i-k >= 0 { // find the element then delete it
			for stack[len(stack)-1] != nums[i-k] {
				ele := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				help = append(help, ele)
			}
			stack = stack[:len(stack)-1]
			for len(help) > 0 { // restore
				stack = append(stack, help[len(help)-1])
				help = help[:len(help)-1]
			}
		}
		if len(stack) >= k {
			// top of stack is the window maximum
			ret = append(ret, stack[len(stack)-1])
		}
	}
	return ret
}

// The above solution problem: we do some useless stack push and pop, there is no need help stack.
func maxSlidingWindowOptimizeWithDequeue(nums []int, k int) []int {
	n := len(nums)
	var dequeue []int
	var ret []int
	for i := 0; i < n; i++ {
		// dequeue sort the elements in window from max(left) to min(right).
		// When we enqueue the current nums[i], we will discard all elements small than nums[i], because the current window's
		// maximum elements must in [max, ..., nums[i]].
		for len(dequeue) > 0 && dequeue[len(dequeue)-1] < nums[i] {
			dequeue = dequeue[:len(dequeue)-1] // pop end
		}
		dequeue = append(dequeue, nums[i])
		if i-k+1 >= 0 {
			// if dequeue[0] == nums[i-k]
			if i-k >= 0 && dequeue[0] == nums[i-k] {
				dequeue = dequeue[1:] // pop top
			}
			ret = append(ret, dequeue[0])
		}
	}
	return ret
}
