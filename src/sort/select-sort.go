package sort

// Every time to select the minimum element, and then swap with the front of the array.
// Time: O(n^2)
// Space: O(1)
func SelectSort(arr []interface{}, comparator func(a interface{}, b interface{}) int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		// find the minimum element in the rest and then swap.
		min := i
		for j := i + 1; j < n; j++ {
			if comparator(arr[min], arr[j]) > 0 {
				min = j
			}
		}
		if min != i {
			// swap
			arr[i], arr[min] = arr[min], arr[i]
		}
	}
}
