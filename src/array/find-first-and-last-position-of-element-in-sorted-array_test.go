package array

import (
	"fmt"
	"testing"
)

func TestSearchRange(t *testing.T) {
	res := searchRange([]int{8}, 8)
	fmt.Println(res)

	res = searchRange([]int{8}, 7)
	fmt.Println(res)

	res = searchRange([]int{8}, 9)
	fmt.Println(res)

	res = searchRange([]int{8, 8}, 8)
	fmt.Println(res)

	res = searchRange([]int{8, 9}, 7)
	fmt.Println(res)

	res = searchRange([]int{5, 7, 7, 8, 8, 10}, 8)
	fmt.Println(res)
}
