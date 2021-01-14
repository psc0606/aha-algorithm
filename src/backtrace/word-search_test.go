package backtrace

import "testing"

func TestExist(t *testing.T) {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	actual := exist(board, "ABCCED")
	expected := true
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual = exist(board, "SEE")
	expected = true
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual = exist(board, "ABCB")
	expected = false
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual = exist(board, "ASA")
	expected = true
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
}
