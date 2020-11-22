package list

// https://leetcode-cn.com/problems/merge-two-sorted-lists/
// https://leetcode-cn.com/problems/he-bing-liang-ge-pai-xu-de-lian-biao-lcof/
// Time: O(m + n)
// Space: O(1)
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	head := l1
	if l2.Val < l1.Val {
		head = l2
		l2 = l2.Next
	} else {
		l1 = l1.Next
	}
	tail := head
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			tail.Next = l1
			l1 = l1.Next
		} else {
			tail.Next = l2
			l2 = l2.Next
		}
		tail = tail.Next
	}
	for l1 != nil {
		tail.Next = l1
		tail = tail.Next
		l1 = l1.Next
	}
	for l2 != nil {
		tail.Next = l2
		tail = tail.Next
		l2 = l2.Next
	}
	return head
}
