package tree

import (
	"fmt"
	"testing"
)

func TestDisjointSet(t *testing.T) {
	ds := new(DisjointSet)
	ds.Init(5000)
	for i := 0; i < 100; i++ {
		ds.MakeSet(i)
	}
	fmt.Println(ds.father)

	ds.Union(0, 1)
	actual, expected := ds.Same(0, 1), true
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual, expected = ds.Same(1, 2), false
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual, expected = ds.Same(2, 3), false
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	ds.Union(4, 5)
	ds.Union(1, 4)
	actual, expected = ds.Same(0, 4), true
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual, expected = ds.Same(1, 4), true
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual, expected = ds.Same(0, 5), true
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	ds.Union(6, 0)
	ds.Union(7, 0)
	ds.Union(8, 0)

	actual, expected = ds.Same(0, 6), true
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual, expected = ds.Same(0, 7), true
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual, expected = ds.Same(0, 8), true
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual, expected = ds.Same(1, 8), true
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	actual, expected = ds.Same(4, 8), true
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
	fmt.Println(ds.father)
}
