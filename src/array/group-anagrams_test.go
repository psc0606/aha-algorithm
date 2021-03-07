package array

import (
	"fmt"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	arr := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	actual := groupAnagrams(arr)
	fmt.Println(actual)
}
