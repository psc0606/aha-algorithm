package list

// https://leetcode-cn.com/problems/reverse-linked-list/
// in: 1->2->3->4->5->NULL
// out: 5->4->3->2->1->NULL

// 1 -> 2 -> 3 -> 4
// |   |
// p   q
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var p *ListNode = nil
	q := head
	for {
		// next pointer
		if q == nil {
			return p
		}
		next := q.Next
		q.Next = p
		p = q
		q = next
	}
}
