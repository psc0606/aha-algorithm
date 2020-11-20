package optimalization

import "testing"

func TestMassage(t *testing.T) {
	max := massage2([]int{1, 2, 3, 1})
	expected := 4
	if max != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, max)
	}

	max = massage2([]int{2, 1})
	expected = 2
	if max != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, max)
	}

	max = massage2([]int{2, 1, 4, 5, 3, 1, 1, 3})
	expected = 12
	if max != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, max)
	}
}
