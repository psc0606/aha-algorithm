package tree

import (
	"testing"
)

func TestLowestCommonAncestor(t *testing.T) {
	root, _ := BuildTree([]interface{}{3, 5, 1, 6, 2, 0, 8})
	actual := lowestCommonAncestor(root, root.Left.Left, root.Left.Right)
	expected := root.Left
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
}

func TestLowestCommonAncestorRecursively(t *testing.T) {
	root, _ := BuildTree([]interface{}{3, 5, 1, 6, 2, 0, 8})
	actual := lowestCommonAncestorRecursively(root, root.Left.Left, root.Left.Right)
	expected := root.Left
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
}
