package stack

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	result := isValid("()")
	expected := true
	if result != expected {
		t.Errorf("expected is [%t], actual is [%t]", expected, result)
	}

	result = isValid("()[]{}")
	expected = true
	if result != expected {
		t.Errorf("expected is [%t], actual is [%t]", expected, result)
	}

	result = isValid("(]")
	expected = false
	if result != expected {
		t.Errorf("expected is [%t], actual is [%t]", expected, result)
	}

	result = isValid("([)]")
	expected = false
	if result != expected {
		t.Errorf("expected is [%t], actual is [%t]", expected, result)
	}

	result = isValid("{[]}")
	expected = true
	if result != expected {
		t.Errorf("expected is [%t], actual is [%t]", expected, result)
	}

	result = isValid("({{{{}}}))")
	expected = false
	if result != expected {
		t.Errorf("expected is [%t], actual is [%t]", expected, result)
	}

	result = isValid("([}}])")
	expected = false
	if result != expected {
		t.Errorf("expected is [%t], actual is [%t]", expected, result)
	}

	result = isValid("({[)")
	expected = false
	if result != expected {
		t.Errorf("expected is [%t], actual is [%t]", expected, result)
	}
}
