package optimalization

import "testing"

func TestLongestPalindrome(t *testing.T) {
	actual := longestPalindrome("babad")
	expected := "bab"
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual = longestPalindrome("bcbca")
	expected = "bcb"
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual = longestPalindrome("cbccb")
	expected = "bccb"
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual = longestPalindrome("a")
	expected = "a"
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual = longestPalindrome("aa")
	expected = "aa"
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual = longestPalindrome("abc")
	expected = "a"
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
}
