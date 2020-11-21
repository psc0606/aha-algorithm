package array

import (
	"aha-algorithm/src/util"
	"testing"
)

func TestArrayRankTransform(t *testing.T) {
	rank := arrayRankTransform([]int{40, 10, 20, 30})
	expected := []int{4, 1, 2, 3}
	if !util.SliceEqual(rank, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, rank)
	}

	rank = arrayRankTransform([]int{100, 100, 100})
	expected = []int{1, 1, 1}
	if !util.SliceEqual(rank, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, rank)
	}

	rank = arrayRankTransform([]int{37, 12, 28, 9, 100, 56, 80, 5, 12})
	expected = []int{5, 3, 4, 2, 8, 6, 7, 1, 3}
	if !util.SliceEqual(rank, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, rank)
	}
}
