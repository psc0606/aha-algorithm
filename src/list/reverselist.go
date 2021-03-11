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

func reverseListRecursively(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	// return sublist head only
	// the tail of sublist is head.Next
	p := reverseListRecursively(head.Next)
	// If sublist is nil, set current node to the head.
	if p == nil {
		p = head
	}
	if head.Next != nil {
		head.Next.Next = head
		head.Next = nil
	}
	return p
}
