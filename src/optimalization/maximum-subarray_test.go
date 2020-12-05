package optimalization

import "testing"

func TestMaxSubArray(t *testing.T) {
	max := maxSubArray([]int{8, -19, 5, -4, 20})
	expected := 21
	if max != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, max)
	}

	max = maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	expected = 6
	if max != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, max)
	}

	max = maxSubArray2([]int{8, -19, 5, -4, 20})
	expected = 21
	if max != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, max)
	}

	max = maxSubArray2([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	expected = 6
	if max != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, max)
	}
}
