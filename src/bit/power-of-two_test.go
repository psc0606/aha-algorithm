package bit

import "testing"

func TestIsPowerOfTwo(t *testing.T) {
	actual, expected := isPowerOfTwo(2), true
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual, expected = isPowerOfTwo(1), true
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
}
