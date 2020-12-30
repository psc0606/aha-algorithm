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
