package string

import (
	"reflect"
	"testing"
)

func TestGetNext(t *testing.T) {
	actual, expected := getNext("a"), []int{-1}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual, expected = getNext("ab"), []int{-1, 0}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual, expected = getNext("aa"), []int{-1, 0}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual, expected = getNext("aba"), []int{-1, 0, 0}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual, expected = getNext("abab"), []int{-1, 0, 0, 1}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual, expected = getNext("aaaa"), []int{-1, 0, 1, 2}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual, expected = getNext("aaaaa"), []int{-1, 0, 1, 2, 3}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
}

func TestKmp(t *testing.T) {
	actual, expected := kmp("abc", "ababc"), 1
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual, expected = kmp("aba", "bababa"), 2
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual, expected = kmp("aba", "bababa"), 2
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
}
