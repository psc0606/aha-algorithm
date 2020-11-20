package list

import (
	"fmt"
	"strconv"
)

// create list by slice for test
func BuildList(arr []int) *ListNode {
	var head *ListNode
	p := head
	for _, ele := range arr {
		node := &ListNode{
			Val:  ele,
			Next: nil,
		}
		if p == nil {
			head = node
			p = head
		} else {
			p.Next = node
			p = p.Next
		}
	}
	return head
}

// test two list whether equal or not
func TestListEqual(head1 *ListNode, head2 *ListNode) bool {
	if head1 == head2 {
		return true
	}
	p := head1
	q := head2
	for p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}
		p = p.Next
		q = q.Next
	}
	if p == q {
		return true
	} else {
		return false
	}
}

// print list like 1->2->3->4
func PrintList(head *ListNode) {
	p := head
	var slice string
	for p != nil {
		if len(slice) == 0 {
			slice += strconv.Itoa(p.Val)
		} else {
			slice += "->" + strconv.Itoa(p.Val)
		}
		p = p.Next
	}
	fmt.Println(slice)
}
