package sort

// Bubble Sort
// Time: O(n^2)
// Space: O(1)
func BubbleSort(arr []interface{}, comparator func(a interface{}, b interface{}) int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 1; j < n-i; j++ {
			// if arr[j-1] > arr[j], then swap both
			if comparator(arr[j-1], arr[j]) > 0 {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}
}
