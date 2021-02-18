package sort

// https://en.wikipedia.org/wiki/Shellsort
// Shellsort, also known as Shell sort or Shell's method, is an in-place comparison sort.
// It can be seen as either a generalization of sorting by exchange (bubble sort) or sorting by insertion (insertion sort).
// The method starts by sorting pairs of elements far apart from each other, then progressively reducing the gap between elements to be compared.
// By starting with far apart elements, it can move some out-of-place elements into position faster than
// a simple nearest neighbor exchange. Donald Shell published the first version of this sort in 1959.
// The running time of Shellsort is heavily dependent on the gap sequence it uses.
// For many practical variants, determining their time complexity remains an open problem.

// Shellsort is an optimization of insertion sort that allows the exchange of items that are far apart.
func ShellSort(arr []interface{}, comparator func(a interface{}, b interface{}) int) {
	n := len(arr)
	// Using Marcin Ciura's gap sequence, with an inner insertion sort.
	//gaps := []int{701, 301, 132, 57, 23, 10, 4, 1}

	// Start with the largest gap and work down to a gap of 1.
	//for _, gap := range gaps {
	for gap := n / 2; gap > 0; gap /= 2 {
		for i := gap; i < n; i++ {
			// Then do insertion sort within each group. each element gap is `gap`
			key := arr[i]
			j := i
			for ; j-gap >= 0 && comparator(key, arr[j-gap]) < 0; j -= gap {
				arr[j] = arr[j-gap]
			}
			arr[j] = key
		}
	}
}
