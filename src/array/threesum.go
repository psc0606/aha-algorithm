package array

import "sort"

// https://leetcode-cn.com/problems/3sum/
// https://leetcode-cn.com/problems/3sum/solution/san-shu-zhi-he-by-leetcode-solution/
// Time: O(n^2 + nlogn) = O(n^2)
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	var ret [][]int
	// first loop find the element-1
	for i := 0; i < n; i++ {
		// first element or next different element
		if i == 0 || nums[i] != nums[i-1] {
			k := n - 1
			// second loop find element-2
			for j := i + 1; j < n; j++ {
				if j == i+1 || nums[j] != nums[j-1] {
					for k > j && nums[i]+nums[j]+nums[k] > 0 {
						k--
					}
					if k > j {
						ok := checkZero(nums[i], nums[j], nums[k])
						if ok {
							ret = append(ret, []int{nums[i], nums[j], nums[k]})
						}
					}
				}
			}
		}
	}
	return ret
}

func checkZero(a, b, c int) bool {
	return a+b+c == 0
}
