package tree

import (
	"fmt"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	root, _ := BuildTree([]interface{}{3, 9, 20, nil, nil, 15, 7})
	actual := levelOrder(root)
	fmt.Print(actual)

	root, _ = BuildTree([]interface{}{3, 9, 20, nil, 8, 15, 7})
	actual = levelOrder(root)
	fmt.Print(actual)

	root, _ = BuildTree([]interface{}{3, 9, 20, nil, nil, 15, 7})
	actual = levelOrder2(root)
	fmt.Print(actual)

	root, _ = BuildTree([]interface{}{3, 9, 20, nil, 8, 15, 7})
	actual = levelOrder2(root)
	fmt.Print(actual)
}
