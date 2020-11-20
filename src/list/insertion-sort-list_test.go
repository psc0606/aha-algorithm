package list

import (
	"testing"
)

func TestInsertionSortList(t *testing.T) {
	head := BuildList([]int{4, 2, 1, 3})
	ListPrint(head)
	sortHead := insertionSortList(head)
	ListPrint(sortHead)
	expected := BuildList([]int{1, 2, 3, 4})
	if ListEqualCmp(expected, sortHead) {
	}

	head = BuildList([]int{-1, 5, 3, 4, 0})
	ListPrint(head)
	sortHead = insertionSortList(head)
	ListPrint(sortHead)
	expected = BuildList([]int{-1, 0, 3, 4, 5})
	if ListEqualCmp(expected, sortHead) {
	}
}
