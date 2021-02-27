package backtrace

import (
	"aha-algorithm/src/util"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	actual := letterCombinations("")
	expected := []string{}
	if !util.ArrayEqualWithoutSeq(actual, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual = letterCombinations("2")
	expected = []string{"a", "b", "c"}
	if !util.ArrayEqualWithoutSeq(actual, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual = letterCombinations("23")
	expected = []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}
	if !util.ArrayEqualWithoutSeq(actual, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual = letterCombinations("234")
	expected = []string{"adg", "adh", "adi", "aeg", "aeh", "aei", "afg", "afh", "afi", "bdg", "bdh", "bdi", "beg", "beh", "bei", "bfg", "bfh", "bfi", "cdg", "cdh", "cdi", "ceg", "ceh", "cei", "cfg", "cfh", "cfi"}
	if !util.ArrayEqualWithoutSeq(actual, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
}
