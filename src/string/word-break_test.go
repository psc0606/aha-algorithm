package string

import "testing"

func TestWordBreak(t *testing.T) {
	actual := wordBreak("catstandog", []string{"cat", "cats", "stand", "and", "dog"})
	expected := false
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual = wordBreak("catstanddog", []string{"cat", "cats", "stand", "and", "dog"})
	expected = true
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual = wordBreak("catstandog", []string{"cat", "cats", "stand", "and", "dog", "tan"})
	expected = true
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual = wordBreak("leetcode", []string{"leet", "code"})
	expected = true
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual = wordBreak("leetcode", []string{"leetcode"})
	expected = true
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual = wordBreak("leetcode", []string{})
	expected = false
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual = wordBreak("l", []string{"l"})
	expected = true
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual = wordBreak("lll", []string{"l"})
	expected = true
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
}
