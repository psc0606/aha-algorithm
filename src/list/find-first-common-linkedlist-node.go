package list

// https://leetcode-cn.com/problems/liang-ge-lian-biao-de-di-yi-ge-gong-gong-jie-dian-lcof/

// Time: O(n)
// Space: O(1)
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p := headA
	lengthA := 0
	for p != nil {
		p = p.Next
		lengthA++
	}
	q := headB
	lengthB := 0
	for q != nil {
		q = q.Next
		lengthB++
	}
	p = headA
	q = headB
	for i := 0; i < abs(lengthA, lengthB); i++ {
		if lengthA > lengthB {
			p = p.Next
		} else if lengthA < lengthB {
			q = q.Next
		}
	}

	for p != nil && q != nil {
		if p == q {
			break
		}
		p = p.Next
		q = q.Next
	}
	return p
}

func abs(a, b int) int {
	if a < b {
		return b - a
	} else {
		return a - b
	}
}

// A better solution, do it like cycle linked list.
//    |-length a-|
//       1 -> 2
//             \
//              --> 5 -> 6 -> 7
//             /    |-length c-|
//  1 -> 3 -> 4
//  |-length b-|
// satisfied condition:
// a + c + b = b + c + a
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	p := headA
	q := headB
	for p != q {
		if p != nil {
			p = p.Next
		} else {
			p = headB
		}
		if q != nil {
			q = q.Next
		} else {
			q = headA
		}
	}
	return p
}
