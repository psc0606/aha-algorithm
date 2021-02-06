package optimalization

import "testing"

func TestUniquePaths(t *testing.T) {
	actual := uniquePaths2(3, 2)
	expected := 3
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	actual = uniquePaths2(3, 7)
	expected = 28
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	actual = uniquePaths2(59, 9)
	expected = 5743572120
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}
}
