package array

import (
	"fmt"
	"testing"
)

func TestMergeIntervals(t *testing.T) {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	actual := mergeIntervals(intervals)
	fmt.Println(actual)
}
