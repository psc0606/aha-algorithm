package array

import (
	"math/rand"
	"time"
)

// https://leetcode-cn.com/problems/top-k-frequent-elements/

// Solution: partition from quick sort, it will update in place. it's not fitted for big data.
// Time: O(1)
// Space: O(n)
func topKFrequent(nums []int, k int) []int {
	freqs := make(map[int]int)
	// count frequency
	for _, num := range nums {
		freqs[num]++
	}
	var values [][]int
	for k, v := range freqs {
		values = append(values, []int{k, v})
	}
	start := 0
	end := len(values) - 1
	index := partition(values, start, end)
	for index != k-1 {
		if index > k-1 {
			end = index - 1
			index = partition(values, start, end)
		}
		if index < k-1 {
			start = index + 1
			index = partition(values, start, end)
		}
	}
	var ret []int
	for _, v := range values[:index+1] {
		ret = append(ret, v[0])
	}
	return ret
}

func partition(values [][]int, start int, end int) int {
	rand.Seed(time.Now().UnixNano())
	picked := rand.Int()%(end-start+1) + start
	// random pick a element as pivot
	// swap the pivot to the last element
	pivot := values[picked][1]
	values[picked], values[end] = values[end], values[picked]

	i := start
	for j := start; j < end; j++ {
		if values[j][1] >= pivot {
			// swap
			values[i], values[j] = values[j], values[i]
			i++
		}
	}
	// pivot be swapped
	values[i], values[end] = values[end], values[i]
	return i // position of pivot
}

// Solution two: base on heap, it's fitted for big data.
type IHeap [][2]int

func (h IHeap) Len() int { return len(h) }

func (h IHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }

func (h IHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}
