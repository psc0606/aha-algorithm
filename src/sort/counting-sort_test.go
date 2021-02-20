package sort

import (
	"aha-algorithm/src/util"
	"fmt"
	"testing"
)

func TestCountingSort(t *testing.T) {
	comparator := func(a, b int) int {
		return a - b
	}

	arr := []int{5}
	CountingSort(arr, 10)
	if !IsIntSorted(arr, comparator) {
		t.Errorf("expected is [%v], actual is [%v]", true, false)
	}

	arr = []int{5, 3, 2, 1, 6, 8}
	fmt.Println(arr)
	CountingSort(arr, 10)
	fmt.Println(arr)
	if !IsIntSorted(arr, comparator) {
		t.Errorf("expected is [%v], actual is [%v]", true, false)
	}

	arr = []int{5, 3, 2, 1, 6, 8, 7, 11}
	fmt.Println(arr)
	CountingSort(arr, 11)
	fmt.Println(arr)
	if !IsIntSorted(arr, comparator) {
		t.Errorf("expected is [%v], actual is [%v]", true, false)
	}

	arr = util.RandArray(1000)
	fmt.Println(arr)
	CountingSort(arr, 1000)
	fmt.Println(arr)
	if !IsIntSorted(arr, comparator) {
		t.Errorf("expected is [%v], actual is [%v]", true, false)
	}
}
