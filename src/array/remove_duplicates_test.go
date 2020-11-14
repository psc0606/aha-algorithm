package array

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	count := removeDuplicates(nums)
	if count != 5 {
		t.Errorf("expected is [%d], actual is [%d]", 5, count)
	}
	// deep equal two slice
	if !reflect.DeepEqual(nums[0:count], []int{0, 1, 2, 3, 4}) {
		t.Errorf("expected equal")
	}
}
