package array

import "testing"

func TestTrap2(t *testing.T) {
	actual, expected := trap2([]int{2, 0, 2}), 2
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	actual, expected = trap2([]int{4, 2, 0, 3, 2, 5}), 9
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}
}
