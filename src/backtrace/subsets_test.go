package backtrace

import (
	"fmt"
	"testing"
)

func TestSubsets(t *testing.T) {
	nums := []int{1, 2, 3}
	actual := subsets(nums)
	fmt.Println(actual)
}

func TestSubsetsWithBits(t *testing.T) {
	nums := []int{1, 2, 3}
	actual := subsetsWithBits(nums)
	fmt.Println(actual)
}
