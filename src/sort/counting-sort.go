package sort

// Counting sort is a lineal time sort algorithm, input limited [0-x]
// Time: O(n)
// Space: O(n)
// `arr` - all array elements limit to [0-x]
// `k` - k is the maximum of array.
func CountingSort(arr []int, k int) {
	c := make([]int, k+1)
	for _, v := range arr {
		c[v]++
	}
	index := 0
	for i := 0; i <= k; i++ {
		for j := 0; j < c[i]; j++ {
			arr[index] = i
			index++
		}
	}
}
