package array

// https://leetcode-cn.com/problems/decompress-run-length-encoded-list/
// like ziplist in redis
// [freq, val] = [nums[2*i], nums[2*i+1]]
// input: nums = [1,2,3,4]
// output: [2,4,4,4]
func decompressRLElist(nums []int) []int {
	var decodeArr []int
	for i := 0; i < len(nums); i += 2 {
		freq := nums[i]
		val := nums[i+1]
		for j := 0; j < freq; j++ {
			decodeArr = append(decodeArr, val)
		}
	}
	return decodeArr
}
