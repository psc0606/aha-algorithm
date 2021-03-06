package stack

import "testing"

func TestDecodeString(t *testing.T) {
	actual := decodeString("3[a]2[bc]")
	expected := "aaabcbc"
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual = decodeString("2[3[a]2[bc]]")
	expected = "aaabcbcaaabcbc"
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual = decodeString("abc")
	expected = "abc"
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual = decodeString("4[abc]")
	expected = "abcabcabcabc"
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
}
