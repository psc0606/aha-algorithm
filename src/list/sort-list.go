package list

// https://leetcode-cn.com/problems/sort-list/

// Merge sort from top to down
// Time: O(nlogn)
// Space: O(logn) by call stack
// TOP -> DOWN
// 3 -> 2 -> 5 -> 1
//        |
// 2 -> 3    1 -> 5
//        |
// 1 -> 2 -> 3 -> 5
// recursively
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	left, right := head, slow.Next
	slow.Next = nil
	left, right = sortList(left), sortList(right)
	return mergeTwoLists(left, right)
}

// Merge from down to top.
// Time: O(nlogn)
// Space: O(1)
// loop
func sortListII(head *ListNode) *ListNode {
	return nil
}
