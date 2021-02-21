package array

import "testing"

func TestSearch(t *testing.T) {
	arr := []int{4, 5, 6, 7, 0, 1, 2}
	actual := search(arr, 0)
	expected := 4
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	arr = []int{0, 1, 2, 4, 5, 6, 7}
	actual = search(arr, 0)
	expected = 0
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	arr = []int{1, 2, 4, 5, 6, 7, 0}
	actual = search(arr, 0)
	expected = 6
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	arr = []int{2, 4, 5, 6, 7, 0, 1}
	actual = search(arr, 0)
	expected = 5
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	arr = []int{5, 1, 3}
	actual = search(arr, 3)
	expected = 2
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}
}
