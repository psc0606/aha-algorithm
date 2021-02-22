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

// merge two slice into one
func TestMergeTwoSlice(t *testing.T) {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6}
	s1 = append(s1, s2...)
	fmt.Println(s1)
}

func TestSlice(t *testing.T) {
	str := "123456"

	// only one element
	fmt.Println(str[1:2])
	// empty
	fmt.Println(str[2:2])

	fmt.Println(str[1:])
}

func TestOutOfRange(t *testing.T) {
	str := "123456"
	fmt.Println(str[5:len(str)])
	// error, out of range, compile error
	// fmt.Println(str[5:len(str)+1])
}

func TestInsert(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("catched panic：%s\n", r)
		}
	}()

	var arr []int
	// throw panic, but I will catch it by a defined defer func.
	arr[10] = 1

	arr = make([]int, 11)
	// ok
	arr[10] = 1
}

func TestArrayAppend(t *testing.T) {
	// init capacity is 0, max capacity is 10.
	arr := make([]int, 0, 10)
	fmt.Println(len(arr)) // 0
	fmt.Println(arr)

	for i := 0; i < 11; i++ {
		// append will auto expand the capacity of array.
		// arr = append(arr, i)

		// index out of range
		arr[i] = i
	}
	fmt.Println(arr)
}
