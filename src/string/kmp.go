package string

// https://www.nowcoder.com/practice/bb1615c381cc4237919d1aa448083bcc
// Time: O(n)
// Space: O(n)
// `S` - pattern string
// `T` - target string
// return matching times.
func kmp(S string, T string) int {
	m := len(S)
	n := len(T)
	next := getNext(S)
	ans := 0
	for i, j := 0, 0; i < n; {
		// when j = -1, we should move i to next, j(is -1) set to 0.
		if j == -1 || S[j] == T[i] {
			i++
			j++
		} else {
			j = next[j]
		}
		// find matching, if will want to find more, we must continue.
		if j == m {
			j = next[j]
			ans++
		}
	}
	return ans
}

// kmp next
// next[j]=k, j is the position of mismatching on pattern, k is the new position should jump to.
func getNext(pattern string) []int {
	m := len(pattern)
	next := make([]int, m+1)
	next[0] = -1
	k, j := -1, 0
	for j < m {
		if k == -1 || pattern[k] == pattern[j] {
			j++
			k++
			next[j] = k
		} else {
			k = next[k]
		}
	}
	return next
}
