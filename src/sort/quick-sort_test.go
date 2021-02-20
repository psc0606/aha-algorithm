package sort

import (
	"aha-algorithm/src/util"
	"fmt"
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	comparator := func(a interface{}, b interface{}) int {
		return a.(int) - b.(int)
	}

	arr := []interface{}{1}
	QuickSort(arr, comparator)
	expected := []interface{}{1}
	actual := arr
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	arr = []interface{}{5, 3}
	QuickSort(arr, comparator)
	expected = []interface{}{3, 5}
	actual = arr
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	arr = []interface{}{5, 3, 2, 1, 6, 8}
	QuickSort(arr, comparator)
	expected = []interface{}{1, 2, 3, 5, 6, 8}
	actual = arr
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	arr = util.RandArrayInterface(10)
	fmt.Println(arr)
	fmt.Println(IsSorted(arr, comparator))
	QuickSort(arr, comparator)
	fmt.Println(arr)
	if !IsSorted(arr, comparator) {
		t.Errorf("expected is [%v], actual is [%v]", true, false)
	}

	arr = util.RandArrayInterface(1000)
	fmt.Println(arr)
	fmt.Println(IsSorted(arr, comparator))
	QuickSort(arr, comparator)
	fmt.Println(arr)
	if !IsSorted(arr, comparator) {
		t.Errorf("expected is [%v], actual is [%v]", true, false)
	}
}
