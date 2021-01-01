package stack

// https://leetcode-cn.com/problems/min-stack/
// required - the time of GetMin() method is O(1)

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
type MinStack struct {
	elements []int
	mins     []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{elements: []int{}, mins: []int{}}
}

func (this *MinStack) Push(x int) {
	this.elements = append(this.elements, x)
	if len(this.mins) == 0 {
		this.mins = append(this.mins, x)
		return
	}
	if this.GetMin() > x {
		this.mins = append(this.mins, x)
	} else {
		this.mins = append(this.mins, this.GetMin())
	}
}

func (this *MinStack) Pop() {
	this.mins = this.mins[:len(this.mins)-1]
	this.elements = this.elements[:this.Len()-1]
}

func (this *MinStack) Top() int {
	// assert this stack is not empty
	return this.elements[len(this.elements)-1]
}

func (this *MinStack) GetMin() int {
	// assert this stack is not empty
	return this.mins[len(this.mins)-1]
}

func (this *MinStack) Len() int {
	return len(this.elements)
}
