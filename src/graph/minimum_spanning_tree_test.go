package graph

import "testing"

func TestMiniSpanningTree(t *testing.T) {
	cost := [][]int{{1, 3, 3}, {1, 2, 1}, {2, 3, 1}}
	actual, expected := miniSpanningTree(3, 3, cost), 2
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	cost = [][]int{{1, 2, 6}, {1, 3, 1}, {1, 4, 5}, {2, 3, 5}, {3, 4, 5}, {2, 5, 3}, {3, 5, 6}, {3, 6, 4}, {4, 6, 2}, {5, 6, 6}}
	actual, expected = miniSpanningTree(6, 10, cost), 15
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}
}
