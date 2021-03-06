package stack

// https://leetcode-cn.com/problems/valid-parentheses/
// @see https://leetcode-cn.com/problems/decode-string/
// First:
// Time: O(n)
// Space: O(n)
func isValid(s string) bool {
	if len(s) == 0 || len(s)%2 == 1 {
		return false
	}
	var stack []int32
	for _, c := range s {
		if isLeftBracket(c) {
			stack = append(stack, c)
		} else if isRightBracket(c) {
			stackLen := len(stack)
			// must match the bracket on the top
			if stackLen > 0 && isMatchedBracket(stack[stackLen-1], c) {
				stack = stack[:stackLen-1]
			} else {
				return false
			}
		} else {
			// should not occur this
		}
	}
	// must be zero at last
	return len(stack) == 0
}

func isLeftBracket(c int32) bool {
	if c == '(' || c == '[' || c == '{' {
		return true
	} else {
		return false
	}
}

func isRightBracket(c int32) bool {
	if c == ')' || c == ']' || c == '}' {
		return true
	} else {
		return false
	}
}

func isMatchedBracket(c1 int32, c2 int32) bool {
	return c1 == '(' && c2 == ')' || c1 == '[' && c2 == ']' || c1 == '{' && c2 == '}'
}

// Second:
// A better simple solution is remove the match brackets: (),[],{}
// Time: O(n)
// Space: O(1)
// ({{{{}}})) -> ({{{}})) -> ({{})) -> ({)) -> can not remove matched brackets any more.
