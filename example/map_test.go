package main

import (
	"fmt"
	"testing"
)

func TestGetNotExists(t *testing.T) {
	ht := make(map[int]bool)
	v, ok := ht[0]
	// key not exists, but return default false.
	if v != false {
		t.Errorf("expected is [%v], actual is [%v]", false, v)
	}
	// ok tell you the element exists or not.
	if ok != false {
		t.Errorf("expected is [%v], actual is [%v]", false, v)
	}
}

func TestMapIterator(t *testing.T) {
	ht := make(map[int]int)
	ht[0] = 0
	ht[1] = 1
	for k, v := range ht {
		fmt.Println(k, v)
	}
	// not exists, but return 0
	v := ht[3]
	fmt.Println(v)
}
