package list

// https://leetcode-cn.com/problems/insertion-sort-list/
// 4->2->1->3
// last->4, next->2
// Another better solution to swap list node is only swap the value, not the whole node.
func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	last := head
	next := head.Next
	for next != nil {
		if last.Val <= next.Val {
			last = next
			next = next.Next
			continue
		}
		// remove the node pointed by next
		last.Next = next.Next

		find := head
		var pre *ListNode
		for find.Val < next.Val {
			pre = find
			find = find.Next
		}
		// means head is small than next.Val
		if pre == nil {
			next.Next = find
			head = next
		} else {
			pre.Next = next
			next.Next = find
		}

		// recover to the position
		next = last.Next
	}
	return head
}
