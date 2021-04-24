package array

import (
	"reflect"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	actual, expected := spiralOrder(matrix), []int{1, 2, 3, 6, 9, 8, 7, 4, 5}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	matrix = [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
	}
	actual, expected = spiralOrder(matrix), []int{1, 2, 3, 4, 8, 7, 6, 5}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	matrix = [][]int{
		{1, 2, 3, 4},
	}
	actual, expected = spiralOrder(matrix), []int{1, 2, 3, 4}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
}
