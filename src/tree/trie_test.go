package tree

import "testing"

func TestInsert(t *testing.T) {
	obj := Constructor()
	obj.Insert("hello")
	obj.Insert("helle")
	actual := obj.Search("hello")
	expected := true
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual = obj.Search("helle")
	expected = true
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual = obj.Search("hell")
	expected = false
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual = obj.StartsWith("hel")
	expected = true
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual = obj.StartsWith("helle")
	expected = true
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual = obj.StartsWith("hellx")
	expected = false
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
}
