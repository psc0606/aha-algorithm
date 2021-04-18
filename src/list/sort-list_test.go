package list

import (
	"testing"
)

func TestSortList(t *testing.T) {
	head := BuildList([]int{5, 2, 1, 4})
	actual := sortList(head)
	expected := BuildList([]int{1, 2, 4, 5})
	if !TestListEqual(expected, actual) {
		t.Errorf("expected is [%v], actual is [%v]", PrintList(expected), PrintList(actual))
	}

	head = BuildList([]int{5})
	actual = sortList(head)
	expected = BuildList([]int{5})
	if !TestListEqual(expected, actual) {
		t.Errorf("expected is [%v], actual is [%v]", PrintList(expected), PrintList(actual))
	}

	head = BuildList([]int{5, 2})
	actual = sortList(head)
	expected = BuildList([]int{2, 5})
	if !TestListEqual(expected, actual) {
		t.Errorf("expected is [%v], actual is [%v]", PrintList(expected), PrintList(actual))
	}

	head = BuildList([]int{5, 2, 1})
	actual = sortList(head)
	expected = BuildList([]int{1, 2, 5})
	if !TestListEqual(expected, actual) {
		t.Errorf("expected is [%v], actual is [%v]", PrintList(expected), PrintList(actual))
	}

	head = BuildList([]int{3, 5, 2, 1})
	actual = sortList(head)
	expected = BuildList([]int{1, 2, 3, 5})
	if !TestListEqual(expected, actual) {
		t.Errorf("expected is [%v], actual is [%v]", PrintList(expected), PrintList(actual))
	}
}

func TestSortListII(t *testing.T) {
	head := BuildList([]int{5, 2, 1, 4})
	actual := sortListII(head)
	expected := BuildList([]int{1, 2, 4, 5})
	if !TestListEqual(expected, actual) {
		t.Errorf("expected is [%v], actual is [%v]", PrintList(expected), PrintList(actual))
	}

	head = BuildList([]int{5})
	actual = sortListII(head)
	expected = BuildList([]int{5})
	if !TestListEqual(expected, actual) {
		t.Errorf("expected is [%v], actual is [%v]", PrintList(expected), PrintList(actual))
	}

	head = BuildList([]int{5, 2})
	actual = sortListII(head)
	expected = BuildList([]int{2, 5})
	if !TestListEqual(expected, actual) {
		t.Errorf("expected is [%v], actual is [%v]", PrintList(expected), PrintList(actual))
	}

	head = BuildList([]int{5, 2, 1})
	actual = sortListII(head)
	expected = BuildList([]int{1, 2, 5})
	if !TestListEqual(expected, actual) {
		t.Errorf("expected is [%v], actual is [%v]", PrintList(expected), PrintList(actual))
	}

	head = BuildList([]int{3, 5, 2, 1})
	actual = sortListII(head)
	expected = BuildList([]int{1, 2, 3, 5})
	if !TestListEqual(expected, actual) {
		t.Errorf("expected is [%v], actual is [%v]", PrintList(expected), PrintList(actual))
	}
}
