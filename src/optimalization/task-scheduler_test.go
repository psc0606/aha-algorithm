package optimalization

import "testing"

func TestLeastInterval(t *testing.T) {
	tasks := []byte{'A', 'A', 'A', 'B', 'B', 'B'}
	actual := leastInterval(tasks, 2)
	expected := 8
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	tasks = []byte{'A', 'A', 'A', 'B', 'B', 'B'}
	actual = leastInterval(tasks, 0)
	expected = 6
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	tasks = []byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'}
	actual = leastInterval(tasks, 2)
	expected = 16
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	tasks = []byte{'A', 'B', 'A'}
	actual = leastInterval(tasks, 2)
	expected = 4
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}
}
