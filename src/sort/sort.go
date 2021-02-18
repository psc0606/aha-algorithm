package sort

func IsSorted(arr []interface{}, comparator func(a interface{}, b interface{}) int) bool {
	for i := 1; i < len(arr); i++ {
		if comparator(arr[i], arr[i-1]) < 0 {
			return false
		}
	}
	return true
}
