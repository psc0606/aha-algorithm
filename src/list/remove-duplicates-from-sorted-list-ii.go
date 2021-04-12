package list

// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/
func deleteDuplicatesII(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{Val: 0, Next: head}
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		// find duplicate to be deleted.
		if cur.Next.Val == cur.Next.Next.Val {
			x := cur.Next.Val // keep the duplicate value in meed.
			// remove all
			p := cur.Next
			for p != nil && p.Val == x {
				cur.Next = p.Next
				p = p.Next
			}
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}
