package tree

import (
	"fmt"
	"testing"
)

func TestPathSum(t *testing.T) {
	//       1
	//    2    3
	//  4  5
	//
	arr := []interface{}{1, 2, 3, 4, 5}
	root, _ := BuildTree(arr)
	actual := pathSum(root, 4)
	fmt.Println(actual)
}
