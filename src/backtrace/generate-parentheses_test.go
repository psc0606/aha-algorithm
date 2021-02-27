package backtrace

import (
	"aha-algorithm/src/util"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	// [()]
	actual := generateParenthesis(1)
	expected := []string{"()"}
	if !util.ArrayEqualWithoutSeq(actual, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	// [(()) ()()]
	actual = generateParenthesis(2)
	expected = []string{"()()", "(())"}
	if !util.ArrayEqualWithoutSeq(actual, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}

	// [((())) (()()) (())() ()(()) ()()()]
	actual = generateParenthesis(3)
	expected = []string{"((()))", "(()())", "(())()", "()(())", "()()()"}
	if !util.ArrayEqualWithoutSeq(actual, expected) {
		t.Errorf("expected is [%v], actual is [%v]", expected, actual)
	}
}
