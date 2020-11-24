package list

// https://leetcode-cn.com/problems/sort-list/
// Time: O(nlogn)
// Space: O(1)
// merge sort
func sortList(head *ListNode) *ListNode {
	step := 1
	p := head

	length := 0
	for p != nil {
		length++
		p = p.Next
	}

	for step < length/2 {
		p = head
		i := 0
		for i < step {

		}
		step *= 2
	}
	return nil
}
