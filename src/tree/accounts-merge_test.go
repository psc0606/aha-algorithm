package tree

import (
	"fmt"
	"testing"
)

func TestAccountsMerge(t *testing.T) {
	arr := [][]string{
		{"John", "johnsmith@mail.com", "john00@mail.com"},
		{"John", "johnnybravo@mail.com"},
		{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
		{"Mary", "mary@mail.com"},
	}
	actual := accountsMerge(arr)
	fmt.Println(actual)
}
