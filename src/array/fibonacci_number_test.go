package array

import "testing"

func TestFibonacciNumber(t *testing.T) {
	actual := fibonacciNumberRecursively(0)
	expected := 0
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	actual = fibonacciNumberRecursively(1)
	expected = 1
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	actual = fibonacciNumberRecursively(2)
	expected = 1
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	actual = fibonacciNumberRecursively(3)
	expected = 2
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	actual = fibonacciNumberRecursively(4)
	expected = 3
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	actual = fibonacciNumberRecursively(5)
	expected = 5
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	actual = fibonacciNumberRecursively(6)
	expected = 8
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	actual = fibonacciNumberRecursively(7)
	expected = 13
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	// dp
	actual = fibonacciNumberDp(0)
	expected = 0
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	actual = fibonacciNumberDp(1)
	expected = 1
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	actual = fibonacciNumberDp(2)
	expected = 1
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	actual = fibonacciNumberDp(3)
	expected = 2
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	actual = fibonacciNumberDp(4)
	expected = 3
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	actual = fibonacciNumberDp(5)
	expected = 5
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	actual = fibonacciNumberDp(6)
	expected = 8
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	actual = fibonacciNumberDp(7)
	expected = 13
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	actual = fibonacciNumberDp(50)
	expected = 12586269025
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	// dp2 space
	actual = fibonacciNumberDpOptimizeSpace(0)
	expected = 0
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	actual = fibonacciNumberDpOptimizeSpace(1)
	expected = 1
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	actual = fibonacciNumberDpOptimizeSpace(2)
	expected = 1
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	actual = fibonacciNumberDpOptimizeSpace(3)
	expected = 2
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	actual = fibonacciNumberDpOptimizeSpace(4)
	expected = 3
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	actual = fibonacciNumberDpOptimizeSpace(5)
	expected = 5
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	actual = fibonacciNumberDpOptimizeSpace(6)
	expected = 8
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	actual = fibonacciNumberDpOptimizeSpace(7)
	expected = 13
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	actual = fibonacciNumberDpOptimizeSpace(50)
	expected = 12586269025
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}
}
