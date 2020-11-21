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
