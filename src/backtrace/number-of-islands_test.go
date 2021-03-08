package backtrace

import "testing"

func TestNumIslands(t *testing.T) {
	grids := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	actual := numIslands(grids)
	expected := 3
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	grids = [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	actual = numIslands(grids)
	expected = 1
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	grids = [][]byte{
		{'0', '0', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	actual = numIslands(grids)
	expected = 0
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	grids = [][]byte{
		{'0', '0', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	actual = numIslands(grids)
	expected = 1
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	grids = [][]byte{
		{'0', '0', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '0'},
	}
	actual = numIslands(grids)
	expected = 2
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	grids = [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	actual = numIslands(grids)
	expected = 1
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
}

func TestNumIslandsByDisjointSet(t *testing.T) {
	grids := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	actual := numIslandsByDisjointSet(grids)
	expected := 3
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	grids = [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	actual = numIslandsByDisjointSet(grids)
	expected = 1
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	grids = [][]byte{
		{'0', '0', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	actual = numIslandsByDisjointSet(grids)
	expected = 0
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	grids = [][]byte{
		{'0', '0', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	actual = numIslandsByDisjointSet(grids)
	expected = 1
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	grids = [][]byte{
		{'0', '0', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '0'},
	}
	actual = numIslandsByDisjointSet(grids)
	expected = 2
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	grids = [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	actual = numIslandsByDisjointSet(grids)
	expected = 1
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
}
