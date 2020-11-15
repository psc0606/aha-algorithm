package list

// https://leetcode-cn.com/problems/odd-even-linked-list/
// Time: O(n)
// Space: O(1)
// This is my self implementation, but it is a bit more difficult to understand.
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	// point to odd node
	p := head
	var oddTail *ListNode

	// point to first even node
	evenHead := head.Next
	// pointer to even list evenTail
	var evenTail *ListNode
	for {
		// point to next even node
		if p == nil {
			oddTail.Next = evenHead
			return head
		}
		q := p.Next
		if q == nil {
			// odd evenTail pointer to he even head
			p.Next = evenHead
			return head
		}
		oddTail = p
		p.Next = q.Next
		p = p.Next

		if evenTail == nil {
			evenTail = q
		} else {
			evenTail.Next = q
			evenTail = evenTail.Next
		}
		evenTail.Next = nil
	}
}

// other people's better implementation with only few codes
func oddEvenListOptimize(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// case 1:
	// 1->2->3->4->NULL
	// |  |
	// odd  even
	//       odd  even
	// case 2:
	// 1->2->3->NULL
	// |  |
	// odd  even
	//       odd
	odd, even, eventHead := head, head.Next, head.Next
	for odd.Next != nil && even.Next != nil {
		odd.Next = even.Next
		even.Next = even.Next.Next
		odd = odd.Next
		even = even.Next
	}
	odd.Next = eventHead
	return head
}
