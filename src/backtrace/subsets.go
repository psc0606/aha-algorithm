package backtrace

// https://leetcode-cn.com/problems/subsets/
// Time: O(n^2)
func subsets(nums []int) [][]int {
	//n := len(nums)
	var ret [][]int
	dfsSubsets(nums, 0, []int{}, &ret)
	return ret
}

func dfsSubsets(nums []int, index int, path []int, result *[][]int) {
	if index == len(nums) {
		arr := append([]int{}, path...)
		*result = append(*result, arr)
		return
	}
	path = append(path, nums[index])
	dfsSubsets(nums, index+1, path, result)
	path = path[:len(path)-1] // delete tail
	dfsSubsets(nums, index+1, path, result)
}

// https://leetcode-cn.com/problems/subsets/solution/zi-ji-by-leetcode-solution/
// Every element have two status: in-set or not-in-set.
// We can use `1` represents in-set, `0` represents not-in-set.
// eg: [1, 2, 3]
// 000 -> []
// 001 -> [3]
// 010 -> [2]
// 011 -> [2, 3]
// 100 -> [1]
// 101 -> [1, 3]
// 110 -> [1, 2]
// 111 -> [1, 2, 3]
// Time: O(n * 2^n)
func subsetsWithBits(nums []int) [][]int {
	var ret [][]int
	n := len(nums)
	for mask := 0; mask < 1<<n; mask++ { // 1 << n = 2^n
		cur := mask
		var arr []int
		index := 0
		for cur > 0 {
			if cur&0x01 == 1 {
				arr = append(arr, nums[index])
			}
			cur >>= 1
			index++
		}
		ret = append(ret, arr)
	}
	return ret
}
