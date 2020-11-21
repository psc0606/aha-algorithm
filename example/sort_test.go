package main

import (
	"fmt"
	"sort"
	"testing"
)

func TestSortInt(t *testing.T) {
	arr := []int{37, 12, 28, 9, 100, 56, 80, 5, 12}
	sort.Ints(arr)
	fmt.Println(arr)
}
