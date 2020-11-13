package list

import "testing"

func TestReverseList(t *testing.T) {
	tail := reverseList(nil)
	if tail != nil {
		t.Error("expected nil")
	}

	head := &ListNode{
		Val:  0,
		Next: nil,
	}
	tail = reverseList(head)
	if tail != head {
		t.Errorf("expected head[%p] = tail[%p]", head, tail)
	}

	head = &ListNode{
		Val: 0,
		Next: &ListNode{
			Val:  1,
			Next: nil,
		},
	}
	tail = reverseList(head)
	if tail.Val != 1 {
		t.Errorf("expected tail [%d]=[%d]", tail.Val, 1)
	}

	head = &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}
	tail = reverseList(head)
	if tail.Val != 2 {
		t.Errorf("expected tail [%d]=[%d]", tail.Val, 2)
	}
}
