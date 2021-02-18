package sort

import (
	"aha-algorithm/src/util"
	"fmt"
	"reflect"
	"testing"
)

func TestShellSort(t *testing.T) {
	comparator := func(a interface{}, b interface{}) int {
		return a.(int) - b.(int)
	}

	arr := []interface{}{1}
	ShellSort(arr, comparator)
	expected := []interface{}{1}
	actual := arr
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	arr = []interface{}{5, 3}
	ShellSort(arr, comparator)
	expected = []interface{}{3, 5}
	actual = arr
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	arr = []interface{}{5, 3, 2, 1, 6, 8}
	ShellSort(arr, comparator)
	expected = []interface{}{1, 2, 3, 5, 6, 8}
	actual = arr
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	arr = util.RandArrayInterface(10)
	fmt.Println(arr)
	fmt.Println(IsSorted(arr, comparator))
	ShellSort(arr, comparator)
	fmt.Println(arr)
	if !IsSorted(arr, comparator) {
		t.Errorf("expected is [%v], actual is [%v]", true, false)
	}

	arr = util.RandArrayInterface(1000)
	fmt.Println(arr)
	fmt.Println(IsSorted(arr, comparator))
	BubbleSort(arr, comparator)
	fmt.Println(arr)
	if !IsSorted(arr, comparator) {
		t.Errorf("expected is [%v], actual is [%v]", true, false)
	}
}
