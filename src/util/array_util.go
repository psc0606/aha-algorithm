package util

func ArrayEqualWithoutSeq(arr1 []string, arr2 []string) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	ht := make(map[string]bool)
	for _, ele := range arr1 {
		ht[ele] = true
	}
	for _, ele := range arr2 {
		_, ok := ht[ele]
		if !ok {
			return false
		}
	}
	return true
}

func ArraySum(arr []int) int {
	ret := 0
	for _, v := range arr {
		ret += v
	}
	return ret
}
