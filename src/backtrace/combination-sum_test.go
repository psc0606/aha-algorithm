package backtrace

import (
	"fmt"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	arr := []int{2, 3, 6, 7}
	actual := combinationSum(arr, 7)
	fmt.Println(actual)
}
