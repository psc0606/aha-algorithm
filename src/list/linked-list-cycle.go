package list

// https://leetcode-cn.com/problems/linked-list-cycle/
// Time: O(1)
// Space: O(1)
func hasCycle(head *ListNode) bool {
	slow := head
	fast := head
	for {
		if slow == nil || fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
		// If cycle exists, the fast pointer and slow pointer will collide within the cycle sooner or later.
		if slow == fast {
			return true
		}
	}
}

// More simple, but it will break the original list.
const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

func hasCycle2(head *ListNode) bool {
	p := head
	for p != nil {
		if p.Val == MaxInt {
			return true
		} else {
			p.Val = MaxInt
		}
		p = p.Next
	}
	return false
}
