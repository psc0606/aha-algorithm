package array

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	digits := []int{1, 3, 9}
	actual := plusOne(digits)
	expected := []int{1, 4, 0}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	digits = []int{8}
	actual = plusOne(digits)
	expected = []int{9}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	digits = []int{9}
	actual = plusOne(digits)
	expected = []int{1, 0}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	digits = []int{9, 9}
	actual = plusOne(digits)
	expected = []int{1, 0, 0}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
}
