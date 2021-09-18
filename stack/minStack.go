package stack

type MinStack struct {
	data []int
	min  int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	s := new(MinStack)
	s.data = make([]int, 0)
	s.min = 0
	return *s
}

func (this *MinStack) Push(val int) {
	if len(this.data) == 0 {
		this.min = val
	}
	this.data = append(this.data, (val - this.min))
	if this.min > val {
		this.min = val
	}
}

func (this *MinStack) Pop() {
	p := this.data[len(this.data)-1]
	if p < 0 {
		this.min = this.min - this.data[len(this.data)-1]
	}
	this.data = this.data[:len(this.data)-1]
}

func (this *MinStack) Top() int {
	p := this.data[len(this.data)-1]
	if p < 0 {
		return this.min
	} else {
		return p + this.min
	}
}

func (this *MinStack) GetMin() int {
	return this.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
