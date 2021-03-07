package list

// https://leetcode-cn.com/problems/swap-nodes-in-pairs/
// Time: O(n)
// Space: O(1), but it not really swap the nodes, only replace the value.
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	p, q := head, head.Next
	for p != nil && q != nil {
		p.Val, q.Val = q.Val, p.Val
		p = q.Next
		if p != nil {
			q = p.Next
		}
	}
	return head
}

// If we must to swap both neighbour nodes, how to do it?
func swapPairs2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	p, q := head, head.Next
	var last *ListNode // record last already swapped pointer.
	for p != nil && q != nil {
		p.Next = q.Next
		q.Next = p
		if last != nil {
			last.Next = q
		} else {
			head = q
		}
		last = p

		// after swap
		p = p.Next
		if p != nil {
			q = p.Next
		}
	}
	return head
}

// Simplify code with dummy head.
func swapPairs3(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{Next: head}
	p := dummy // record the last swapped pointer
	for p.Next != nil && p.Next.Next != nil {
		node1 := p.Next
		node2 := p.Next.Next

		// swap
		node1.Next = node2.Next
		node2.Next = node1
		p.Next = node2

		p = node1 // tail of last swapped pointer
	}
	return dummy.Next
}

// Recursively
func swapPairs4(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	head.Next = swapPairs4(newHead.Next)
	newHead.Next = head
	return newHead
}
