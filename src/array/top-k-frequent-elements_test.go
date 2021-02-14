package array

import (
	"reflect"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	actual := topKFrequent([]int{1, 1, 1, 2, 2, 3}, 1)
	expected := []int{1}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual = topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)
	expected = []int{1, 2}
	if !reflect.DeepEqual(array2map(actual), array2map(expected)) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual = topKFrequent([]int{3, 5, 1, 2, 2, 3}, 2)
	expected = []int{2, 3}
	if !reflect.DeepEqual(array2map(actual), array2map(expected)) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	// actual := topKFrequent([]int{33333333, 1, 1, 1, 2, 2}, 2)
	for i := 0; i < 1000; i++ {

		actual := topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)
		expected := []int{1, 2}
		if !reflect.DeepEqual(array2map(actual), array2map(expected)) {
			t.Errorf("expected is [%v], actual is [%v]", expected, actual)
		}
	}
}

func array2map(arr []int) map[int]bool {
	ret := make(map[int]bool)
	for _, v := range arr {
		ret[v] = true
	}
	return ret
}
