package array

import (
	"reflect"
	"testing"
)

func TestIntersectIII(t *testing.T) {
	actual, expected := intersectIII([]int{1, 3, 7, 10}, []int{3, 10, 100}), []int{3, 10}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual, expected = intersectIII([]int{1, 3, 3, 7, 10}, []int{3, 10, 100}), []int{3, 10}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual, expected = intersectIII([]int{1, 3, 3, 7, 10}, []int{3, 3, 10, 100}), []int{3, 3, 10}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
}
