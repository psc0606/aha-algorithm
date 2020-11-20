package optimalization

// https://leetcode-cn.com/problems/the-masseuse-lcci/
// Time: O(n)
// Space: O(n)
// solved by greed algorithm
func massage(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	op := make([]int, len(nums))
	op[0] = nums[0]
	op[1] = max(nums[1], op[0])
	for i := 2; i < len(nums); i++ {
		op[i] = max(op[i-2]+nums[i], op[i-1])
	}
	return op[len(nums)-1]
}

// Space: O(1)
func massage2(nums []int) int {
	op1 := 0
	op2 := 0
	for i := 0; i < len(nums); i++ {
		tmp := op2
		op2 = max(op1+nums[i], op2)
		op1 = tmp
	}
	return op2
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
