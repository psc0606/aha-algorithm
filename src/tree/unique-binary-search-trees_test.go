package tree

import "testing"

func TestNumTrees(t *testing.T) {
	actual := numTrees(0)
	expected := 0
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	actual = numTrees(1)
	expected = 1
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	actual = numTrees(2)
	expected = 2
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	actual = numTrees(3)
	expected = 5
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	actual = numTrees(4)
	// when 1 is the root, 5
	// when 2 is the root, 2
	expected = 5 + 2 + 2 + 5
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}
}
