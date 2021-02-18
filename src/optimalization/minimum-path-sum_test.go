package optimalization

import "testing"

func TestMinPathSum(t *testing.T) {
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	actual := minPathSum(grid)
	expected := 7
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	grid = [][]int{{1, 3, 7}, {1, 5, 1}, {4, 2, 1}}
	actual = minPathSum(grid)
	expected = 9
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}
}
