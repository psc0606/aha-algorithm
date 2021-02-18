package sort

// Insertion sort is the simplest sort algorithm, it will be effect when data is very small.
// Time: O(n^2)
// Space: O(1)
func InsertSort(arr []interface{}, comparator func(a interface{}, b interface{}) int) {
	n := len(arr)
	if n < 2 {
		return
	}
	for i := 1; i < n; i++ {
		// arr[i] > arr[j]
		j := i - 1
		key := arr[i]
		for ; j >= 0 && comparator(key, arr[j]) < 0; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = key
	}
}
