package math

import (
	"math"
	"testing"
)

func TestMyPow2(t *testing.T) {
	actual, expected := myPow2(2.1, 3), 9.261
	if math.Abs(actual-expected) > 0.0000001 {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual, expected = myPow2(2.1, 4), 19.4481
	if math.Abs(actual-expected) > 0.0000001 {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual, expected = myPow2(2, 7), 128
	if math.Abs(actual-expected) > 0.0000001 {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual, expected = myPow2(2, -1), 0.5
	if math.Abs(actual-expected) > 0.0000001 {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
}

func TestMyPow3(t *testing.T) {
	actual, expected := myPow3(2.1, 3), 9.261
	if math.Abs(actual-expected) > 0.0000001 {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual, expected = myPow3(2.1, 4), 19.4481
	if math.Abs(actual-expected) > 0.0000001 {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual, expected = myPow3(2, 7), 128
	if math.Abs(actual-expected) > 0.0000001 {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual, expected = myPow3(2, -1), 0.5
	if math.Abs(actual-expected) > 0.0000001 {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
}
