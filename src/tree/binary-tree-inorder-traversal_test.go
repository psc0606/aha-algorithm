package tree

import (
	"reflect"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	root, _ := BuildTree([]interface{}{3, 9, 20, nil, nil, 15, 7})
	actual := inorderTraversal(root)
	expected := []int{9, 3, 15, 20, 7}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	root, _ = BuildTree([]interface{}{3})
	actual = inorderTraversal(root)
	expected = []int{3}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	root, _ = BuildTree([]interface{}{3, 9, nil})
	actual = inorderTraversal(root)
	expected = []int{9, 3}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
}

func TestInorderTraversalRecursively(t *testing.T) {
	root, _ := BuildTree([]interface{}{3, 9, 20, nil, nil, 15, 7})
	actual := inorderTraversalRecursively(root)
	expected := []int{9, 3, 15, 20, 7}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	root, _ = BuildTree([]interface{}{3})
	actual = inorderTraversalRecursively(root)
	expected = []int{3}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	root, _ = BuildTree([]interface{}{3, 9, nil})
	actual = inorderTraversalRecursively(root)
	expected = []int{9, 3}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
}
