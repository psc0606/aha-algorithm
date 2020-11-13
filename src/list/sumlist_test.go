package list

import "testing"

func TestAddTwoNumbers(t *testing.T) {
	var head1 *ListNode = nil
	var head2 *ListNode = nil
	resultHead := addTwoNumbers(head1, head2)
	if resultHead != nil {
		t.Error("expect nil")
	}

	// 7 -> nil
	// 5 -> nil
	head1 = &ListNode{
		Val:  7,
		Next: nil,
	}
	head2 = &ListNode{
		Val:  5,
		Next: nil,
	}
	resultHead = addTwoNumbers(head1, head2)
	if resultHead.Val != 1 {
		t.Error("expect nil")
	}
	if resultHead.Next.Val != 2 {
		t.Error("expect nil")
	}

	// 7 -> 1 -> nil
	// 5 -> nil
	head1 = &ListNode{
		Val: 7,
		Next: &ListNode{
			Val: 1,
		},
	}
	head2 = &ListNode{
		Val:  5,
		Next: nil,
	}
	resultHead = addTwoNumbers(head1, head2)
	if resultHead.Val != 7 {
		t.Error("expect 7")
	}
	if resultHead.Next.Val != 6 {
		t.Error("expect 6")
	}

	// 9 -> 9 -> 5 -> nil
	// 5 -> nil
	head1 = &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val:  5,
				Next: nil,
			},
		},
	}
	head2 = &ListNode{
		Val:  5,
		Next: nil,
	}
	resultHead = addTwoNumbers(head1, head2)
	if resultHead.Val != 1 {
		t.Error("expect 7")
	}
	if resultHead.Next.Val != 0 {
		t.Error("expect 6")
	}
}
