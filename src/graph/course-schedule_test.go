package graph

import "testing"

func TestCanFinish(t *testing.T) {
	prerequisites := [][]int{{1, 0}, {2, 1}}
	actual, expected := canFinish(3, prerequisites), true
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	prerequisites = [][]int{{1, 0}}
	actual, expected = canFinish(2, prerequisites), true
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	prerequisites = [][]int{{1, 0}, {0, 1}}
	actual, expected = canFinish(2, prerequisites), false
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	prerequisites = [][]int{{1, 0}, {2, 1}, {0, 2}}
	actual, expected = canFinish(3, prerequisites), false
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	prerequisites = [][]int{{1, 0}, {2, 1}, {0, 2}, {3, 1}}
	actual, expected = canFinish(4, prerequisites), false
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
}

func TestCanFinishByDfs(t *testing.T) {
	prerequisites := [][]int{{1, 0}, {2, 1}}
	actual, expected := canFinishByDfs(3, prerequisites), true
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	prerequisites = [][]int{{1, 0}}
	actual, expected = canFinishByDfs(2, prerequisites), true
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	prerequisites = [][]int{{1, 0}, {0, 1}}
	actual, expected = canFinishByDfs(2, prerequisites), false
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	prerequisites = [][]int{{1, 0}, {2, 1}, {0, 2}}
	actual, expected = canFinishByDfs(3, prerequisites), false
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	prerequisites = [][]int{{1, 0}, {2, 1}, {0, 2}, {3, 1}}
	actual, expected = canFinishByDfs(4, prerequisites), false
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	prerequisites = [][]int{{2, 0}, {1, 0}, {3, 1}, {3, 2}, {1, 3}}
	actual, expected = canFinishByDfs(4, prerequisites), false
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
}
