package array

import (
	"reflect"
	"testing"
)

func TestIntersectIV(t *testing.T) {
	actual, expected := intersectIV([]int{1, 3, 7, 10, 20}, []int{3, 10, 11, 12, 13, 18, 19, 20, 21, 30, 35, 40, 78, 100}), []int{3, 10, 20}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual, expected = intersectIV([]int{3}, []int{3, 10, 11, 12, 13, 18, 19, 20, 21, 30, 35, 40, 78, 100}), []int{3}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual, expected = intersectIV([]int{3, 100}, []int{3, 10, 11, 12, 13, 18, 19, 20, 21, 30, 35, 40, 78, 100}), []int{3, 100}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual, expected = intersectIV([]int{33, 40}, []int{3, 10, 11, 12, 13, 18, 19, 20, 21, 30, 35, 40, 78, 100}), []int{40}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual, expected = intersectIV([]int{33, 40}, []int{3, 10, 11, 12, 13, 18, 19, 20, 21, 30, 35, 40, 40, 78, 100}), []int{40}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	//
	//actual, expected = intersectIV([]int{33, 40, 40}, []int{3, 10, 11, 12, 13, 18, 19, 20, 21, 30, 35, 40, 78, 100}), []int{40}
	actual, expected = intersectIV([]int{33, 40}, []int{3, 10, 11, 12, 13, 18, 19, 20, 21, 30, 35, 40, 78, 100}), []int{40}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
}
