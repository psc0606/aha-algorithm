package backtrace

import (
	"fmt"
	"testing"
)

func TestPermute(t *testing.T) {
	nums := []int{1}
	actual := permute(nums)
	fmt.Println(actual)

	nums = []int{1, 2}
	actual = permute(nums)
	fmt.Println(actual)

	nums = []int{1, 2, 3}
	actual = permute(nums)
	fmt.Println(actual)
}
