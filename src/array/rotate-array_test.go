package array

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	expected := []int{5, 6, 7, 1, 2, 3, 4}
	if !reflect.DeepEqual(nums, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, nums)
	}

	nums = []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 1)
	expected = []int{7, 1, 2, 3, 4, 5, 6}
	if !reflect.DeepEqual(nums, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, nums)
	}

	nums = []int{-1, -100, 3, 99}
	rotate(nums, 2)
	expected = []int{3, 99, -1, -100}
	if !reflect.DeepEqual(nums, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, nums)
	}
}

func TestRotate3(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate3(nums, 3)
	expected := []int{5, 6, 7, 1, 2, 3, 4}
	if !reflect.DeepEqual(nums, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, nums)
	}

	nums = []int{1, 2, 3, 4, 5, 6, 7}
	rotate3(nums, 1)
	expected = []int{7, 1, 2, 3, 4, 5, 6}
	if !reflect.DeepEqual(nums, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, nums)
	}

	nums = []int{-1, -100, 3, 99}
	rotate3(nums, 2)
	expected = []int{3, 99, -1, -100}
	if !reflect.DeepEqual(nums, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, nums)
	}

	nums = []int{-1}
	rotate3(nums, 2)
	expected = []int{-1}
	if !reflect.DeepEqual(nums, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, nums)
	}
}
