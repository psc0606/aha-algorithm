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

func TestReverseListRecursively(t *testing.T) {
	head := BuildList([]int{1})
	actual, expected := reverseListRecursively(head), BuildList([]int{1})
	if !TestListEqual(expected, actual) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	head = BuildList([]int{1, 2})
	actual, expected = reverseListRecursively(head), BuildList([]int{2, 1})
	if !TestListEqual(expected, actual) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	head = BuildList([]int{1, 2, 3})
	actual, expected = reverseListRecursively(head), BuildList([]int{3, 2, 1})
	if !TestListEqual(expected, actual) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
}
