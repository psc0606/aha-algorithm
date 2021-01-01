package election

// https://leetcode-cn.com/problems/majority-element/
// `Boyer–Moore majority vote algorithm` aha-algorithm/paper/A Fast Majority Vote Algorithm.pdf
// Time: O(n)
// Space: O(1)
// example: [3,2,3]
func majorityElement(nums []int) int {
	// assert nums not empty

	cand := nums[0]
	k := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == cand {
			k++
		} else {
			k--
		}
		if k == 0 {
			cand = nums[i]
			k = 1
		}
	}
	return cand
}
