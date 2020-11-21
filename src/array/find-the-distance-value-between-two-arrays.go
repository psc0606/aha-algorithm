package array

// https://leetcode-cn.com/problems/find-the-distance-value-between-two-arrays/
// Time: O(m*n)
// Space: O(1)
func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	res := 0
	for _, ele1 := range arr1 {
		match := true
		for _, ele2 := range arr2 {
			if abs(ele1-ele2) >= d {
				match = false
				break
			}
		}
		if match {
			res++
		}
	}
	return res
}
