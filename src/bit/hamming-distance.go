package bit

// https://leetcode-cn.com/problems/hamming-distance/
func hammingDistance(x int, y int) int {
	return numberOfBit1(x ^ y)
}

// number of bit-1
// input only positive number
func numberOfBit1(x int) int {
	count := 0
	for x > 0 {
		if x&0x01 == 1 {
			count++
		}
		x >>= 1
	}
	return count
}

// use a bit trick: n & (n-1)
func numberOfBit1WithTrick(x int) int {
	c := 0
	for x > 0 {
		c++
		x &= x - 1
	}
	return c
}
