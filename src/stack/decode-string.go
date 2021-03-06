package stack

// https://leetcode-cn.com/problems/decode-string/
// @see https://leetcode-cn.com/problems/valid-parentheses

type tuple struct {
	num     int
	content string
}

// https://leetcode-cn.com/problems/decode-string/comments/415775
// The key problem is how to solve nested "[" and "]"
// eg: 3[af4[b]]
// stack is:
// |         | content="b"
// |(4, "af")|
// |(3, "")  |
// `s` - we can assume the input string is always legal.
func decodeString(s string) string {
	var stack []tuple
	num := 0
	content := ""
	for _, c := range s {
		if c >= '0' && c <= '9' {
			num = int(c-'0') + num*10
		} else if c == '[' {
			stack = append(stack, tuple{num: num, content: content})
			num, content = 0, ""
		} else if c == ']' {
			top := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			content = top.content + duplicateString(content, top.num)
		} else {
			content += string(c)
		}
	}
	return content
}

func duplicateString(str string, num int) string {
	var ret string
	for i := 0; i < num; i++ {
		ret += str
	}
	return ret
}
