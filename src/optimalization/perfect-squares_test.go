package optimalization

import (
	"testing"
)

func TestNumSquares(t *testing.T) {
	actual := numSquares(23)
	expected := 4 // 9 + 9 + 4 + 1
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual = numSquares(16)
	expected = 1 // 16
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual = numSquares(1)
	expected = 1 // 1
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual = numSquares(2)
	expected = 2 // 1 + 1
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual = numSquares(3)
	expected = 3 // 1 + 1 + 1
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual = numSquares(4)
	expected = 1 // 4
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual = numSquares(12)
	expected = 3 // 4 + 4 + 4
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
}
