package cache

import "testing"

func TestLRU(t *testing.T) {
	// capacity 0
	obj := Constructor(0)
	obj.Put(1, 1)
	actual, expected := obj.Get(1), -1
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	// capacity 1
	obj = Constructor(1)
	obj.Put(1, 1)
	actual, expected = obj.Get(1), 1
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	obj.Put(2, 2)
	actual, expected = obj.Get(1), -1
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	obj.Put(3, 3)
	actual, expected = obj.Get(3), 3
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
	actual, expected = obj.Size(), 1
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	// capacity 2
	obj = Constructor(2)
	obj.Put(1, 1)
	actual, expected = obj.Get(1), 1
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
	actual, expected = obj.Size(), 1
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	obj.Put(2, 2)
	actual, expected = obj.Get(1), 1
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
	actual, expected = obj.Get(2), 2
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
	actual, expected = obj.Get(1), 1
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
	actual, expected = obj.Size(), 2
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	obj.Put(1, 1)
	obj.Put(2, 2)
	actual, expected = obj.Get(1), 1
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
	actual, expected = obj.Size(), 2
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
	obj.Put(3, 3)
	actual, expected = obj.Get(2), -1
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
	actual, expected = obj.Size(), 2
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
	obj.Put(4, 4)
	actual, expected = obj.Get(1), -1
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
	actual, expected = obj.Get(3), 3
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
	actual, expected = obj.Get(4), 4
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	//
	obj.Put(2, 1)
	obj.Put(2, 2)
	actual, expected = obj.Get(2), 2
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
	obj.Put(1, 1)
	obj.Put(4, 1)
	actual, expected = obj.Get(2), -1
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
}
