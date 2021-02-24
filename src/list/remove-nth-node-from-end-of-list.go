package list

// https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/
// First calculate the length of list, then find the N-th node to be deleted.
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := 0
	p := head
	for ; p != nil; length++ {
		p = p.Next
	}
	if n > length || n < 1 {
		// invalid n, out range of list length
		return nil
	}
	if n == length {
		return head.Next
	}
	p = head
	var q *ListNode
	for i := 0; i < length-n; i++ {
		q = p
		p = p.Next
	}
	q.Next = p.Next
	return head
}

// Implemented by two pointer: fast and slow
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	fast, slow := dummy, dummy
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	if fast == nil {
		// that means n out range of list length, return nil.
		return nil
	}
	// move the fast pointer to the tail, then slow pointer points to the node before N-th node from end of list.
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}
