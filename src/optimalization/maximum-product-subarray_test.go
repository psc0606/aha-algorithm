package optimalization

import "testing"

func TestMaxProduct(t *testing.T) {
	arr := []int{2, 3, -2, 4}
	actual := maxProduct(arr)
	expected := 6
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	arr = []int{2}
	actual = maxProduct(arr)
	expected = 2
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	arr = []int{-2}
	actual = maxProduct(arr)
	expected = -2
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	arr = []int{-2, 2, -3}
	actual = maxProduct(arr)
	expected = 12
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	arr = []int{-1, -2, -9, -6}
	actual = maxProduct(arr)
	expected = 108
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	arr = []int{2, -5, -2, -4, 3}
	actual = maxProduct(arr)
	expected = 24
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
}
