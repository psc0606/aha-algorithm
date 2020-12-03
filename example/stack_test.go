package main

import (
	"fmt"
	"testing"
)

// go no stack implement, we can use slice to do it.
func TestStackPush(t *testing.T) {
	var stack []int
	stack = append(stack, 1)
	stack = append(stack, 2)
	stack = append(stack, 3)
	fmt.Println(stack)
}

func TestStackPop(t *testing.T) {
	var stack []int
	stack = append(stack, 1)
	stack = append(stack, 2)
	stack = append(stack, 3)
	stack = stack[:len(stack)-1]
	fmt.Println(stack)
}

func TestStackPeek(t *testing.T) {
	var stack []int
	stack = append(stack, 1)
	stack = append(stack, 2)
	stack = append(stack, 3)
	fmt.Println(stack[len(stack)-1])
}
