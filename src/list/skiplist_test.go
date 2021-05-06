package list

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	obj := ConstructorSkiplist()
	obj.Add(1)
	obj.print()
	obj.Add(2)
	obj.print()
	obj.Add(3)
	obj.print()
	actual, expected := obj.Search(1), true
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
	actual, expected = obj.Search(2), true
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
	actual, expected = obj.Search(3), true
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
	actual, expected = obj.Search(-1), false
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
	obj.Add(4)
	obj.print()
	obj.Add(8)
	obj.print()
	obj.Add(10)
	actual, expected = obj.Search(4), true
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
	actual, expected = obj.Search(8), true
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
	actual, expected = obj.Search(10), true
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
	for i := 11; i < 100; i++ {
		obj.Add(i)
	}
	actual, expected = obj.Search(99), true
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	obj.print()
	obj.Erase(1)
	actual, expected = obj.Search(1), false
	if actual != expected {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
	obj.print()

	for i := 0; i < 110; i++ {
		obj.Erase(i)
		obj.print()
	}

	fmt.Println("----------------------------------------------------")
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		random := rand.Intn(100)
		fmt.Printf("Add %d\n", random)
		obj.Add(random)
		obj.print()
	}
}
