package tree

import (
	"fmt"
	"testing"
)

func TestPreorderTraversalBfs(t *testing.T) {
	root, _ := BuildTree([]interface{}{3, 1, 2})
	actual := preorderTraversalBfs(root)
	fmt.Print(actual)
}
