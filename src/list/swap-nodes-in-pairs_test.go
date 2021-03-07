package list

import "testing"

func TestSwapPairs2(t *testing.T) {
	head := BuildList([]int{1, 2, 3, 4})
	actual := swapPairs2(head)
	expected := BuildList([]int{2, 1, 4, 3})
	if !TestListEqual(expected, actual) {
		t.Errorf("expected is [%v], actual is [%v]", PrintList(expected), PrintList(actual))
	}

	head = BuildList([]int{1, 2, 3})
	actual = swapPairs2(head)
	expected = BuildList([]int{2, 1, 3})
	if !TestListEqual(expected, actual) {
		t.Errorf("expected is [%v], actual is [%v]", PrintList(expected), PrintList(actual))
	}
}

func TestSwapPairs3(t *testing.T) {
	head := BuildList([]int{1, 2, 3, 4})
	actual := swapPairs3(head)
	expected := BuildList([]int{2, 1, 4, 3})
	if !TestListEqual(expected, actual) {
		t.Errorf("expected is [%v], actual is [%v]", PrintList(expected), PrintList(actual))
	}

	head = BuildList([]int{1, 2, 3})
	actual = swapPairs3(head)
	expected = BuildList([]int{2, 1, 3})
	if !TestListEqual(expected, actual) {
		t.Errorf("expected is [%v], actual is [%v]", PrintList(expected), PrintList(actual))
	}
}

func TestSwapPairs4(t *testing.T) {
	head := BuildList([]int{1, 2, 3, 4})
	actual := swapPairs4(head)
	expected := BuildList([]int{2, 1, 4, 3})
	if !TestListEqual(expected, actual) {
		t.Errorf("expected is [%v], actual is [%v]", PrintList(expected), PrintList(actual))
	}

	head = BuildList([]int{1, 2, 3})
	actual = swapPairs4(head)
	expected = BuildList([]int{2, 1, 3})
	if !TestListEqual(expected, actual) {
		t.Errorf("expected is [%v], actual is [%v]", PrintList(expected), PrintList(actual))
	}
}
