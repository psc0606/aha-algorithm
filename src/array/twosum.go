package array

// https://leetcode-cn.com/problems/two-sum/
// nums = [2, 7, 11, 15], target = 9 -> [0, 1]
// traverse twice
func twoSum(nums []int, target int) []int {
	if nums == nil {
		return nil
	}

	// key is num, value is index
	hashmap := make(map[int]int)
	for index, num := range nums {
		hashmap[num] = index
	}

	// make a slice
	indexArray := make([]int, 2)
	for index, num := range nums {
		find, ok := hashmap[target-num]
		// not find
		if !ok {
			continue
		}
		// find one, but not itself
		if find != index {
			indexArray[0] = index
			indexArray[1] = find
			break
		}
	}
	return indexArray
}

func twoSumWithOnceTraversal(nums []int, target int) []int {
	if nums == nil {
		return nil
	}

	// key is num, value is index
	hashmap := make(map[int]int)
	// make a slice
	indexArray := make([]int, 2)
	for index, num := range nums {
		find, ok := hashmap[target-num]
		if ok {
			// find one, but not itself
			if find != index {
				indexArray[0] = find
				indexArray[1] = index
				break
			}
		} else {
			hashmap[num] = index
		}
	}
	return indexArray
}
