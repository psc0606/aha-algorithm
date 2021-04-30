package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTypeOf(t *testing.T) {
	m := new(map[int]string)
	n := make(map[int]string)
	fmt.Println("m:", reflect.TypeOf(m)) // *map[int]string
	fmt.Println("n:", reflect.TypeOf(n)) // map[int]string
}
