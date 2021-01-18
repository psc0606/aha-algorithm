package util

const IntMax = int(^uint(0) >> 1)
const IntMin = ^IntMax

func Max(a int, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}
