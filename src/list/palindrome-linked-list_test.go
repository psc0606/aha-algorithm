package list

import "testing"

func TestIsPalindrome(t *testing.T) {
	head := BuildList([]int{4, 2, 1, 3})
	PrintList(head)
	actual := isPalindrome(head)
	expected := false
	if expected != actual {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	head = BuildList([]int{1})
	PrintList(head)
	actual = isPalindrome(head)
	expected = true
	if expected != actual {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	head = BuildList([]int{1, 1})
	PrintList(head)
	actual = isPalindrome(head)
	expected = true
	if expected != actual {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	head = BuildList([]int{1, 3, 1})
	PrintList(head)
	actual = isPalindrome(head)
	expected = true
	if expected != actual {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	head = BuildList([]int{1, 3, 3, 1})
	PrintList(head)
	actual = isPalindrome(head)
	expected = true
	if expected != actual {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	head = BuildList([]int{1, 3, 2, 1})
	PrintList(head)
	actual = isPalindrome(head)
	expected = false
	if expected != actual {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
}
