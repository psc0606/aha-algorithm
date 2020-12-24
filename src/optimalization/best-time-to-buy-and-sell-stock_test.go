package optimalization

import "testing"

func TestMaxProfit(t *testing.T) {
	max := maxProfit([]int{7, 1, 5, 3, 6, 4})
	expected := 5
	if max != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, max)
	}

	max = maxProfit([]int{7, 6, 4, 3, 1})
	expected = 0
	if max != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, max)
	}
}
