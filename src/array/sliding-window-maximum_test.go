package array

import (
	"reflect"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	actual, expected := maxSlidingWindow(nums, 3), []int{3, 3, 5, 5, 6, 7}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	nums = []int{1, 3, 2, 4, 5, 3, 6, 7}
	actual, expected = maxSlidingWindow(nums, 3), []int{3, 4, 5, 5, 6, 7}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	nums = []int{1}
	actual, expected = maxSlidingWindow(nums, 1), []int{1}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
}

func TestMaxSlidingWindowOptimizeWithDequeue(t *testing.T) {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	actual, expected := maxSlidingWindowOptimizeWithDequeue(nums, 3), []int{3, 3, 5, 5, 6, 7}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	nums = []int{1, 3, 2, 4, 5, 3, 6, 7}
	actual, expected = maxSlidingWindowOptimizeWithDequeue(nums, 3), []int{3, 4, 5, 5, 6, 7}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	nums = []int{1}
	actual, expected = maxSlidingWindowOptimizeWithDequeue(nums, 1), []int{1}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	nums = []int{1, -1}
	actual, expected = maxSlidingWindowOptimizeWithDequeue(nums, 1), []int{1, -1}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	nums = []int{1, -1, 1}
	actual, expected = maxSlidingWindowOptimizeWithDequeue(nums, 2), []int{1, 1}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	nums = []int{1, -1, -1}
	actual, expected = maxSlidingWindowOptimizeWithDequeue(nums, 1), []int{1, -1, -1}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	nums = []int{1, -1, -1}
	actual, expected = maxSlidingWindowOptimizeWithDequeue(nums, 2), []int{1, -1}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
}
