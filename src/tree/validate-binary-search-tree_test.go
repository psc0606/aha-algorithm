package tree

import "testing"

func TestIsValidBST(t *testing.T) {
	root, _ := BuildTree([]interface{}{5, 1, 4, nil, nil, 3, 6})
	actual := isValidBST(root)
	expected := true
	if actual == expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	root, _ = BuildTree([]interface{}{2, 1, 3})
	actual = isValidBST(root)
	expected = true
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	root, _ = BuildTree([]interface{}{1, 2, 3})
	actual = isValidBST(root)
	expected = false
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	root, _ = BuildTree([]interface{}{5, 1, 6, nil, 8})
	actual = isValidBST(root)
	expected = false
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
}
