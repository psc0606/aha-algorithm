package tree

import "testing"

func TestKthSmallest(t *testing.T) {
	root, _ := BuildTree([]interface{}{3, 1, 4, nil, 2})
	actual, expected := kthSmallest(root, 1), 1
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	root, _ = BuildTree([]interface{}{3, 1, 4, nil, 2})
	actual, expected = kthSmallest(root, 2), 2
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
}

func TestKthSmallestBfs(t *testing.T) {
	root, _ := BuildTree([]interface{}{3, 1, 4, nil, 2})
	actual, expected := kthSmallestBfs(root, 1), 1
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	root, _ = BuildTree([]interface{}{3, 1, 4, nil, 2})
	actual, expected = kthSmallestBfs(root, 2), 2
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
}
