package tree

import (
	"testing"
)

func TestBuildTree(t *testing.T) {
	arr := []interface{}{5, 1, 4, nil, nil, 3, 6}
	root, err := BuildTree(arr)
	if root == nil {
		t.Error("expected not nil.", err)
	}

	// illegal array, it can not be built into a tree.
	//        5
	//     /     \
	//    1      nil
	//   / \    /   \
	// nil nil 3    6
	arr = []interface{}{5, 1, nil, nil, nil, 3, 6}
	root, err = BuildTree(arr)
	if root != nil {
		t.Error("expected not nil.", err)
	}

	// empty tree
	arr = []interface{}{nil}
	root, err = BuildTree(arr)
	if root != nil {
		t.Error("expected nil.", err)
	}
}
