package optimalization

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	actual := lengthOfLongestSubstring("abcabcbb")
	expected := 3
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual = lengthOfLongestSubstring("abcadea")
	expected = 5
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual = lengthOfLongestSubstring("a")
	expected = 1
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual = lengthOfLongestSubstring("aa")
	expected = 1
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual = lengthOfLongestSubstring("aab")
	expected = 2
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual = lengthOfLongestSubstring("ab")
	expected = 2
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual = lengthOfLongestSubstring("")
	expected = 0
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual = lengthOfLongestSubstring("")
	expected = 0
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
}

func TestLengthOfLongestSubstring2(t *testing.T) {
	actual := lengthOfLongestSubstring2("abcabcbb")
	expected := 3
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual = lengthOfLongestSubstring2("abcadea")
	expected = 5
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual = lengthOfLongestSubstring2("a")
	expected = 1
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual = lengthOfLongestSubstring2("aa")
	expected = 1
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual = lengthOfLongestSubstring2("aab")
	expected = 2
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual = lengthOfLongestSubstring2("ab")
	expected = 2
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual = lengthOfLongestSubstring2("")
	expected = 0
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual = lengthOfLongestSubstring2("")
	expected = 0
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
}
