package main

import (
	"container/list"
	"fmt"
	"testing"
)

func TestListTraversal(t *testing.T) {
	l := list.New()
	l.PushBack(13)
	l.PushBack(2)
	l.PushBack(3)
	length := l.Len()
	if length != 3 {
		t.Errorf("expected is [%d], actual is [%d]", 3, length)
	}

	// !warning: iterate list but should not remove anyone.
	for e := l.Front(); e != nil; e = e.Next() {
		// cast to int
		v := (e.Value).(int)
		fmt.Println(v)
	}

	println("------------------------------------------")
	// traversal with safe remove
	var next *list.Element
	for e := l.Front(); e != nil; e = next {
		next = e.Next()
		v := (e.Value).(int)
		if v == 2 {
			// before you remove you must save it's Next point.
			l.Remove(e)
		} else {
			fmt.Println(v)
		}
	}
}

func TestPushNil(t *testing.T) {
	// support push multiple nil
	l := list.New()
	l.PushBack(nil)
	l.PushBack(nil)
	l.PushBack(nil)
	for e := l.Front(); e != nil; e = e.Next() {
		// cast to int
		//v := (e.Value).(*int)
		fmt.Println(e.Value)
		//fmt.Println(v)
	}
}
