package sort

// Time: O(nlogn)
// Space: O(n)
func MergeSort(arr []interface{}, comparator func(a interface{}, b interface{}) int) {
	n := len(arr)
	if n <= 1 {
		return
	}
	step := 1
	for step < n {
		step *= 2
		for i := 0; i < n; i += step {
			start, end := i, i+step-1
			mid := (start + end + 1) / 2
			doMerge(arr, comparator, start, min(mid, n-1), min(end, n-1))
		}
	}
}

func doMerge(arr []interface{}, comparator func(a interface{}, b interface{}) int, start, mid, end int) {
	i := start
	j := mid
	sorted := make([]interface{}, end-start+1)
	k := 0
	for i < mid && j <= end {
		if comparator(arr[i], arr[j]) < 0 {
			sorted[k] = arr[i]
			i++
		} else {
			sorted[k] = arr[j]
			j++
		}
		k++
	}
	// the rest of left
	for ; i < mid; i++ {
		sorted[k] = arr[i]
		k++
	}
	// the rest of right
	for ; j <= end; j++ {
		sorted[k] = arr[j]
		k++
	}
	// copy the sorted
	k = 0
	for z := start; z <= end; z++ {
		arr[z] = sorted[k]
		k++
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
