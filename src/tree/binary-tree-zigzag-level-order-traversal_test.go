package tree

import (
	"fmt"
	"testing"
)

func TestZigzagLevelOrder(t *testing.T) {
	root, _ := BuildTree([]interface{}{1, 2, 3, 4, 5})
	actual := zigzagLevelOrder(root)
	fmt.Println(actual)
}
