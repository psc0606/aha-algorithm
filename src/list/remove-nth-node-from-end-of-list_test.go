package list

import (
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	head := BuildList([]int{1, 2, 3, 4, 5})
	actual := removeNthFromEnd(head, 1)
	expected := BuildList([]int{1, 2, 3, 4})
	if !TestListEqual(expected, actual) {
		t.Errorf("expected is [%v], actual is [%v]", PrintList(expected), PrintList(actual))
	}

	head = BuildList([]int{1, 2, 3, 4, 5})
	actual = removeNthFromEnd(head, 2)
	expected = BuildList([]int{1, 2, 3, 5})
	if !TestListEqual(expected, actual) {
		t.Errorf("expected is [%v], actual is [%v]", PrintList(expected), PrintList(actual))
	}

	head = BuildList([]int{1, 2, 3, 4, 5})
	actual = removeNthFromEnd(head, 4)
	expected = BuildList([]int{1, 3, 4, 5})
	if !TestListEqual(expected, actual) {
		t.Errorf("expected is [%v], actual is [%v]", PrintList(expected), PrintList(actual))
	}

	head = BuildList([]int{1, 2, 3, 4, 5})
	actual = removeNthFromEnd(head, 5)
	expected = BuildList([]int{2, 3, 4, 5})
	if !TestListEqual(expected, actual) {
		t.Errorf("expected is [%v], actual is [%v]", PrintList(expected), PrintList(actual))
	}

	head = BuildList([]int{1, 2})
	actual = removeNthFromEnd(head, 2)
	expected = BuildList([]int{2})
	if !TestListEqual(expected, actual) {
		t.Errorf("expected is [%v], actual is [%v]", PrintList(expected), PrintList(actual))
	}

	head = BuildList([]int{1, 2})
	actual = removeNthFromEnd(head, 1)
	expected = BuildList([]int{1})
	if !TestListEqual(expected, actual) {
		t.Errorf("expected is [%v], actual is [%v]", PrintList(expected), PrintList(actual))
	}

	head = BuildList([]int{1})
	actual = removeNthFromEnd(head, 1)
	expected = BuildList([]int{})
	if !TestListEqual(expected, actual) {
		t.Errorf("expected is [%v], actual is [%v]", PrintList(expected), PrintList(actual))
	}
}

func TestRemoveNthFromEnd2(t *testing.T) {
	head := BuildList([]int{1, 2, 3, 4, 5})
	actual := removeNthFromEnd2(head, 1)
	expected := BuildList([]int{1, 2, 3, 4})
	if !TestListEqual(expected, actual) {
		t.Errorf("expected is [%v], actual is [%v]", PrintList(expected), PrintList(actual))
	}

	head = BuildList([]int{1, 2, 3, 4, 5})
	actual = removeNthFromEnd2(head, 2)
	expected = BuildList([]int{1, 2, 3, 5})
	if !TestListEqual(expected, actual) {
		t.Errorf("expected is [%v], actual is [%v]", PrintList(expected), PrintList(actual))
	}

	head = BuildList([]int{1, 2, 3, 4, 5})
	actual = removeNthFromEnd2(head, 4)
	expected = BuildList([]int{1, 3, 4, 5})
	if !TestListEqual(expected, actual) {
		t.Errorf("expected is [%v], actual is [%v]", PrintList(expected), PrintList(actual))
	}

	head = BuildList([]int{1, 2, 3, 4, 5})
	actual = removeNthFromEnd2(head, 5)
	expected = BuildList([]int{2, 3, 4, 5})
	if !TestListEqual(expected, actual) {
		t.Errorf("expected is [%v], actual is [%v]", PrintList(expected), PrintList(actual))
	}

	head = BuildList([]int{1, 2})
	actual = removeNthFromEnd2(head, 2)
	expected = BuildList([]int{2})
	if !TestListEqual(expected, actual) {
		t.Errorf("expected is [%v], actual is [%v]", PrintList(expected), PrintList(actual))
	}

	head = BuildList([]int{1, 2})
	actual = removeNthFromEnd2(head, 1)
	expected = BuildList([]int{1})
	if !TestListEqual(expected, actual) {
		t.Errorf("expected is [%v], actual is [%v]", PrintList(expected), PrintList(actual))
	}

	head = BuildList([]int{1})
	actual = removeNthFromEnd2(head, 1)
	expected = BuildList([]int{})
	if !TestListEqual(expected, actual) {
		t.Errorf("expected is [%v], actual is [%v]", PrintList(expected), PrintList(actual))
	}
}
