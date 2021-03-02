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

	// self define sort, such as descend.
	sort.Sort(IntSlice(arr))
	fmt.Println(arr)
}

type IntSlice []int

func (s IntSlice) Len() int           { return len(s) }
func (s IntSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s IntSlice) Less(i, j int) bool { return s[i] > s[j] }
