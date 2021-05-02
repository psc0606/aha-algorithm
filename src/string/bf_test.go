package string

import (
	"reflect"
	"testing"
)

func TestBf(t *testing.T) {
	actual, expected := bf("a", "aba"), []int{0, 2}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual, expected = bf("a", "ab"), []int{0}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual, expected = bf("aab", "aaaaaaaaaaaaaab"), []int{12}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual, expected = bf("aba", "abababac"), []int{0, 2, 4}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual = bf("", "aba")
	if len(actual) != 0 {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual = bf("aba", "")
	if len(actual) != 0 {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual = bf("aba", "a")
	if len(actual) != 0 {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual = bf("aba", "ab")
	if len(actual) != 0 {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual, expected = bf("aba", "aba"), []int{0}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
}
