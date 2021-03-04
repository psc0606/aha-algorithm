package list

// https://leetcode-cn.com/problems/linked-list-cycle-ii/
// Solution1: Use a hashtable to record all visited node.
// Time: O(n)
// Space: O(n)
func detectCycle(head *ListNode) *ListNode {
	ht := make(map[*ListNode]bool)
	p := head
	for p != nil {
		_, ok := ht[p]
		if ok {
			return p
		}
		ht[p] = true
		p = p.Next
	}
	return nil
}

// https://leetcode-cn.com/problems/linked-list-cycle-ii/solution/huan-xing-lian-biao-ii-by-leetcode-solution/
// A faster pointer and a slow pointer meet with each other. Proof is important.
// Time: O(n)
// Space: O(1)
func detectCycleWithO1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	fast, slow := head, head

	// find the first meet point
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}
	// no circle at all
	if fast == nil || fast.Next == nil {
		return nil
	}
	fast = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}
