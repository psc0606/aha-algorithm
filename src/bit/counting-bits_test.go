package bit

import (
	"aha-algorithm/src/util"
	"testing"
)

func TestCountBits(t *testing.T) {
	actual := countBits(0)
	expected := []int{0}
	if !util.SliceEqual(actual, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual = countBits(1)
	expected = []int{0, 1}
	if !util.SliceEqual(actual, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual = countBits(2)
	expected = []int{0, 1, 1}
	if !util.SliceEqual(actual, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual = countBits(3)
	expected = []int{0, 1, 1, 2}
	if !util.SliceEqual(actual, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual = countBits(4)
	expected = []int{0, 1, 1, 2, 1}
	if !util.SliceEqual(actual, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
}
