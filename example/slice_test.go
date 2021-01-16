package main

import (
	"fmt"
	"testing"
)

func TestAppendNil(t *testing.T) {
	var slice []interface{}
	// can append more than one nil
	slice = append(slice, nil)
	slice = append(slice, nil)
	slice = append(slice, nil)
	fmt.Println(slice)
}

func TestAppend(t *testing.T) {
	var slice []int
	slice = append(slice, 1)
	slice = append(slice, 2)
	slice = append(slice, 3)
	slice = append(slice, 4)
	fmt.Println(slice)

	// init capacity
	slice = make([]int, 2)
	slice = append(slice, 1)
	slice = append(slice, 2)
	slice = append(slice, 3)
	slice = append(slice, 4)
	fmt.Println(slice)
}

func TestSlice(t *testing.T) {
	str := "123456"

	// only one element
	fmt.Println(str[1:2])
	// empty
	fmt.Println(str[2:2])

	fmt.Println(str[1:])
}
