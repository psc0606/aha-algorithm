package stack

import "testing"

func TestMinStack(t *testing.T) {
	obj := Constructor()

	obj.Push(-2)
	actual := obj.GetMin()
	expected := -2
	if expected != actual {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	obj.Push(0)
	actual = obj.GetMin()
	expected = -2
	if expected != actual {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	obj.Push(-3)
	actual = obj.GetMin()
	expected = -3
	if expected != actual {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	obj.Push(5)
	actual = obj.GetMin()
	expected = -3
	if expected != actual {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	obj.Push(10)
	actual = obj.GetMin()
	expected = -3
	if expected != actual {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual = obj.Len()
	expected = 5
	if expected != actual {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	obj.Pop()
	actual = obj.GetMin()
	expected = -3
	if expected != actual {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	obj.Pop()
	actual = obj.GetMin()
	expected = -3
	if expected != actual {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	obj.Pop()
	actual = obj.GetMin()
	expected = -2
	if expected != actual {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	obj.Pop()
	actual = obj.GetMin()
	expected = -2
	if expected != actual {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
}
