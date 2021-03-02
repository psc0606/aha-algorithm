package array

import (
	"fmt"
	"testing"
)

func TestThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	actual := threeSum(nums)
	fmt.Println(actual)

	nums = []int{}
	actual = threeSum(nums)
	fmt.Println(actual)

	nums = []int{0}
	actual = threeSum(nums)
	fmt.Println(actual)

	nums = []int{0, 0}
	actual = threeSum(nums)
	fmt.Println(actual)

	nums = []int{0, 0, 0}
	actual = threeSum(nums)
	fmt.Println(actual)

	// must not generate duplicate tuple
	nums = []int{0, 0, 0, 0}
	actual = threeSum(nums)
	fmt.Println(actual) // [[0 0 0]]

	// must not generate duplicate tuple
	nums = []int{0, 0, 0, 0, 0}
	actual = threeSum(nums)
	fmt.Println(actual) // [[0 0 0]]

	// must not generate duplicate tuple
	nums = []int{1, -1, 0, 0, 0}
	actual = threeSum(nums)
	fmt.Println(actual) // [[-1 0 1] [0 0 0]]

	// must not generate duplicate tuple
	nums = []int{-1, -1, 2, 2, 2}
	actual = threeSum(nums)
	fmt.Println(actual) // [[-1 -1 2]]
}
