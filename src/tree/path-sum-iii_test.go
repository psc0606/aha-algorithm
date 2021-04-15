package tree

import "testing"

func TestPathSumIII3(t *testing.T) {
	arr := []interface{}{5, 4, 8, 11, 1, 4, 6, 7, 2, nil, nil, 5, 1}
	root, _ := BuildTree(arr)
	actual, expected := pathSumIII3(root, 17), 3
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	arr = []interface{}{1, 1, 1}
	root, _ = BuildTree(arr)
	actual, expected = pathSumIII3(root, 1), 3
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	arr = []interface{}{1, 1, 1}
	root, _ = BuildTree(arr)
	actual, expected = pathSumIII3(root, 2), 2
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	arr = []interface{}{0, 1, 1}
	root, _ = BuildTree(arr)
	actual, expected = pathSumIII3(root, 1), 4
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	arr = []interface{}{1}
	root, _ = BuildTree(arr)
	actual, expected = pathSumIII3(root, 1), 1
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	//arr = []interface{}{1, 0, 1, 1, 2, 0, -1, 0, 1, -1, 0, -1, 0, 1, 0}
	//root, _ = BuildTree(arr)
	//actual, expected = pathSumIII3(root, 2), 13
	//if actual != expected {
	//	t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	//}
}
