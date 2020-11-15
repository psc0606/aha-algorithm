package main

import (
	"fmt"
	"testing"
)

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
