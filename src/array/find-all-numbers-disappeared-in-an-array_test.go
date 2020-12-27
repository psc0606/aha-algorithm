package array

import (
	"reflect"
	"testing"
)

func TestFindDisappearedNumbers(t *testing.T) {
	missing := findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1})
	expected := []int{5, 6}
	if !reflect.DeepEqual(missing, expected) {
		t.Errorf("expected equal")
	}
}
