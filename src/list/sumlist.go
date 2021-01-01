package list

// https://leetcode-cn.com/problems/sum-lists-lcci/
// https://leetcode-cn.com/problems/add-two-numbers/
// the natural order
// the input two list may not be same size

// example:
// 7 -> 1 -> 6
//     + 	    =  808
// 5 -> 9 -> 2
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	pl1 := reverseList(l1)
	pl2 := reverseList(l2)
	var head *ListNode = nil
	var tail *ListNode = nil
	// 进位
	var acc = 0
	for {
		var sum int
		if pl1 == nil && pl2 == nil {
			break
		} else if pl1 != nil && pl2 != nil {
			sum = pl1.Val + pl2.Val + acc
		} else if pl1 != nil {
			sum = pl1.Val + acc
		} else if pl2 != nil {
			sum = pl2.Val + acc
		}

		var node *ListNode = nil
		if sum >= 10 {
			node = &ListNode{
				Val: sum % 10,
			}
			acc = 1
		} else {
			node = &ListNode{
				Val: sum,
			}
			acc = 0
		}

		if head == nil {
			head = node
			tail = head
		} else {
			tail.Next = node
			tail = node
		}
		if pl1 != nil {
			pl1 = pl1.Next
		}
		if pl2 != nil {
			pl2 = pl2.Next
		}
	}

	// at last check acc
	if acc == 1 {
		tail.Next = &ListNode{
			Val: 1,
		}
	}
	return reverseList(head)
}
