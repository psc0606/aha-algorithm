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
func sortListII(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// n is the length of list
	n, p := 0, head
	for p != nil {
		n++
		p = p.Next
	}

	dummy := &ListNode{Next: head}
	for step := 1; step <= n; step *= 2 {
		prev, curr := dummy, dummy
		for curr != nil { // every sub-list
			leftHead := prev.Next
			rightHead := split(prev.Next, step)
			curr = split(rightHead, step)
			prev.Next = mergeTwoLists(leftHead, rightHead)
			for prev.Next != nil {
				prev = prev.Next
			}
			prev.Next = curr
		}
	}
	return dummy.Next
}

// return right head of split list, the left tail pointer to null.
func split(head *ListNode, step int) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{Next: head}
	curr := dummy
	for i := 0; i < step && curr.Next != nil; i++ {
		curr = curr.Next
	}
	right := curr.Next
	curr.Next = nil
	return right
}
