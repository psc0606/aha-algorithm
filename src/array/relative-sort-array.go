package array

// https://leetcode-cn.com/problems/relative-sort-array/
func relativeSortArray(arr1 []int, arr2 []int) []int {
	var hash [1001]int
	for i := 0; i < len(arr1); i++ {
		hash[arr1[i]]++
	}

	k := 0
	for i := 0; i < len(arr2); i++ {
		for j := 0; j < hash[arr2[i]]; j++ {
			arr1[k] = arr2[i]
			k++
		}
		// already used, then clear
		hash[arr2[i]] = 0
	}

	for i := 0; i < len(hash); i++ {
		for j := 0; j < hash[i]; j++ {
			arr1[k] = i
			k++
		}
	}
	return arr1
}
