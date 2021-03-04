package list

import "testing"

func TestDetectCycle(t *testing.T) {
	head := &ListNode{
		Val: 1,
	}
	actual := detectCycle(head)
	var expected *ListNode
	if expected != actual {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	node1 := &ListNode{Val: 1, Next: nil}
	node2 := &ListNode{Val: 2}
	node1.Next = node2
	node2.Next = node1
	actual = detectCycle(node1)
	expected = node1
	if expected != actual {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	node1 = &ListNode{Val: 3}
	node2 = &ListNode{Val: 2}
	node3 := &ListNode{Val: 0}
	node4 := &ListNode{Val: -4}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2
	actual = detectCycle(node1)
	expected = node2
	if expected != actual {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
}

func TestDetectCycleWithO1(t *testing.T) {
	head := &ListNode{
		Val: 1,
	}
	actual := detectCycleWithO1(head)
	var expected *ListNode
	if expected != actual {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	node1 := &ListNode{Val: 1, Next: nil}
	node2 := &ListNode{Val: 2}
	node1.Next = node2
	node2.Next = node1
	actual = detectCycleWithO1(node1)
	expected = node1
	if expected != actual {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	node1 = &ListNode{Val: 3}
	node2 = &ListNode{Val: 2}
	node3 := &ListNode{Val: 0}
	node4 := &ListNode{Val: -4}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2
	actual = detectCycleWithO1(node1)
	expected = node2
	if expected != actual {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
}
