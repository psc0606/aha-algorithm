package string

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	actual, expected := longestCommonPrefix([]string{"flower", "flow", "flight"}), "fl"
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", actual, expected)
	}

	actual, expected = longestCommonPrefix([]string{"flower", "flow", "flight", ""}), ""
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", actual, expected)
	}

	actual, expected = longestCommonPrefix([]string{""}), ""
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", actual, expected)
	}

	actual, expected = longestCommonPrefix([]string{"ab", "a"}), "a"
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", actual, expected)
	}
}

func TestLongestCommonPrefixII(t *testing.T) {
	actual, expected := longestCommonPrefixII([]string{"flower", "flow", "flight"}), "fl"
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", actual, expected)
	}

	actual, expected = longestCommonPrefixII([]string{"flower", "flow", "flight", ""}), ""
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", actual, expected)
	}

	actual, expected = longestCommonPrefixII([]string{""}), ""
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", actual, expected)
	}

	actual, expected = longestCommonPrefixII([]string{"a"}), "a"
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", actual, expected)
	}

	actual, expected = longestCommonPrefixII([]string{"ab", "a"}), "a"
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", actual, expected)
	}
}
