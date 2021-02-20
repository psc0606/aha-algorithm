package sort

import (
	"math/rand"
	"time"
)

func QuickSort(arr []interface{}, comparator func(a interface{}, b interface{}) int) {
	n := len(arr)
	if n <= 1 {
		return
	}
	PartialQuickSort(arr, comparator, 0, n-1)
}

func PartialQuickSort(arr []interface{}, comparator func(a interface{}, b interface{}) int, start, end int) {
	if start < end {
		index := partition(arr, comparator, start, end)
		PartialQuickSort(arr, comparator, start, index-1)
		PartialQuickSort(arr, comparator, index+1, end)
	}
}

func partition(arr []interface{}, comparator func(a interface{}, b interface{}) int, start, end int) int {
	rand.Seed(time.Now().UnixNano())
	picked := rand.Int()%(end-start+1) + start
	// random pick a element as pivot
	// swap the pivot to the last element
	pivot := arr[picked]
	arr[picked], arr[end] = arr[end], arr[picked]

	i := start
	for j := start; j < end; j++ {
		if comparator(arr[j], pivot) < 0 {
			// swap
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	// pivot be swapped
	arr[i], arr[end] = arr[end], arr[i]
	return i // position of pivot
}
