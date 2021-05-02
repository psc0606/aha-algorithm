package string

// String matching - Brute Force Algorithm.
// O(m*(n-m+1))
func bf(S string, T string) (ans []int) {
	m, n := len(S), len(T)
	if m == 0 || n == 0 {
		return ans
	}
	for i := 0; i < n; i++ {
		j, k := 0, i
		for j < m && k < n && S[j] == T[k] {
			j++
			k++
		}
		if j == m {
			ans = append(ans, i)
		}
	}
	return ans
}
