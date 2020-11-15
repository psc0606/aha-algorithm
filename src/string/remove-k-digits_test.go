package string

import "testing"

func TestRemoveKdigits(t *testing.T) {
	min := removeKdigits("1432219", 3)
	actual := "1219"
	if min != actual {
		t.Errorf("expected is [%s], actual is [%s]", actual, min)
	}

	min = removeKdigits("10200", 1)
	actual = "200"
	if min != actual {
		t.Errorf("expected is [%s], actual is [%s]", actual, min)
	}

	min = removeKdigits("10", 2)
	actual = "0"
	if min != actual {
		t.Errorf("expected is [%s], actual is [%s]", actual, min)
	}
}
