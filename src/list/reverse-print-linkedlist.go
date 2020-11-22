package list

import "aha-algorithm/src/util"

// https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/
func reversePrint(head *ListNode) []int {
	var stack []int
	p := head
	for p != nil {
		stack = append(stack, p.Val)
		p = p.Next
	}
	res := make([]int, len(stack))
	j := 0
	for i := len(stack) - 1; i >= 0; i-- {
		res[j] = stack[i]
		j++
	}
	return res
}

// update in place
func reversePrint2(head *ListNode) []int {
	var stack []int
	p := head
	for p != nil {
		stack = append(stack, p.Val)
		p = p.Next
	}
	return util.SliceReverse(stack)
}
