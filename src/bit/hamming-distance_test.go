package bit

import "testing"

func TestHammingDistance(t *testing.T) {
	count := numberOfBit1(1)
	expected := 1
	if count != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, count)
	}

	count = numberOfBit1(3)
	expected = 2
	if count != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, count)
	}

	count = numberOfBit1(0)
	expected = 0
	if count != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, count)
	}

	count = numberOfBit1(8)
	expected = 1
	if count != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, count)
	}
}
