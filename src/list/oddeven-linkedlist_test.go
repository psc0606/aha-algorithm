package list

import "testing"

func TestOddEvenList(t *testing.T) {
	head := &ListNode{
		Val:  1,
		Next: nil,
	}
	newHead := oddEvenList(head)
	if head != newHead {
		t.Errorf("expected same")
	}

	head = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}
	newHead = oddEvenList(head)
	if head != newHead {
		t.Errorf("expected same")
	}

	head = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	newHead = oddEvenList(head)
	if head != newHead {
		t.Errorf("expected same")
	}

	head = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}
	newHead = oddEvenList(head)
	if head != newHead {
		t.Errorf("expected same")
	}
}
