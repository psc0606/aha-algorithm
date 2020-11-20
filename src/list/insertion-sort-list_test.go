package list

import (
	"testing"
)

func TestInsertionSortList(t *testing.T) {
	head := BuildList([]int{4, 2, 1, 3})
	PrintList(head)
	sortHead := insertionSortList(head)
	PrintList(sortHead)
	expected := BuildList([]int{1, 2, 3, 4})
	if TestListEqual(expected, sortHead) {
	}

	head = BuildList([]int{-1, 5, 3, 4, 0})
	PrintList(head)
	sortHead = insertionSortList(head)
	PrintList(sortHead)
	expected = BuildList([]int{-1, 0, 3, 4, 5})
	if TestListEqual(expected, sortHead) {
	}
}
