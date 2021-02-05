package heap

import (
	"testing"
)

func TestPeek(t *testing.T) {
	h := new(MaxHeap)
	defer func() {
		if r := recover(); r != nil {
			expected := "empty heap, peek panic"
			if r != expected {
				t.Errorf("expected is [%v], actual is [%v]", expected, r)
			}
		}
	}()
	h.Peek()
}

func TestRemove(t *testing.T) {
	h := new(MaxHeap)
	arr := []int{10, 7, 2, 5, 1}
	h.buildHeap(arr)
	actual := h.Size()
	expected := len(arr)
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	actual = h.Remove()
	expected = 10
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	actual = h.Size()
	expected = 4
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	actual = h.Remove()
	expected = 7
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	h.Insert(10)
	actual = h.Peek()
	expected = 10
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	h.Insert(7)
	actual = h.Peek()
	expected = 10
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}

	h.Insert(16)
	actual = h.Peek()
	expected = 16
	if actual != expected {
		t.Errorf("expected is [%d], actual is [%d]", expected, actual)
	}
}
