package list

// https://leetcode-cn.com/problems/palindrome-linked-list/
// Time: O(n)
// Space: O(1)
// We can also use a fast pointer and a slow pointer find the middle point.
func isPalindrome(head *ListNode) bool {
	count := 0
	p := head
	for p != nil {
		p = p.Next
		count++
	}
	if count < 2 {
		return true
	}
	// half of the list
	semi := (count + 1) / 2
	head2 := head
	for i := 0; i < semi; i++ {
		head2 = head2.Next
	}
	head2 = _reverseList(head2)
	for head != nil && head2 != nil {
		if head.Val != head2.Val {
			return false
		}
		head = head.Next
		head2 = head2.Next
	}
	return true
}

func _reverseList(head *ListNode) *ListNode {
	var last *ListNode = nil
	curr := head
	for curr != nil {
		tmp := curr.Next
		curr.Next = last
		last = curr
		curr = tmp
	}
	return last
}
