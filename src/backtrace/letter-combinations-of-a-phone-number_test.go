package backtrace

import (
	"fmt"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	actual := letterCombinations("")
	fmt.Println(actual)

	actual = letterCombinations("2")
	fmt.Println(actual)

	actual = letterCombinations("23")
	fmt.Println(actual)

	actual = letterCombinations("234")
	fmt.Println(actual)
}
