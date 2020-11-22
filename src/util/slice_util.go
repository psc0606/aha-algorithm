package util

// Test two slice equal or not
func SliceEqual(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	i := 0
	for ; i < len(a) && a[i] == b[i]; i++ {
	}
	return i == len(a)
}

// reverse the slice in place
func SliceReverse(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
