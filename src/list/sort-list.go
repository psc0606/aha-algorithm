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

	// find mid of list
	var findMid func(head *ListNode) *ListNode
	findMid = func(head *ListNode) *ListNode {
		// at least two node
		dummy := &ListNode{Next: head}
		slow, fast := dummy, dummy
		for fast != nil && fast.Next != nil {
			slow = slow.Next
			fast = fast.Next.Next
		}
		return slow
	}

	// merge two sort list
	var mergeSortList func(head1 *ListNode, head2 *ListNode) *ListNode
	mergeSortList = func(head1 *ListNode, head2 *ListNode) *ListNode {
		if head1 == nil {
			return head2
		}
		if head2 == nil {
			return head1
		}
		var head, tail *ListNode
		for head1 != nil && head2 != nil {
			if head1.Val < head2.Val {
				if head == nil {
					head = head1
					tail = head1
				} else {
					tail.Next = head1
					tail = tail.Next
				}
				head1 = head1.Next
			} else {
				if head == nil {
					head = head2
					tail = head2
				} else {
					tail.Next = head2
					tail = tail.Next
				}
				head2 = head2.Next
			}
		}
		for ; head1 != nil; head1 = head1.Next {
			tail.Next = head1
			tail = tail.Next
		}
		for ; head2 != nil; head2 = head2.Next {
			tail.Next = head2
			tail = tail.Next
		}
		return head
	}

	mid := findMid(head)
	left, right := head, mid.Next
	mid.Next = nil
	left = sortList(left)
	right = sortList(right)
	return mergeSortList(left, right)
}

// Merge from down to top.
// Time: O(nlogn)
// Space: O(1)
// loop
func sortListII(head *ListNode) *ListNode {
	return nil
}
