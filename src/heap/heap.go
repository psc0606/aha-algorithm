package heap

// max root heap
type MaxHeap struct {
	arr []int
}

// Time complexity is O(nlogn), not O(n).
// build a heap by array
func (h *MaxHeap) buildHeap(arr []int) {
	for _, e := range arr {
		h.Insert(e)
	}
}

// return true, if the heap is empty.
func (h *MaxHeap) IsEmpty() bool {
	return len(h.arr) == 0
}

// return the number of elements in the heap.
// Time: O(1)
func (h *MaxHeap) Size() int {
	return len(h.arr)
}

// private method for keep the the attribute of the heap when a element is inserted into the heap.
func (h *MaxHeap) shiftUp(index int) {
	parentIndex := (index - 1) / 2
	if index < 1 || parentIndex < 0 {
		return
	}
	// current index larger than parent, then swap both.
	if h.arr[index] > h.arr[parentIndex] {
		swap(h.arr, index, parentIndex)
	}
	h.shiftUp(parentIndex)
}

// private method for keep the attribute of the heap when a element is deleted from the heap top.
func (h *MaxHeap) shiftDown(index int) {
	leftChildIndex := 2*index + 1
	rightChildIndex := leftChildIndex + 1

	var indexOfMaxChild int
	if leftChildIndex < h.Size() && rightChildIndex < h.Size() {
		indexOfMaxChild = max(h.arr, leftChildIndex, rightChildIndex)
	} else if leftChildIndex < h.Size() {
		indexOfMaxChild = max(h.arr, leftChildIndex, leftChildIndex)
	} else if rightChildIndex < h.Size() {
		indexOfMaxChild = max(h.arr, rightChildIndex, rightChildIndex)
	} else {
		return
	}

	if h.arr[index] < h.arr[indexOfMaxChild] {
		swap(h.arr, index, indexOfMaxChild)
		h.shiftDown(indexOfMaxChild)
	} else if h.arr[index] < h.arr[indexOfMaxChild] {
		swap(h.arr, index, indexOfMaxChild)
		h.shiftDown(indexOfMaxChild)
	} else {
		return
	}
}

// insert a element into the heap
func (h *MaxHeap) Insert(ele int) {
	// append it to the last, then `shitUp` to keep the attribute of heap.
	h.arr = append(h.arr, ele)
	h.shiftUp(h.Size() - 1)
}

// remove and return the element of heap top.
func (h *MaxHeap) Remove() int {
	// delete the top element, then move the last element to the top,
	// then `shiftDown` to keep the attribute of heap.
	if h.IsEmpty() {
		panic("empty heap, peek panic")
	}
	result := h.arr[0]
	h.arr[0] = h.arr[h.Size()-1]
	h.arr = h.arr[0 : h.Size()-1]
	h.shiftDown(0)
	return result
}

// peek root of heap
func (h *MaxHeap) Peek() int {
	if len(h.arr) > 0 {
		return h.arr[0]
	}
	// fmt.Sprintf("%v", i)
	panic("empty heap, peek panic")
}

// private method
func swap(arr []int, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// private method
func max(arr []int, i int, j int) int {
	if arr[i] > arr[j] {
		return i
	} else {
		return j
	}
}
